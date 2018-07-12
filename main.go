package main

import (
	"fmt"
	"math"
	"time"
)

const maxNum = 100000 // 素数を調べる最大値

func main() {

	list := []int{} // 整数リスト

	// 整数のリストに初期値2～maxNumをセット(0,1は素数ではないので除外)
	for j := 0; j < maxNum-2; j++ {
		list = append(list, j+2)
	}

	// エラトステネスの篩
	start := time.Now()
	furui(list)
	end := time.Now()

	// 普通に素数を求める
	start2 := time.Now()
	normal(list)
	end2 := time.Now()

	fmt.Printf("エラトステネスの篩 -> %f秒\n", (end.Sub(start)).Seconds())
	fmt.Printf("普通に素数を求める -> %f秒\n", (end2.Sub(start2)).Seconds())
}

// エラトステネスの篩
func furui(list []int) {

	// ループ最大値取得(maxNumの平方根)
	loopMax := math.Sqrt(maxNum)

	var tmp []int
	var target = 1

	// 素数を調べる(maxNumの平方根までループすればいい)
	for index := 0; target < int(loopMax); index++ {
		target = list[index] // 素数候補取得

		for i := 0; i < len(list); i++ {
			if list[i]%target != 0 {
				// 素数候補で割り切れる数をふるい落とす(割り切れない数をtmpにセット)
				tmp = append(tmp, list[i])
			} else if list[i] == target {
				// 素数候補自身はOK、tmpにセット
				tmp = append(tmp, list[i])
			}
		}

		list = tmp // ふるい落とした結果をセット
		tmp = nil

	}

	// 素数を表示
//	for _, sosu := range list {
//		fmt.Println(sosu)
//	}
}

// 普通に素数を求める
func normal(list []int) {

	var tmp2 []int

	var isNotSosuFlg = false

	// 最大値まで回していき、自分自身以外で割り切れない（約数がない）を求める
	for j := 0; j < len(list); j++ {

		// list[0]から始めて自分自身までループし割り切れるか？
		for k := 0; k < len(list); k++ {
			if (list[j]%list[k] == 0) && (list[j] != list[k]) {
				// 自分自身以外で割り切れたら素数ではない
				isNotSosuFlg = true
				break
			} else if list[k] == list[j] {
				// 上記以外で自分自身は素数
				break
			}
		}

		if !isNotSosuFlg {
			tmp2 = append(tmp2, list[j])
		} else {
			isNotSosuFlg = false
		}
	}

	list = tmp2 // 結果をセット

	// 素数を表示
	//for _, sosu := range list {
//		fmt.Println(sosu)
	//}
}

