package redis

import "bluebell/models"

func GetPostIDsInOrder(p *models.ParamsPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	return rdb.ZRevRange(key, start, end).Result()
}
