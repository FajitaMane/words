package main

import (
	"fmt"
	_ "strconv"
)

type Board struct {
	grid [][]*Slot
}

//helper for defining board points
type Point struct {
	i int
	j int
}

//max board index
const MBR = 14

//the NewBoard method returns a new instance of the standard Board
func NewBoard() *Board {
	//first, set all the difficult to determine points of modifiers
	dub_letter_points := make([]Point, 6)
	dub_letter_points[0] = Point{0, 3}
	dub_letter_points[1] = Point{0, 11}
	dub_letter_points[2] = Point{6, 3}
	dub_letter_points[3] = Point{8, 3}
	dub_letter_points[4] = Point{7, 4}
	dub_letter_points[5] = Point{6, 6}
	trip_letter_points := make([]Point, 2)
	trip_letter_points[0] = Point{1, 5}
	trip_letter_points[1] = Point{5, 5}
	//empty tile
	t := Tile{uint8(156)}
	//instantiate the empty grid
	grid := make([][]*Slot, 15)
	for i := range grid {
		grid[i] = make([]*Slot, 15)
		for j := range grid[i] {
			if (i%MBR == 0 && j%MBR == 0) ||
				((i%MBR == 0 || j%MBR == 0) && (i == 7 || j == 7)) {
				grid[i][j] = &Slot{&t, byte('T')}
				//make the double word tiles
			} else if (i == j && !(i < 10 && i > 4)) ||
				(i+j == MBR && !(i < 10 && i > 4)) {
				grid[i][j] = &Slot{&t, byte('D')}
			} else {
				grid[i][j] = &Slot{&t, byte('n')}
			}
		}
	}
	//make the starting tile a double word
	grid[7][7] = &Slot{&t, byte('D')}
	//set the double letter tiles
	for _, p := range dub_letter_points {
		grid[p.i][p.j] = &Slot{&t, byte('d')}
		grid[p.j][p.i] = &Slot{&t, byte('d')}
		grid[MBR-p.j][MBR-p.i] = &Slot{&t, byte('d')}
		grid[MBR-p.i][MBR-p.j] = &Slot{&t, byte('d')}
	}
	//set the triple letter tiles
	for _, p := range trip_letter_points {
		grid[p.i][p.j] = &Slot{&t, byte('t')}
		grid[p.j][p.i] = &Slot{&t, byte('t')}
		grid[MBR-p.i][MBR-p.j] = &Slot{&t, byte('t')}
		grid[MBR-p.i][p.j] = &Slot{&t, byte('t')}
		grid[p.i][MBR-p.j] = &Slot{&t, byte('t')}
		grid[MBR-p.j][p.i] = &Slot{&t, byte('t')}
		grid[p.j][MBR-p.j] = &Slot{&t, byte('t')}
	}
	return &Board{grid}
}

func (board *Board) Print() {
	for x := range board.grid {
		for y := range board.grid[x] {
			fmt.Print(string(board.grid[x][y].modifier), "   ")
		}
		fmt.Print("\n\n")
	}
}
