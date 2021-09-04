package main

import (
	"fmt"
)

type obj struct {
	x        int
	y        int
	typeChar string
}

type character struct {
	x        int
	y        int
	typeChar string
	hp       int
	pocket   map[string]int
}

type room struct {
	size   int
	matrix [16][16]obj
}

func newRoomStr(o [16][16]string) room {
	var result room
	result.size = len(o)
	for i := 0; i < result.size; i++ {
		for j := 0; j < result.size; j++ {
			result.matrix[i][j].typeChar = o[i][j]
			result.matrix[i][j].x = i
			result.matrix[i][j].y = j
		}
	}
	return result
}

func (m room) writeMap() {
	fmt.Println("\033[2J") //only in bash
	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			o := m.matrix[i][j]
			fmt.Print(o.typeChar)
		}
		fmt.Print("\n")
	}
}
func movePlayer(x int, y int, pl *character, m *room) bool { //bool true - new lvl
	if m.matrix[x][y].typeChar == " " {
		m.matrix[pl.x][pl.y].typeChar = " "
		m.matrix[x][y].typeChar = pl.typeChar
		pl.x, pl.y = x, y
	} else if m.matrix[x][y].typeChar == "⊞" { //Trapdoor
		if pl.pocket["key"] > 0 {
			m.matrix[pl.x][pl.y].typeChar = " "
			m.matrix[x][y].typeChar = pl.typeChar
			pl.x, pl.y = x, y
			pl.pocket["key"]--
			return true
		}
	} else if m.matrix[x][y].typeChar == "⚿" {
		m.matrix[pl.x][pl.y].typeChar = " "
		m.matrix[x][y].typeChar = pl.typeChar
		pl.x, pl.y = x, y
		pl.pocket["key"]++
	}
	m.writeMap()
	return false
}

func game(player *character, currentRoom *room) bool {
	movePlayer(1, 1, player, currentRoom)
	var key string
	var x int
	var y int
	nextLvl := false
	for {
		fmt.Scanf("%s\n", &key)
		switch key[0] {
		case 'w':
			x = player.x - 1
			y = player.y
		case 'a':
			x = player.x
			y = player.y - 1
		case 's':
			x = player.x + 1
			y = player.y
		case 'd':
			x = player.x
			y = player.y + 1
		case 'v':
			nextLvl = false
			return nextLvl
		default:
			x = player.x
			y = player.y
		}
		if movePlayer(x, y, player, currentRoom) == true {
			nextLvl = true
			return nextLvl
		}
	}
}

func main() {
	player := character{
		x:        1,
		y:        1,
		typeChar: "o",
		hp:       100,
		pocket:   make(map[string]int),
	}
	room0 := newRoomStr([16][16]string{
		{"■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", "⚿", " ", " ", " ", " ", "⊞", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■"},
	})
	room1 := newRoomStr([16][16]string{
		{"■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■"},
		{"■", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "■", " ", " ", " ", "■"},
		{"■", " ", "■", "■", " ", "■", " ", " ", " ", "■", " ", " ", " ", "■", " ", "■"},
		{"■", " ", "⊞", "■", " ", "■", "■", " ", " ", "■", " ", "■", "■", "■", " ", "■"},
		{"■", " ", " ", "■", " ", " ", "■", " ", " ", "■", " ", " ", "■", " ", " ", "■"},
		{"■", " ", "■", "■", "■", " ", "■", "■", " ", "■", "■", " ", "■", " ", "■", "■"},
		{"■", " ", " ", " ", "■", " ", " ", "■", " ", "■", "⚿", " ", "■", " ", " ", "■"},
		{"■", "■", "■", " ", "■", "■", " ", "■", " ", "■", " ", " ", "■", "■", " ", "■"},
		{"■", " ", " ", " ", " ", "■", " ", "■", " ", "■", " ", " ", "■", " ", " ", "■"},
		{"■", " ", "■", "■", "■", "■", " ", "■", " ", "■", "■", "■", "■", " ", " ", "■"},
		{"■", " ", " ", " ", " ", "■", " ", "■", " ", "■", " ", " ", " ", " ", "■", "■"},
		{"■", "■", "■", "■", " ", "■", " ", "■", " ", "■", " ", "■", "■", "■", "■", "■"},
		{"■", " ", " ", " ", " ", "■", " ", "■", " ", "■", " ", " ", " ", " ", " ", "■"},
		{"■", " ", "■", "■", "■", "■", " ", "■", " ", "■", "■", "■", "■", "■", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", "■", " ", " ", " ", " ", " ", " ", " ", "■"},
		{"■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■", "■"},
	})
	if game(&player, &room0) == false {
		return
	}
	if game(&player, &room1) == false {
		return
	}
	fmt.Println("\033[2J")
	fmt.Println("Win!!!!!!!!")
}

//---------------- old things ---------------------------------------------------------------

/*

	time.Sleep(time.Second / 60)

		{"*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", " ", "*", " ", " ", " ", " ", "*"},
		{"*", " ", "*", " ", "*", "*", " ", "*"},
		{"*", " ", "*", "*", "*", " ", " ", "*"},
		{"*", " ", " ", " ", " ", " ", " ", "*"},
		{"*", "*", "*", " ", "*", "*", " ", "*"},
		{"*", " ", " ", " ", "*", " ", " ", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*"},

		{"■", "■", "■", "■", "■", "■", "■", "■"},
		{"■", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", "■"},
		{"■", " ", " ", " ", " ", " ", " ", "■"},
		{"■", "■", "■", "■", "■", "■", "■", "■"},
*/
