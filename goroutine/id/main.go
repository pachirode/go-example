package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func processId() {
	// 获取当前进程的 id
	pid := os.Getpid()
	fmt.Println("process id:", pid)

	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("找不到进程:", err)
		return
	}

	err = proc.Kill()
	if err != nil {
		fmt.Println("杀死进程失败")
	}

	fmt.Println("进程正在运行")
}

func gId() int {
	buf := make([]byte, 32)
	n := runtime.Stack(buf, false) //仅获取当前 goroutine 的堆栈信息，返回值表示当前堆栈长度
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		//i := i go 1.22-
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i, gId())
		}()
	}
	wg.Wait()
}
