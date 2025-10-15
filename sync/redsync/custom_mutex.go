package main

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type MiniRedisMutex struct {
	name   string
	expiry time.Duration
	conn   redis.Cmdable
}

func NewMutex(name string, expiry time.Duration, conn redis.Cmdable) *MiniRedisMutex {
	return &MiniRedisMutex{name, expiry, conn}
}

// Lock 加锁操作，随机生成唯一的 value 作为锁的标识
func (m *MiniRedisMutex) Lock(ctx context.Context, value string) (bool, error) {
	reply, err := m.conn.SetNX(ctx, m.name, value, m.expiry).Result()
	if err != nil {
		return false, err
	}

	return reply, nil
}

var deleteScript = `
	local val = redis.call("GET", KEYS[1])
	if val == ARGV[1] then
		return redis.call("DEL", KEYS[1])
	elseif val == false then
		return -1
	else
		return 0
	end
`

func (m *MiniRedisMutex) Unlock(ctx context.Context, value string) (bool, error) {
	status, err := m.conn.Eval(ctx, deleteScript, []string{m.name}, value).Result()
	if err != nil {
		return false, err
	}
	if status == int64(-1) {
		return false, errors.New("expired")
	}
	return status != int64(0), nil
}
