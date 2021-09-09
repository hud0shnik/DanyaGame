package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type obj struct {
	x        int
	y        int
	typeChar string
}

type room struct {
	matrix [4][4]obj
}

func newRoomStr(o [4][4]string) room {
	var result room
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result.matrix[i][j].typeChar = o[i][j]
			result.matrix[i][j].x = i
			result.matrix[i][j].y = j
		}
	}
	return result
}

func writeMap(m *room) {
	fmt.Println("\033[2J") //only in bash
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			o := m.matrix[i][j]
			fmt.Print(o.typeChar)
		}
		fmt.Print("\n")
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; ; i++ {
		x := rand.Intn(4)
		y := rand.Intn(4)
		if m.matrix[x][y].typeChar == " " {
			m.matrix[x][y].typeChar = "2"
			break
		}
		if i > 1000 {
			return //end of game
		}
	}
	fmt.Println("\033[2J")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			o := m.matrix[i][j]
			fmt.Print(o.typeChar)
		}
		fmt.Print("\n")
	}
}

func d(m *room) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if m.matrix[i][j].typeChar == m.matrix[i][j+1].typeChar && m.matrix[i][j].typeChar != " " {
				neww, _ := strconv.Atoi(m.matrix[i][j].typeChar)
				m.matrix[i][j].typeChar = " "
				m.matrix[i][j+1].typeChar = strconv.Itoa(neww * 2)
			} else if m.matrix[i][j+1].typeChar == " " {
				m.matrix[i][j+1].typeChar = m.matrix[i][j].typeChar
				m.matrix[i][j].typeChar = " "
			}
		}
		for j := 3; j > 0; j-- {
			if m.matrix[i][j].typeChar == " " {
				m.matrix[i][j].typeChar = m.matrix[i][j-1].typeChar
				m.matrix[i][j-1].typeChar = " "
			}
		}

	}
	writeMap(m)
}

func a(m *room) {
	for i := 0; i < 4; i++ {
		for j := 3; j > 0; j-- { // _ _ 2 4
			if m.matrix[i][j].typeChar != " " && m.matrix[i][j].typeChar == m.matrix[i][j-1].typeChar {
				neww, _ := strconv.Atoi(m.matrix[i][j].typeChar)
				m.matrix[i][j].typeChar = " "
				m.matrix[i][j-1].typeChar = strconv.Itoa(neww * 2)
			} else if m.matrix[i][j-1].typeChar == " " {
				m.matrix[i][j-1].typeChar = m.matrix[i][j].typeChar
				m.matrix[i][j].typeChar = " "
			}
		}
		for j := 3; j > 0; j-- {
			if m.matrix[i][j-1].typeChar == " " {
				m.matrix[i][j-1].typeChar = m.matrix[i][j].typeChar
				m.matrix[i][j-1].typeChar = " "
			}
		}

	}
	writeMap(m)
}

func game(currentRoom *room) bool {
	var key string
	nextLvl := false
	writeMap(currentRoom)
	for {
		fmt.Scanf("%s\n", &key)
		switch key[0] {
		case 'w':
		case 'a':
			a(currentRoom)
		case 's':
		case 'd':
			d(currentRoom)
		case 'v': //v = exit
			nextLvl = false
			return nextLvl
		default:
			break
		}
	}
}

func main() {
	room := newRoomStr([4][4]string{
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
	})
	game(&room)
}
