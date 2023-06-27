package redis

import (
	"bluebell/models"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

func GetIDsFromKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	return rdb.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamsPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return GetIDsFromKey(key, p.Page, p.Size)
}

func GetCommunityPostIDsInOrder(
	p *models.ParamsCommunityPostList) ([]string, error) {

	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	communityIDStr := strconv.Itoa(int(p.CommunityID))
	communityKey := getRedisKey(KeyCommunityPostSetPrefix + communityIDStr)
	cacheKey := orderKey + communityIDStr
	if rdb.Exists(cacheKey).Val() < 1 {
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(cacheKey, redis.ZStore{
			Aggregate: "MAX",
		}, communityKey, orderKey)
		pipeline.Expire(cacheKey, 60*time.Second)
		if _, err := pipeline.Exec(); err != nil {
			return nil, err
		}
	}

	return GetIDsFromKey(cacheKey, p.Page, p.Size)
}
