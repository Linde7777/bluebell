package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

const oneWeekInSeconds = 24 * 7 * 60 * 60
const scorePerVote = 432

var ErrVoteTimeExpire = errors.New("the voting time has expired")

func CreatePost(postID int64) error {
	_, err := rdb.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()
	return err
}

// VoteForPost takes an argument `value` which is the
// `direction` in `models.ParamVoteData`,
// 1 for upvoting, -1 for downvoting, 0 for canceling the vote,
func VoteForPost(userID, postID string, value float64) error {
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}

	oldValue := rdb.ZScore(getRedisKey(
		KeyPostVotedZSetPrefix+postID), userID).Val()
	var direction float64
	if value > oldValue {
		direction = 1
	} else {
		direction = -1
	}
	diff := math.Abs(oldValue - value)
	_, err := rdb.ZIncrBy(getRedisKey(KeyPostScoreZSet),
		direction*diff*scorePerVote, postID).Result()
	if err != nil {
		return err
	}

	if value != 0 {
		_, err = rdb.ZAdd(getRedisKey(
			KeyPostVotedZSetPrefix+postID), redis.Z{
			Score: value, Member: userID}).Result()
		if err != nil {
			return err
		}
	} else {
		_, err = rdb.ZRem(getRedisKey(
			KeyPostVotedZSetPrefix+postID), userID).Result()
		if err != nil {
			return err
		}
	}

	return err
}
