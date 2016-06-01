package main

import (
	"container/list"
	_ "fmt"
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
		var val uint8
		letter := alphabet[x : x+1][0]
		tile := Tile{letter}
		val = tile.Frequency()
		for i := 0; uint8(i) < val; i++ {
			if i == 0 {
				bag_list.PushFront(tile)
			} else {
				new_tile := Tile{letter}
				bag_list.PushFront(new_tile)
			}
			//fmt.Println("adding tile #", bag_list.Len(),
			//" to bag_list, letter=", letter)

		}
	}
	var bag Bag
	bag = Bag{bag_list}
	return &bag
}

//seed a new random number and return a letter at index n
func (bag *Bag) Draw() Tile {
	if bag.tiles.Len() == 0 {
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
