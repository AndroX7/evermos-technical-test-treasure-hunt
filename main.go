package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var clear map[string]func()

func init() {
	GenerateBoard()
	endGame = false
	clear = make(map[string]func())
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
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

var key string
var board Board
var step int
var endGame bool

type Board struct {
	Column []Row
}
type Row struct {
	Row []string
}

var winRow, winCol int

func GenerateBoard() {
	var row Row
	var treasureCol, treasureRow int
	rand.Seed(time.Now().Unix())
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 7; j++ {
			if i == 0 || i == 5 || j == 0 || j == 7 {
				row.Row = append(row.Row, "#")
			} else {
				row.Row = append(row.Row, ".")
			}
		}
		board.Column = append(board.Column, row)
		row.Row = []string{}
	}
	for index, col := range board.Column {
		if index == 2 {
			col.Row[2] = "#"
			col.Row[3] = "#"
			col.Row[4] = "#"
		} else if index == 3 {
			col.Row[4] = "#"
			col.Row[6] = "#"
		} else if index == 4 {
			col.Row[1] = "X"
			col.Row[2] = "#"
		}
	}
	for {
		treasureRow = 1 + rand.Intn(6-1+1)
		treasureCol = 1 + rand.Intn(4-1+1)
		if board.Column[treasureCol].Row[treasureRow] != "#" && board.Column[treasureCol].Row[treasureRow] != "X" {
			// board.Column[treasureCol].Row[treasureRow] = "$"
			winRow = treasureRow
			winCol = treasureCol
			break
		}
	}
}
func printBoard() {
	for _, col := range board.Column {
		fmt.Println(strings.Join(col.Row, "  "))
		fmt.Println("")
	}
}

func main() {
	for {
		printBoard()
		fmt.Println("Enter one single key to move(a: up, b: right, c: down, d: left) then enter: ")
		fmt.Scan(&key)
		for i := 0; i < len(board.Column); i++ {
			for j := 0; j < len(board.Column[i].Row); j++ {
				if board.Column[i].Row[j] == "X" {
					fmt.Println(board.Column[i-1].Row[j])
					if strings.ToLower(strings.TrimSpace(key)) == "a" && board.Column[i-1].Row[j] != "#" {
						board.Column[i].Row[j], board.Column[i-1].Row[j] = board.Column[i-1].Row[j], board.Column[i].Row[j]
					} else if strings.ToLower(strings.TrimSpace(key)) == "b" && board.Column[i].Row[j+1] != "#" {
						board.Column[i].Row[j], board.Column[i].Row[j+1] = board.Column[i].Row[j+1], board.Column[i].Row[j]
					} else if strings.ToLower(strings.TrimSpace(key)) == "c" && board.Column[i+1].Row[j] != "#" {
						board.Column[i].Row[j], board.Column[i+1].Row[j] = board.Column[i+1].Row[j], board.Column[i].Row[j]
					} else if strings.ToLower(strings.TrimSpace(key)) == "d" && board.Column[i].Row[j-1] != "#" {
						board.Column[i].Row[j], board.Column[i].Row[j-1] = board.Column[i].Row[j-1], board.Column[i].Row[j]
					}
					if i == winCol && j == winRow {
						endGame = true
						board.Column[i].Row[j] = "$"
					}
					key = ""
					break
				}
			}
		}
		if endGame {
			fmt.Println("Congratulations you won the game ^O^")
			break
		}
		CallClear()
	}
}
