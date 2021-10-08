package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	Map      [][]string
	CurrentX int
	CurrentY int
	Step     int
	TotalDot int
	TotalGet int
}

var data Data
var x, y int

func main() {
	initData()
	printMap()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Input your move : ")
		text, _ := reader.ReadString('\n')
		text = strings.ToLower(strings.Replace(text, "\n", "", -1))

		switch text {
		case "a":
			move(-1, 0)
		case "w":
			move(0, -1)
		case "d":
			move(1, 0)
		case "x":
			move(0, 1)
		}

		printMap()

		if data.TotalDot == 0 {
			break
		}
	}

}

func initData() {
	data.Map = [][]string{{"#", "#", "#", "#", "#", "#", "#", "#"},
		{"#", ".", ".", ".", ".", ".", ".", "#"},
		{"#", ".", "#", "#", "#", ".", ".", "#"},
		{"#", ".", ".", ".", "#", ".", "#", "#"},
		{"#", "X", "#", ".", ".", ".", ".", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#"},
	}

	data.Step = 0
	data.CurrentX = 1
	data.CurrentY = 4
	data.TotalDot = 17
	data.TotalGet = 0
}

func printMap() {
	var i, j int
	p := ""
	sMap := data.Map

	for i = 0; i < 6; i++ {
		for j = 0; j < 8; j++ {
			p += fmt.Sprintf("%s", sMap[i][j])
		}

		p += fmt.Sprintf("\n")
	}

	fmt.Println("=================================")
	fmt.Println("Total Step :", data.Step)
	fmt.Println("Total get dot :", data.TotalGet)
	fmt.Println(p)
	fmt.Println("=================================")
}

func move(x, y int) {
	x += data.CurrentX
	y += data.CurrentY
	data.Step++

	if data.Map[y][x] == "." {
		data.Map[y][x] = "$"
		data.Map[data.CurrentY][data.CurrentX] = " "
		data.TotalGet++
		data.TotalDot--
		data.CurrentY = y
		data.CurrentX = x
	} else if data.Map[y][x] == " " {
		data.Map[data.CurrentY][data.CurrentX] = " "
		data.CurrentY = y
		data.CurrentX = x
		data.Map[y][x] = "$"
	} else if data.Map[y][x] == "#" {
		fmt.Println("you broke obstaqle")
		fmt.Println("")
	}
}
