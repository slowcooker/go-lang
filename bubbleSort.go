package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func main() {
	fmt.Printf("バブルソートを実装していく！")
	// 乱数配列の最大サイズを指定
	var max = 300
	// ランダムの最大値を指定
	var maxNumber = 1000
	rand.Seed(time.Now().UnixNano())
	var loop = rand.Intn(max)
	
	// 乱数配列を作成。
	args := []int{}
    	for i := 0; i < loop; i++ {
		args = append(args, rand.Intn(maxNumber))
    	}
	// ちょっと確認してみる。
	fmt.Println(args, len(args), "\n")
	fmt.Printf("randomMax:" + strconv.Itoa(maxNumber) + "/ Loops:" + strconv.Itoa(loop) + " (Setting" + strconv.Itoa(max) + ")\n")
	
	// ソート開始。
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args); j++ {
			if (j == 0) {
				continue;
			}
			var current = args[j]
			if (args[j - 1] > args[j]) {
				args[j] = args[j - 1]
				args[j - 1] = current
			}
		}
	}
	fmt.Println(args, len(args), "\n")
}
