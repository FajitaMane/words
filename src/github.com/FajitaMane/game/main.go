package main

import (
	"fmt"
)

func main() {
	var bag *Bag
	initBag()
	bag = NewBag()
	fmt.Println("total of ", bag.tiles.Len(), " tiles")
	var i int
	for i = bag.tiles.Len(); i > 0; i = bag.tiles.Len() {
		fmt.Print("drew tile ")
		tile_char := bag.Draw().char
		fmt.Printf(string(tile_char))
		fmt.Println(" tile #", 101-i, " with ", bag.tiles.Len(), " remaining.")
	}
}
