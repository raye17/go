package cache

import "time"

func SetBlackLoginKey(key string, value string, duration time.Duration) error {
	return RedisClient.Set(key, value, duration).Err()
}
func SetFailLoginKey(key string, value string, duration time.Duration) error {
	return RedisClient.Set(key, value, duration).Err()
}
