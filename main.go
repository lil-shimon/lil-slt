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

		res := calc(slot)

		profit := res * bet

		if res > 0 {
			fmt.Printf("%v円もうけました", profit)
			fmt.Println("")
		}

		money = money + profit
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

func calc(slot [][]int) int {
	res := 0
	for x := 0; x < 3; x++ {
		if slot[x][0] == slot[x][1] && slot[x][0] == slot[x][2] {
			calcProfit(slot[x][0], &res)
		}
		for y := 0; y < 3; y++ {
			if slot[0][y] == slot[1][y] && slot[0][y] == slot[2][y] {
				calcProfit(slot[0][y], &res)
			}
		}
	}
	return res
}

func calcProfit(value int, profit *int) {
	switch value {
	case 0:
		*profit += 1
	case 1:
		*profit += 1
	case 2:
		*profit += 2
	case 3:
		*profit += 4
	case 5:
		*profit += 10
	case 6:
		*profit += 20
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
