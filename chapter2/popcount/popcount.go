package popcount

// pc[i] は i のポピュレーションカウントです。
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount は x のポピュレーションカウント(1が設定されているビット数)を返します。
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 練習問題 2.3 単一の式の代わりにループを使う様に PopCount を書き直しなさい。二つのバージョンの性能を比較しなさい。(11.4 説で異なる実装の性能を体系的に非エックスル方法を説明しています。)
func PopCountEx23(x uint64) int {
	var count byte
	for i := 0; i < 8; i++ {
		count += pc[byte(x>>(uint64(i)*8))]
	}
	return int(count)
}

// 練習問題 2.4 引数をビットシフトしながら最下位ビットの検査を 64 回繰り返す事でビット数を数える PopCount のバージョンを作成しなさい。テーブル参照を行うバージョンと性能を比較しなさい。
func PopCountEx24(x uint64) int {
	var count byte
	for i := 0; i < 64; i++ {
		count += byte(x&1)
		x >>= 1
	}
	return int(count)
}

// 練習問題 2.5 式 x&(x-1) は x で 1 が設定されている最下位ビットをクリアします。この事実を使ってビット数を数える PopCount のバージョンを作成し、その性能を評価しなさい。
func PopCountEx25(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
