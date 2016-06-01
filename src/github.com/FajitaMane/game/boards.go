package main

import (
	"fmt"
)

type Board struct {
	grid [][]*Slot
}

//the NewBoard method returns a new instance of the standard Board
func NewBoard() *Board {
	//empty tile
	t := Tile{uint8(156)}
	//instantiate the empty grid
	grid := make([][]*Slot, 15)
	for i := range grid {
		grid[i] = make([]*Slot, 15)
		for j := range grid[i] {
			if i%14 == 0 && j%14 == 0 {
				grid[i][j] = &Slot{&t, byte('T')}
			} else {
				grid[i][j] = &Slot{&t, byte('n')}
			}
		}
	}
	return &Board{grid}
}

func (board *Board) Print() {
	for x := range board.grid {
		for y := range board.grid[x] {
			fmt.Print(board.grid[x][y].modifier, " ")
		}
		fmt.Print("\n")
	}
}
