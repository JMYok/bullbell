package redis

import "bluebell/models"

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
