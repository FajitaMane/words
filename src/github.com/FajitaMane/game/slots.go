package main

import (
	"fmt"
)

/*
Here's what the modifier bytes represent:
C (0x43) Center Tile (also double word)
D (0x44) Double Word
T (0x54) Triple Word
d (0x64) Double Tile
t (0x74) Triple Tile
n (0x6#) No Modifier
*/

type Slot struct {
	tile *Tile
	modifier byte
}

func (slot *Slot) Print() {
	fmt.Println("Tile ", slot.tile.char, " value of ", slot.tile.Value()) 
}