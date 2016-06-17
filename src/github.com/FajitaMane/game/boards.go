package game

import (
	"fmt"
	_ "strconv"
)

type Board struct {
	grid [][]*Slot
}

//helper for defining board points
//TODO change this shit to x and y
type Point struct {
	x int
	y int
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
		grid[p.x][p.y] = &Slot{&t, byte('d')}
		grid[p.y][p.x] = &Slot{&t, byte('d')}
		grid[MBR-p.y][MBR-p.x] = &Slot{&t, byte('d')}
		grid[MBR-p.x][MBR-p.y] = &Slot{&t, byte('d')}
	}
	//set the triple letter tiles
	for _, p := range trip_letter_points {
		grid[p.x][p.y] = &Slot{&t, byte('t')}
		grid[p.y][p.x] = &Slot{&t, byte('t')}
		grid[MBR-p.x][MBR-p.y] = &Slot{&t, byte('t')}
		grid[MBR-p.x][p.y] = &Slot{&t, byte('t')}
		grid[p.x][MBR-p.y] = &Slot{&t, byte('t')}
		grid[MBR-p.y][p.x] = &Slot{&t, byte('t')}
		grid[p.y][MBR-p.y] = &Slot{&t, byte('t')}
	}
	return &Board{grid}
}

func (board *Board) IsLegalPlay(play *Play) bool {
	//all slots must contain non-nil tile characters
	for i := range play.word {
		if play.word[i].tile.char == 0 {
			return false
		}
	}
	//make sure the word won't exceed the board size
	if play.hor && play.index.y+len(play.word) > MBR {
		return false
	}
	if !play.hor && play.index.x+len(play.word) > MBR {
		return false
	}
	/*
		//make sure it doesn't intersect any other tiles
		for x := 0; x < len(board.grid); x++ {
			for y := 0; y < len(board.grid[x]); y++ {
				if board.grid[x][y].tile == nil {
					return false
				}
			}
		}
	*/
	return true
}

func (board *Board) ValueOfPlay(play *Play) uint16 {
	var val uint16
	for i := range play.word {
		val += play.word[i].Value()
	}
	for i := range play.word {
		if play.word[i].modifier == byte('D') {
			val = val * 2
		}
		if play.word[i].modifier == byte('T') {
			val = val * 3
		}
	}
	return val
}

func (board *Board) Print() {
	//offset the first row
	fmt.Print("   ")
	for y := range board.grid {
		if y < 10 {
			fmt.Print(y, "   ")
		} else {
			fmt.Print(y, "  ")
		}
	}
	fmt.Print("\n\n")
	for x := range board.grid {
		if x < 10 {
			fmt.Print(x, "  ")
		} else {
			fmt.Print(x, " ")
		}
		for y := range board.grid[x] {
			fmt.Print(string(board.grid[x][y].modifier), "   ")
		}
		if x < 10 {
			fmt.Print(" ", x)
		} else {
			fmt.Print(x)
		}
		fmt.Print("\n\n")
	}
	//offset the last row
	fmt.Print("   ")
	for y := range board.grid {
		if y < 10 {
			fmt.Print(y, "   ")
		} else {
			fmt.Print(y, "  ")
		}
	}
	fmt.Print("\n")
}

//TODO this could use be more consise
func (board *Board) PrintTiles() {
	//offset the first row
	fmt.Print("   ")
	for y := range board.grid {
		if y < 10 {
			fmt.Print(y, "   ")
		} else {
			fmt.Print(y, "  ")
		}
	}
	fmt.Print("\n\n")
	for x := range board.grid {
		if x < 10 {
			fmt.Print(x, "  ")
		} else {
			fmt.Print(x, " ")
		}
		for y := range board.grid[x] {
			fmt.Print(string(board.grid[x][y].tile.char), "   ")
		}
		if x < 10 {
			fmt.Print(" ", x)
		} else {
			fmt.Print(x)
		}
		fmt.Print("\n\n")
	}
	//offset the last row
	fmt.Print("   ")
	for y := range board.grid {
		if y < 10 {
			fmt.Print(y, "   ")
		} else {
			fmt.Print(y, "  ")
		}
	}
	fmt.Print("\n")
}
