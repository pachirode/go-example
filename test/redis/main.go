package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func Login(mobile, smsCode string, rdb *redis.Client) (string, error) {
	ctx := context.Background()

	// 查找验证码
	captcha, err := GetSmsCaptchaFromRedis(ctx, rdb, mobile)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", fmt.Errorf("invalid sms code or expired")
		}
		return "", err
	}

	if captcha != smsCode {
		return "", fmt.Errorf("invalid sms code")
	}

	return smsCode, nil
}

func main() {
	rdb := NewRedisClient()
	token, err := Login("XXX", "123456", rdb)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)
}
