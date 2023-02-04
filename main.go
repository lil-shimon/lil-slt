package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var bet int
	money := 100

	for money > 0 {
		fmt.Printf("所持金は残り%v円です", money)
		fmt.Println("")
		fmt.Println("いくらBETしますか?:")
		fmt.Scan(&bet)

		if bet > money {
			fmt.Println("所持金より少ない額でBETしてください。")
			fmt.Println("")
			continue
		}

		if bet == 0 {
			fmt.Println("掛け金は１円以上で入力してください")
			fmt.Println("")
			continue
		}

		fmt.Printf("%v円BETします\n", bet)
		money = money - bet

		slot := Generate3x3Slice()

		fmt.Println("")
		drawSlot(slot)
		fmt.Println("")
	}

	fmt.Println("ゲームオーバー")
}

func drawSlot(slot [][]int) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Printf("%d ", slot[x][y])
		}
		fmt.Println()
	}
}

func GenerateRandomSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(7) // 0-7のランダムな数字を生成
	}
	return s
}

func Generate3x3Slice() [][]int {
	rand.Seed(time.Now().UnixNano())
	s := GenerateRandomSlice(9) // 9つの要素を持つスライスを生成
	return [][]int{
		{s[0], s[1], s[2]},
		{s[3], s[4], s[5]},
		{s[6], s[7], s[8]},
	}
}
