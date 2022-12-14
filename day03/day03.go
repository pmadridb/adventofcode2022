package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
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
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstHalf := strings.Split(line[0:len(line)/2], "")
		secondHalf := strings.Split(line[len(line)/2:len(line)], "")
	out:
		for i1 := 0; i1 < len(firstHalf); i1++ {
			for i2 := 0; i2 < len(secondHalf); i2++ {
				if firstHalf[i1] == secondHalf[i2] {
					totalScore += strings.Index(Priority, firstHalf[i1]) + 1
					break out
				}
			}
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
		line1 := strings.Split(scanner.Text(), "")
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		for i1 := 0; i1 < len(line1); i1++ {
			if strings.Index(line2, line1[i1]) >= 0 && strings.Index(line3, line1[i1]) >= 0 {
				totalScore += strings.Index(Priority, line1[i1]) + 1
				break
			}
		}
	}
	fmt.Println(totalScore)
}
