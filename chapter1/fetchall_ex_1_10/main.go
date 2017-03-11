// 練習問題 1.10 大量のデータを生成するウェブサイトを見つけなさい。報告される時間が大きく変化するかを調べるために
// fetchall を2回続けて実行して、キャッシュされているかどうかを調査しなさい。毎回同じ内容を得ているでしょうか。
// fetchall を修正して、その出力をファイルへ保存するようにして調べら得る様にしなさい。
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // ゴルーチンを開始
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // ch チャネルから受信
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // ch チャネルへ送信
		return
	}

	cacheKey := make([]byte, 10)
	_, err = rand.Read(cacheKey)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	filename := fmt.Sprintf("%x.%x", sha256.Sum256([]byte(url)), sha256.Sum256(cacheKey))

	output, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	nbytes, err := io.Copy(output, resp.Body)
	resp.Body.Close() // 資源をリークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
