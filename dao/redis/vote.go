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
var ErrRepeatVoting = errors.New("repeating voting")

func CreatePost(postID int64) error {
	_, err := rdb.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()
	return err
}

// VoteForPost takes an argument `currDirection` which is the
// `direction` in `models.ParamVoteData`,
// 1 for upvoting, -1 for downvoting, 0 for canceling the vote,
func VoteForPost(userID, postID string, currDirection float64) error {
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}

	oldDirection := rdb.ZScore(getRedisKey(
		KeyPostVotedZSetPrefix+postID), userID).Val()
	var futureDirection float64
	if currDirection == oldDirection {
		return ErrRepeatVoting
	}
	if currDirection > oldDirection {
		futureDirection = 1
	} else {
		futureDirection = -1
	}
	diff := math.Abs(oldDirection - currDirection)
	_, err := rdb.ZIncrBy(getRedisKey(KeyPostScoreZSet),
		futureDirection*diff*scorePerVote, postID).Result()
	if err != nil {
		return err
	}

	if currDirection != 0 {
		_, err = rdb.ZAdd(getRedisKey(
			KeyPostVotedZSetPrefix+postID), redis.Z{
			Score: currDirection, Member: userID}).Result()
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

func GetPostVotingData(IDs []string) (data []int64, err error) {
	pipeline := rdb.Pipeline()
	for _, id := range IDs {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}

	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
