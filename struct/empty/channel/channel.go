package main

func main() {
	done := make(chan struct{})

	go func() {
		// Do something
		done <- struct{}{}
	}()

	// Waiting
	<-done
	// Done

	// 其他写法
	go func() {
		// do something
		close(done)
	}()

	<-done
}
