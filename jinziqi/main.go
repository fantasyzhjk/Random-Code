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

func checkBroad() {
	//横向
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][0] == board[i][2] && board[i][0] != " " {
			fmt.Printf("%v win!\n", board[i][0])
			os.Exit(0)
		}
	}
	//纵向
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] && board[0][i] == board[2][i] && board[0][i] != " " {
			fmt.Printf("%v win!\n", board[0][i])
			os.Exit(0)
		}
	}
	//斜向
	if board[0][0] == board[1][1] && board[0][0] == board[2][2] && board[0][0] != " " {
		fmt.Printf("%v win!\n", board[0][0])
		os.Exit(0)
	}
	if board[0][2] == board[1][1] && board[0][2] == board[2][0] && board[0][2] != " " {
		fmt.Printf("%v win!\n", board[0][2])
		os.Exit(0)
	}
	//平局
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return
			}
		}
	}
	fmt.Println("平局")
	os.Exit(0)
}

func main() {
	fmt.Println("Game start!")
	for {
		printBoard()
		checkBroad()
		playerInput()
	}
}
