package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

/**
direction = 1:
	1. 之前没有投过票，现在投赞成票 差值的绝对值：1
	2. 之前投反对票，现在改投赞成票 差值的绝对值：2
direction = 0:
	1. 之前投过赞成票，现在要取消投票 差值的绝对值：1
	2. 之前投过反对票，现在要取消投票 差值的绝对值：1
direction = -1:
	1. 之前没有投过票，现在投反对票 差值的绝对值：1
	2. 之前投赞成票，现在改投反对票 差值的绝对值：2

一天86400秒，86400/200 = 432 ----> 200张赞成票可为帖子续一天

投票限制：
每个帖子从发表之日起，一个星期之内允许用户投票
	1. 到期之后redis->mysql
	2. 到期之后删除 KeyPostVotedZSetPrefix
*/

// VoteForPost currentValue代表用户投的是赞成票还是反对票-1 0 1
func VoteForPost(userID, postID string, currentValue float64) (err error) {
	//1.判读投票限制
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()

	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}

	//2.更新帖子的分数
	//查询当前用户给当前帖子的投票记录(-1 0 1)
	oldValue := client.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()

	//控制分数的增减
	var op float64
	if currentValue > oldValue {
		op = 1
	} else {
		op = -1
	}

	diff := math.Abs(oldValue - currentValue)

	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)

	//3.记录该用户为帖子投票的数据
	if currentValue == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  currentValue,
			Member: userID,
		})
	}
	_, err = pipeline.Exec()
	return err
}

func CreatePost(postID uint64) error {
	pipeline := client.TxPipeline()
	//帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	//帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}
