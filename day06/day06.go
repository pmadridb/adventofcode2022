package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	puzzle1()
	puzzle2()
}

func puzzle1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	i := 0
	for ; i < len(line); i++ {
		signal := line[i : i+4]
		if isUnique(signal) {
			break
		}
	}

	fmt.Println(i + 4)
}

func isUnique(signal string) bool {
	for i := 0; i < len(signal); i++ {
		for j := 0; j < len(signal); j++ {
			if i != j && signal[i] == signal[j] {
				return false
			}
		}
	}
	return true
}

func puzzle2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	i := 0
	for ; i < len(line); i++ {
		signal := line[i : i+14]
		if isUnique(signal) {
			break
		}
	}

	fmt.Println(i + 14)
}
