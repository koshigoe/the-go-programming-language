package main

import (
	"net/http"
	"os"
	"fmt"
	"log"
	"time"
)

// WaitForServer は URL のサーバへ接続を試みます。
// 指数バックオフを使って一分間試みます。
// すべての試みが失敗したらエラーを報告します。
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			log.Printf("OK: %s\n", url)
			return nil // 成功
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // 指数バックオフ
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	if err := WaitForServer(os.Args[1]); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
}
