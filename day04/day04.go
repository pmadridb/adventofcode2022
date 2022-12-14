package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		split := strings.Split(line, ",")
		left := strings.Split(split[0], "-")
		right := strings.Split(split[1], "-")
		leftmin, err := strconv.Atoi(left[0])
		leftmax, err := strconv.Atoi(left[1])
		rightmin, err := strconv.Atoi(right[0])
		rightmax, err := strconv.Atoi(right[1])
		_ = err
		if (leftmin <= rightmin && leftmax >= rightmax) || (leftmin >= rightmin && leftmax <= rightmax) {
			totalScore++
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
		split := strings.Split(line, ",")
		left := strings.Split(split[0], "-")
		right := strings.Split(split[1], "-")
		leftmin, err := strconv.Atoi(left[0])
		leftmax, err := strconv.Atoi(left[1])
		rightmin, err := strconv.Atoi(right[0])
		rightmax, err := strconv.Atoi(right[1])
		_ = err
		if (leftmin >= rightmin && leftmin <= rightmax) || (leftmax >= rightmin && leftmax <= rightmax) || (leftmin <= rightmin && leftmax >= rightmax) || (leftmin >= rightmin && leftmax <= rightmax) {
			totalScore++
		}
	}
	fmt.Println(totalScore)
}
