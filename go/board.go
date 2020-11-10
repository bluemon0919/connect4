package main

import "fmt"

// BoardRow ボードの行数
const BoardRow int = 6

// BoardColumn ボードの列数
const BoardColumn int = 7

// StackMax 縦に積める最大数
const StackMax int = 6

// Board ゲームボードです
type Board struct {
	MainBoard    [BoardRow + 2][BoardColumn + 2]int // ボード自身
	StackedCount [BoardColumn + 2]int               // 縦に何個積んでいるかの情報
	History      []int
}

// NewBoard 新しいボードを生成します
func NewBoard() *Board {
	b := &Board{}
	for i := 0; i < len(b.MainBoard[0]); i++ {
		b.MainBoard[0][i] = -1
		b.MainBoard[BoardRow+1][i] = -1
	}
	for i := 0; i < len(b.MainBoard); i++ {
		b.MainBoard[i][0] = -1
		b.MainBoard[i][BoardColumn+1] = -1
	}
	b.History = make([]int, 0)
	return b
}

// Draw ボードの内容を表示します
func (b *Board) Draw() {
	for _, row := range b.MainBoard {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

// Input 石を投入する
func (b *Board) Input(stone int, column int) error {
	b.History = append(b.History, column)
	if 0 > column || column >= len(b.MainBoard[0]) {
		return fmt.Errorf("範囲外です")
	}
	// ボードの行を算出する
	fmt.Println("Row[column]:", b.StackedCount[column])
	row := b.StackedCount[column]
	if row > StackMax {
		return fmt.Errorf("もう置けませぬ")
	}
	if b.MainBoard[StackMax-row][column] == -1 {
		return fmt.Errorf("そこには置けませぬ")
	}

	// ボードと投入情報を更新する
	// 下段から順に投入する
	b.MainBoard[StackMax-row][column] = stone
	b.StackedCount[column]++
	return nil
}

// Judge 勝敗を判定します
func (b *Board) Judge(stone int, column int) (bool, error) {
	if 0 >= column || column > BoardColumn {
		return false, fmt.Errorf("範囲オーバーです")
	}
	cnt := b.StackedCount[column]
	row := StackMax - cnt + 1

	left := b.left(stone, column-1, row)
	right := b.right(stone, column+1, row)
	down := b.down(stone, column, row+1)
	leftDown := b.leftDown(stone, column-1, row+1)
	rightUp := b.rightUp(stone, column+1, row-1)
	leftUp := b.leftUp(stone, column-1, row-1)
	rightDown := b.rightDown(stone, column+1, row+1)
	/*
		fmt.Println("cnt:", cnt)
		fmt.Println("row:", row)
		fmt.Println("left:", left)
		fmt.Println("right:", right)
		fmt.Println("down:", down)
		fmt.Println(column-1, StackMax-row+1+1, "leftDown:", leftDown)
		fmt.Println(column+1, StackMax-row+1-1, "rightUp:", rightUp)
		fmt.Println(column-1, StackMax-row+1-1, "leftUp:", leftUp)
		fmt.Println(column+1, StackMax-row+1+1, "rightDown:", rightDown)
	*/
	switch {
	case left+right+1 >= 4:
		return true, nil
	case down+1 >= 4:
		return true, nil
	case leftDown+rightUp+1 >= 4:
		return true, nil
	case leftUp+rightDown+1 >= 4:
		return true, nil
	}
	return false, nil
}

func (b *Board) left(stone int, column int, row int) int {
	if b.MainBoard[row][column] == stone {
		return b.left(stone, column-1, row) + 1
	}
	return 0
}

func (b *Board) right(stone int, column int, row int) int {
	if b.MainBoard[row][column] == stone {
		return b.right(stone, column+1, row) + 1
	}
	return 0
}

func (b *Board) down(stone int, column int, row int) int {
	if b.MainBoard[row][column] == stone {
		return b.down(stone, column, row+1) + 1
	}
	return 0
}

func (b *Board) leftUp(stone int, column int, row int) int {
	if b.MainBoard[row][column] == stone {
		return b.leftUp(stone, column-1, row-1) + 1
	}
	return 0
}

func (b *Board) leftDown(stone int, column int, row int) int {
	if b.MainBoard[row][column] == stone {
		return b.leftDown(stone, column-1, row+1) + 1
	}
	return 0
}

func (b *Board) rightUp(stone int, column int, row int) int {
	if b.MainBoard[row][column] == stone {
		return b.rightUp(stone, column+1, row-1) + 1
	}
	return 0
}

func (b *Board) rightDown(stone int, column int, row int) int {
	if b.MainBoard[row][column] == stone {
		return b.rightDown(stone, column+1, row+1) + 1
	}
	return 0
}

// ViewHistory 履歴を表示する
func (b *Board) ViewHistory() {
	for _, i := range b.History {
		fmt.Print(i)
	}
	fmt.Println()
}
