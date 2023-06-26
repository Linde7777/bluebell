package redis

const (
	Prefix                 = "bluebell:"
	KeyPostTimeZSet        = "post:time"
	KeyPostScoreZSet       = "post:score"
	KeyPostVotedZSetPrefix = "post:voted:"
)

func getRedisKey(key string) string {
	return Prefix + key
}
