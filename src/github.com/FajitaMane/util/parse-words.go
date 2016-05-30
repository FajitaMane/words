package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

const illegal_chars = " -'"
const caps = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func beginsWithCap(word string) bool {
	return strings.ContainsAny(word[0:1], caps)
}

func isLegalWord(word string) bool {
	return !beginsWithCap(word) && !strings.ContainsAny(word, illegal_chars)
}

func readLines() ([]string, error) {
	var lines []string
	running := true
	file, err := ioutil.ReadFile("assets/words.dict")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString("\n")
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return lines, err
			}
		}
		lines = append(lines, line)
		if err != nil && err != io.EOF {
			return lines, err
		}
	}
	return lines, err
}

func test() {
	fmt.Println("this should be false: ", beginsWithCap("john"))
	fmt.Println("this should be true: ", beginsWithCap("John:"))
	fmt.Println("this should be false", isLegalWord("ain't"))
	fmt.Println("this should be true", isLegalWord("ubiquitous"))
	fmt.Println("this should be false: ", isLegalWord("Ternary"))
	fmt.Println("this should be false: ", isLegalWord("no go"))
	fmt.Println("this should be true: ", isLegalWord("a"))
	fmt.Println("this should be false: ", isLegalWord("No-go"))
	fmt.Println("this should be true: ", isLegalWord("truE"))
}

func main() {
	test()
}
