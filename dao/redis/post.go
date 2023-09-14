package redis

import (
	"bluebell/models"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func getIDsFromKey(key string, page, size int64) ([]string, error) {
	//2. 确定查询的起始点
	start := (page - 1) * size
	end := start + size - 1

	//3. ZREVRANGE 按分数从大到小的顺序查询
	return client.ZRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	//根据用户传递的字段,从redis中获取id
	//1. 根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return getIDsFromKey(key, p.Page, p.Size)
}

// GetPostVoteData 根据ids查询每篇帖子的投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//使用pipeline一次发送多条命令,减少RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		zap.L().Error("pipeline.Exec() error", zap.Error(err))
		return nil, err
	}

	data = make([]int64, 0, len(cmders))

	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return data, err
}

// GetCommunityPostIDsInOrder orderKey
func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	// 使用zinterstore 把分区的帖子set与帖子分数的zset 生成新的zset
	// 针对新的zset 按之前的逻辑取数据

	//社区的key
	cKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(p.CommunityID)))

	//利用 缓存key 减少zinterstore执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(key).Val() < 1 {
		// 不存在,需要计算
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return getIDsFromKey(key, p.Page, p.Size)
}
