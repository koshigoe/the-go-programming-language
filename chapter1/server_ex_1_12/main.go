// 練習問題 1.12 リサージュ図形のサーバを修正して、URL からパラメータ値を読み取るようにしなさい。
// たとえば、http://localhost:8000/?cycles=20 のような URL では、周回の回数をデフォルトの 5
// ではなく、20 に設定するようにしなさい。
// 文字列パラメータを整数へ変換するために strconv.Atoi 関数を使いなさい。その変換関数のドキュメント
// は go doc strconv.Atoi で見ることができます。
package main

import (
	"log"
	"net/http"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // パレットの最初の色
	blackIndex = 1 // パレットの次の色
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles  := 5     // 発信器 x が完了する周回の回数

		if err := r.ParseForm(); err == nil {
			if v, ok := r.Form["cycles"]; ok {
				if n, err := strconv.Atoi(v[0]); err == nil {
					cycles = n
				}
			}
		}

		lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // 回転の分解能
		size    = 100   // 画像キャンバスは [-size..+size] の範囲を扱う
		nframes = 64    // アニメーションフレーム数
		delay   = 8     // 10ms 単位でのフレーム間の遅延
	)
	freq := rand.Float64() * 3.0 // 発信器 y の相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意: エンコードエラーを無視
}
