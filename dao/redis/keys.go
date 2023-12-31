package redis

// redis key

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"   //zset;帖子及发帖时间
	KeyPostScoreZSet       = "post:score"  //zset;帖子及投票分数
	KeyPostVotedZSetPrefix = "post:voted:" //zset;记录用户及投票类型;参数是帖子post id
	KeyCommunitySetPrefix  = "community:"  //set;保存每个分区下 帖子的id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
