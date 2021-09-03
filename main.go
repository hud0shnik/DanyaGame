package main

import "fmt"

type obj struct {
	x         int
	y         int
	typeOfObj string
}

type mapa struct {
	size int
	map0 [3][3]obj
}

func newMapaStr(o [3][3]string) mapa {
	var result mapa
	result.size = 3
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
}

func main() {
	psevdoMapStr := [3][3]string{
		{"Г", "¯", "˥"},
		{"|", "*", "|"},
		{"L", "_", "˩"},
	}
	psevdoMap := newMapaStr(psevdoMapStr)
	psevdoMap.writeMap()
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
