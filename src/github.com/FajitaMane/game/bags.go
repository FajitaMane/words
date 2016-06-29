package game

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

type Bag struct {
	tiles *list.List
}

var alphabet string

//called by Game before getting a bag
func initBag() {
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
}

//TODO change this to singleton pattern?
func NewBag() *Bag {
	bag_list := list.New()
	//iterate through every letter in the alphabet
	for x, _ := range alphabet {
		letter := alphabet[x : x+1][0]
		fmt.Println("adding ", string(letter))
		val_tile := Tile{letter}
		val := val_tile.Frequency()
		for i := 0; uint8(i) < val; i++ {
			bag_list.PushFront(&Tile{letter})
			//fmt.Println("adding tile #", bag_list.Len(),
			//" to bag_list, letter=", letter)

		}
	}
	var bag Bag
	bag = Bag{bag_list}
	fmt.Println(bag_list.Len())
	return &bag
}

//seed a new random number and return a letter at index n
func (bag *Bag) Draw() Tile {
	fmt.Println("bag size = ", bag.Size())
	if bag.Size() == 0 {
		panic("bag is empty")
	}
	rand.Seed(time.Now().UTC().UnixNano())
	index := rand.Intn(bag.tiles.Len())
	i := 0
	for e := bag.tiles.Front(); e != nil; e = e.Next() {
		if i == index {
			tile := Tile{e.Value.(Tile).char}
			bag.tiles.Remove(e)
			return tile
		}
		i++
	}
	panic("couldn't find index")
}

func (bag *Bag) Size() int {
	return bag.tiles.Len()
}
