package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	smsCaptchaExpire    = 5 * time.Minute
	smsCaptchaKeyPrefix = "sms:captcha:%s"

	loginToken = "login:token"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func SetSmsCaptchaToRedis(ctx context.Context, redis *redis.Client, mobile, captcha string) error {
	key := fmt.Sprintf(smsCaptchaKeyPrefix, mobile)
	return redis.Set(ctx, key, captcha, smsCaptchaExpire).Err()
}

func GetSmsCaptchaFromRedis(ctx context.Context, redis *redis.Client, mobile string) (string, error) {
	key := fmt.Sprintf(smsCaptchaKeyPrefix, mobile)
	return redis.Get(ctx, key).Result()
}
