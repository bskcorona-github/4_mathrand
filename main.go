package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 乱数のシードを設定するために現在時刻を使う
	rand.Seed(time.Now().UnixNano())

	// 生成する乱数の範囲を指定する
	min := 1
	max := 100

	// ランダムな整数を生成する
	randomNumber := rand.Intn(max-min+1) + min

	// 生成したランダムな数値を出力する
	fmt.Println("Random number:", randomNumber)
}
