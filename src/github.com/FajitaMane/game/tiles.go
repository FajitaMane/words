package main

type Tile struct {
	char uint8
}

//the value returned is equal to the letter's frequency * 10
//the in-game value of this Tile is the uint16 / 10
func (t *Tile) Value() uint16 {
	//initialize the map
	var val_map map[uint8]uint16
	val_map = make(map[uint8]uint16)
	//set keys and values
	val_map["A"[0]] = 7
	val_map["B"[0]] = 76
	val_map["C"[0]] = 94
	val_map["D"[0]] = 93
	val_map["E"[0]] = 109
	val_map["F"[0]] = 88
	val_map["G"[0]] = 109
	val_map["H"[0]] = 51
	val_map["I"[0]] = 60
	val_map["J"[0]] = 117
	val_map["K"[0]] = 117
	val_map["L"[0]] = 89
	val_map["M"[0]] = 80
	val_map["N"[0]] = 95
	val_map["O"[0]] = 60
	val_map["P"[0]] = 86
	val_map["Q"[0]] = 119
	val_map["R"[0]] = 107
	val_map["S"[0]] = 45
	val_map["T"[0]] = 50
	val_map["U"[0]] = 95
	val_map["V"[0]] = 112
	val_map["W"[0]] = 55
	val_map["X"[0]] = 116
	val_map["Y"[0]] = 84
	val_map["Z"[0]] = 110

	return val_map[t.char]
}

//returns the number of tiles in the Bag when the game starts
func (t *Tile) Frequency() uint8 {
	//initialize the map
	var val_map map[uint8]uint8
	val_map = make(map[uint8]uint8)
	//set the keys and values
	val_map["A"[0]] = 9
	val_map["B"[0]] = 2
	val_map["C"[0]] = 2
	val_map["D"[0]] = 4
	val_map["E"[0]] = 12
	val_map["F"[0]] = 2
	val_map["G"[0]] = 3
	val_map["H"[0]] = 2
	val_map["I"[0]] = 9
	val_map["J"[0]] = 1
	val_map["K"[0]] = 1
	val_map["L"[0]] = 4
	val_map["M"[0]] = 2
	val_map["N"[0]] = 6
	val_map["O"[0]] = 8
	val_map["P"[0]] = 2
	val_map["Q"[0]] = 1
	val_map["R"[0]] = 6
	val_map["S"[0]] = 4
	val_map["T"[0]] = 6
	val_map["U"[0]] = 4
	val_map["V"[0]] = 2
	val_map["W"[0]] = 2
	val_map["X"[0]] = 1
	val_map["Y"[0]] = 2
	val_map["Z"[0]] = 1

	return val_map[t.char]
}
