package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Paper    = 2
	Scissors = 3
	Rock     = 1
)

func getResult(value1 int, value2 int) int {
	if value1 == Paper && value2 == Scissors {
		return 1
	} else if value1 == Paper && value2 == Rock {
		return -1
	} else if value1 == Scissors && value2 == Paper {
		return -1
	} else if value1 == Scissors && value2 == Rock {
		return 1
	} else if value1 == Rock && value2 == Paper {
		return 1
	} else if value1 == Rock && value2 == Scissors {
		return -1
	}
	return 0
}

func main() {
	puzzle1()
	puzzle2()
}

func mapMove(value string) int {
	if value == "A" || value == "X" {
		return Rock
	} else if value == "B" || value == "Y" {
		return Paper
	}
	return Scissors
}

func puzzle1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		value1 := mapMove(split[0])
		value2 := mapMove(split[1])
		result := getResult(value1, value2)
		if result == -1 {
			totalScore += value2
		} else if result == 0 {
			totalScore = totalScore + 3 + value2
		} else if result == 1 {
			totalScore = totalScore + 6 + value2
		}
	}
	fmt.Println(totalScore)
}

func puzzle2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		value1 := mapMove(split[0])
		string2 := split[1]
		value2 := value1
		if string2 == "X" {
			value2 = getLoser(value1)
		} else if string2 == "Z" {
			value2 = getWinner(value1)
		}

		result := getResult(value1, value2)
		if result == -1 {
			totalScore += value2
		} else if result == 0 {
			totalScore = totalScore + 3 + value2
		} else if result == 1 {
			totalScore = totalScore + 6 + value2
		}
	}
	fmt.Println(totalScore)
}

func getWinner(value int) int {
	if value == Rock {
		return Paper
	}
	if value == Paper {
		return Scissors
	}
	return Rock
}
func getLoser(value int) int {
	if value == Rock {
		return Scissors
	}
	if value == Paper {
		return Rock
	}
	return Paper
}
