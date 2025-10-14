package main

import (
	"context"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

func base() {
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "192.168.52.133:6379",
	})

	pool := goredis.NewPool(client)

	// 管理分布式锁
	rs := redsync.New(pool)

	// 默认会添加锁过期的时间
	mutex := rs.NewMutex("test-redsync")

	ctx := context.Background()

	// 获取锁，获取失败锁以及被其他进程
	if err := mutex.LockContext(ctx); err != nil {
		panic(err)
	}

	dog(mutex, ctx)

	// do something
	println("do something")

	// 释放锁，失败可能以及过期或者不属于该进程
	if _, err := mutex.UnlockContext(ctx); err != nil {
		panic(err)
	}
}

func dog(mutex *redsync.Mutex, ctx context.Context) {
	stopCh := make(chan struct{})
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				if ok, err := mutex.ExtendContext(ctx); !ok || err != nil {
					println("fail to extend mutex")
				}
			case <-stopCh:
				println("exit extend watchdog")
				return
			}
		}
	}()
}

func main() {
	base()
}
