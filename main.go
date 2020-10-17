package main

import (
	"fmt"
	"log"
)

func main() {
	// 新しいボードを用意する
	myBoard := NewBoard()

	for {
		// ボードを表示する
		myBoard.Draw()

		// 順番に石を投入する
		column := 0
		fmt.Println("投入する列を指定してください")
		fmt.Scan(&column)
		stone := getStone()
		err := myBoard.Input(stone, column)
		if err != nil {
			fmt.Println(err)
			fmt.Println("そこには置けません、もう一度")
			stone = getStone()
			continue
		}

		// 判定する
		if judge, err := myBoard.Judge(stone, column); err != nil {
			log.Fatal(err)
		} else if judge {
			// 勝敗が決まった場合は終了する
			break
		} else {
		}
	}
	myBoard.Draw()
	fmt.Println("履歴の表示")
	myBoard.ViewHistory()
	fmt.Println("終了します")
}

var stoneColor = 1

func getStone() int {
	if stoneColor == 1 {
		stoneColor = 2
	} else {
		stoneColor = 1
	}
	return stoneColor
}
