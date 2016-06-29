package game

import (
	"fmt"
	"strconv"

	"github.com/FajitaMane/util"
	"github.com/fatih/color"
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
	trip_letter_points := make([]Point, 2)
	trip_letter_points[0] = Point{1, 5}
	trip_letter_points[1] = Point{5, 5}
	//empty tile
	t := Tile{uint8(0)}
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
		grid[MBR-p.y][MBR-p.x] = &Slot{&t, byte('t')}
		grid[p.y][MBR-p.x] = &Slot{&t, byte('t')}
	}
	return &Board{grid}
}

func (board *Board) GetPlaySlots(point Point, hor bool, length int) []*Slot {
	slots := make([]*Slot, length)
	for i := range slots {
		if hor {
			slots[i] = board.grid[point.x+i][point.y]
		} else {
			slots[i] = board.grid[point.x][point.y+i]
		}
	}
	return slots
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
	//finally, check if the word is in the dictionary
	b, err := util.IsWordInDict(play.Word())
	if err != nil {
		panic(err)
	}
	if !b {
		return false
	}
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

//outputs a representation of the board's tiles to the terminal
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
//prints out a representation of the board's tiles to the terminal
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
}

const M = 3 //margin of three for each column
//prints out the representation of the board that the client sees
func (board *Board) PrintF() {
	//offset the first row
	var first_row string
	for y := range board.grid {
		first_row += strconv.Itoa(y)
		if y < 10 {
			first_row += "   "
		} else {
			first_row += "  "
		}
	}
	color.White(first_row)
	fmt.Print("\n")
	//set the lines between the top and bottom rows
	for y := range board.grid {
		color.Set(color.FgWhite)
		fmt.Print(y + 1)
		if y+1 < 10 {
			fmt.Print("   ")
		} else {
			fmt.Print("  ")
		}
		for x := range board.grid {
			board.PrintSlotF(board.grid[y][x])
			if x < 10 {
				fmt.Print("   ")
			} else {
				fmt.Print("  ")
			}
		}
		color.Set(color.FgWhite)
		if y+1 < 10 {
			fmt.Print("  ")
		} else {
			fmt.Print(" ")
		}
		fmt.Print(y + 1)
		fmt.Print("\n\n")
	}
	//offset the last row
	var last_row string
	for y := range board.grid {
		last_row += strconv.Itoa(y)
		if y < 10 {
			last_row += "   "
		} else {
			last_row += "  "
		}
	}
	color.White(last_row)
}

func (board *Board) PrintSlotF(slot *Slot) {
	if slot.tile.char != 0 {
		color.Set(color.FgBlue)
		fmt.Print(slot.tile.char)
	} else {
		if slot.modifier != byte('n') {
			if slot.modifier == byte('t') {
				color.Set(color.FgGreen)
				fmt.Print("t")
			}
			if slot.modifier == byte('d') {
				color.Set(color.FgCyan)
				fmt.Print("d")
			}
			if slot.modifier == byte('T') {
				color.Set(color.FgYellow)
				fmt.Print("T")
			}
			if slot.modifier == byte('D') {
				color.Set(color.FgRed)
				fmt.Print("D")
			}
		} else {
			fmt.Print(" ")
		}
	}
}
