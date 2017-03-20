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
