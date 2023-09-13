package redis

import (
	"bluebell/models"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	//根据用户传递的字段,从redis中获取id
	//1. 根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	//2. 确定查询的起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	//3. ZREVRANGE 按分数从大到小的顺序查询
	return client.ZRange(key, start, end).Result()
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
