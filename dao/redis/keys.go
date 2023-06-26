package redis

const (
	Prefix = "bluebell:"

	// KeyPostTimeZSet represent a zset named post:time,
	// where post_id is the key, the timestamp of when
	// the post created is the value.
	KeyPostTimeZSet = "post:time"

	// KeyPostScoreZSet represent a zset named post:score,
	// where post_id is the key, score is the value.
	KeyPostScoreZSet = "post:score"

	// KeyPostVotedZSetPrefix act literally as its naming.
	// For example: post:voted:123456 represent a zset
	// stored the mapping user_id <-> voting
	// under a post which id is 123456
	KeyPostVotedZSetPrefix = "post:voted:"

	// KeyCommunityPostSetPrefix act literally as its naming.
	// For example: a set named community:56789 store
	// all the posts in the community which id is 56789
	KeyCommunityPostSetPrefix = "community:"
)

func getRedisKey(key string) string {
	return Prefix + key
}
