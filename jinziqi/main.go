package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

//棋盘
var (
	board = [][]string{
		[]string{" ", " ", " "},
		[]string{" ", " ", " "},
		[]string{" ", " ", " "},
	}
	round = "玩家一"
)

func clear() {
	clear := make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear[runtime.GOOS]()
}

func playerInput() {
	var x, y int
	fmt.Printf("请%v输入落子位置\n", round)
	for {
		fmt.Scanln(&x, &y)
		if ((x <= 3 && x >= 1) && (y <= 3 && y >= 1)) && board[x-1][y-1] == " " {
			if round == "玩家一" {
				board[x-1][y-1] = "O"
				round = "玩家二"
			} else if round == "玩家二" {
				board[x-1][y-1] = "X"
				round = "玩家一"
			}
			break
		} else {
			fmt.Println("非法位置")
		}
	}
}

func printBoard() {
	clear()
	for _, i := range board {
		fmt.Print("| ")
		for _, i1 := range i {
			fmt.Printf("%v ", i1)
		}
		fmt.Println("|")
	}
}

func main() {
	fmt.Println("Game start!")
	for {
		printBoard()
		playerInput()
	}
}
