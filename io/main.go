package main

import (
	"io"
	"net/http"
)

func health() {
	resp, _ := http.Get("xxx/healthz")
	defer resp.Body.Close()

	_, _ := io.Copy(io.Discard, resp.Body)

	if resp.StatusCode != http.StatusOK {
		println("失败")
	}
}

func main() {

}
