package main

import (
	"io"
	"io/ioutil"
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"os"
)

const illegal_chars = " -'"
const caps = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func beginsWithCap(word string) bool {
	return strings.ContainsAny(word[0:1], caps)
}

func isLegalWord(word string) bool {
	return !beginsWithCap(word) && !strings.ContainsAny(word, illegal_chars)
}

func readAndParseLines() ([]string, error) {
	var lines []string
	//running := true
	file, err := ioutil.ReadFile("assets/words.dict")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return lines, err
			}
		}
		if isLegalWord(line) {
			lines = append(lines, line)
		}
		if err != nil && err != io.EOF {
			return lines, err
		}
	}
	return lines, err
}

func writeParsedDict(filename string, lines []string) {
	file, err := os.Create("assets/" + filename)
	if (err != nil) {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line)
		if (err != nil) {
			panic(err)
		}
	}
	writer.Flush()
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
	lines, err := readAndParseLines()
	if (err != nil) {
		panic(err)
	}
	writeParsedDict("test1.dict", lines)
}
