package main

import (
	"log"
	"log/slog"
	"time"

	"github.com/robfig/cron/v3"
)

func base() {
	c := cron.New(cron.WithSeconds())

	_, err := c.AddFunc("* * * * * *", func() {
		println("每秒执行任务：", time.Now().Format("2006-04-02 15:04:05"))
	})

	if err != nil {
		log.Fatal("添加任务失败：%v", err)
	}

	_, err = c.AddFunc("*/5 * * * * *", func() {
		println("每五秒执行任务：", time.Now().Format("2006-04-02 15:04:05"))
	})
	if err != nil {
		log.Fatal("添加任务失败：%v", err)
	}

	c.Start()
	defer c.Stop()

	time.Sleep(10 * time.Second)
}

type MyJob struct {
	name  string
	count int
}

func (j *MyJob) Run() {
	j.count++
	if j.count == 2 {
		panic("第二次执行触发异常")
	}

	println("每隔 5s 执行任务，count：%d\n", j.name, j.count)
}

func myJobDemo() {
	var (
		spec = "@every 5s"
		job  = &MyJob{name: "test"}
	)

	c := cron.New(cron.WithSeconds())
	_, err := c.AddJob(spec, job)
	if err != nil {
		println("添加任务失败")
	}

}

type cronLogger struct{}

func newCronLogger() *cronLogger {
	return &cronLogger{}
}

func (l *cronLogger) Info(msg string, keysAndValues ...any) {
	slog.Info(msg, keysAndValues...)
}

// Error implements the cron.Logger interface.
func (l *cronLogger) Error(err error, msg string, keysAndValues ...any) {
	slog.Error(msg, append(keysAndValues, "err", err)...)
}

func cronLoggerDemo() {
	logger := &cronLogger{}
	cron.New(
		cron.WithSeconds(),
		cron.WithLogger(logger),
	)
}

func main() {
	//base()
	logger := &cronLogger{}
	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(logger),
		cron.WithChain(
			cron.SkipIfStillRunning(logger),
			cron.Recover(logger),
		),
	)

	var (
		spec = "@every 5s"
		job  = &MyJob{name: "test"}
	)

	_, err := c.AddJob(spec, job)
	if err != nil {
		println("添加任务失败")
	}

	c.Start()
	defer c.Stop()

	time.Sleep(30 * time.Second)
}
