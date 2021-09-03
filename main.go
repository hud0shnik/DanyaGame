package main

import (
	"fmt"
	"time"
)

type obj struct {
	x         int
	y         int
	typeOfObj string
}

type mapa struct {
	size int
	map0 [8][8]obj
}

func newMapaStr(o [8][8]string) mapa {
	var result mapa
	result.size = 8
	for i := 0; i < result.size; i++ {
		for j := 0; j < result.size; j++ {
			result.map0[i][j].typeOfObj = o[i][j]
			result.map0[i][j].x = i
			result.map0[i][j].y = j
		}
	}
	return result
}

func (m mapa) writeMap() {
	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			o := m.map0[i][j]
			fmt.Print(o.typeOfObj)
		}
		fmt.Print("\n")
	}
	time.Sleep(time.Second / 60)
}
func (m *mapa) movePlayer(x int, y int, pl *obj) {
	m.map0[pl.x][pl.y].typeOfObj = " "
	m.map0[x][y].typeOfObj = pl.typeOfObj
	pl.x, pl.y = x, y
}
func main() {
	player := obj{
		x:         1,
		y:         1,
		typeOfObj: "o",
	}
	psevdoMapStr := [8][8]string{
		{"Г", "¯", "¯", "¯", "¯", "¯", "¯", "˥"},
		{"|", " ", " ", " ", " ", " ", " ", "|"},
		{"|", " ", " ", " ", " ", " ", " ", "|"},
		{"|", " ", " ", " ", " ", " ", " ", "|"},
		{"|", " ", " ", " ", " ", " ", " ", "|"},
		{"|", " ", " ", " ", " ", " ", " ", "|"},
		{"|", " ", " ", " ", " ", " ", " ", "|"},
		{"L", "_", "_", "_", "_", "_", "_", "˩"},
	}
	psevdoMap := newMapaStr(psevdoMapStr)
	for {
		for j := 1; j < 7; j++ {
			for i := 1; i < 7; i++ {
				psevdoMap.movePlayer(j, i, &player)
				psevdoMap.writeMap()

			}
			for i := 5; i > 0; i-- {
				psevdoMap.movePlayer(j, i, &player)
				psevdoMap.writeMap()

			}
		}
		for j := 6; j > 0; j-- {
			for i := 1; i < 7; i++ {
				psevdoMap.movePlayer(j, i, &player)
				psevdoMap.writeMap()

			}
			for i := 5; i > 0; i-- {
				psevdoMap.movePlayer(j, i, &player)
				psevdoMap.writeMap()

			}
		}
	}
}

//----------------
/*


func newMapa(o [3][3]obj) mapa {
	return mapa{
		map0: o,
	}
}


*/
/*
	o00 := obj{
		x:         0,
		y:         0,
		typeOfObj: "/",
	}
	o01 := obj{
		x:         0,
		y:         1,
		typeOfObj: "|",
	}
	o10 := obj{
		x:         1,
		y:         0,
		typeOfObj: "|",
	}
	o11 := obj{
		x:         1,
		y:         1,
		typeOfObj: "/",
	}
	matrix := [2][2]obj{
		{o00, o01},
		{o10, o11},
	}

	m0 := newMapa(matrix)
	m0.size = 2
	fmt.Println(m0)
	m0.writeMap()*/

/*
	fmt.Println("101")
	fmt.Println("\033[2J")
	fmt.Println("102")*/
