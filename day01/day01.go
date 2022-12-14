package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	puzzle1()
	puzzle2()
}

func puzzle1() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mostCalories, elf := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			food, err := strconv.Atoi(line)
			_ = err
			elf += food
		} else {
			if elf > mostCalories {
				mostCalories = elf
			}
			elf = 0
		}
	}
	fmt.Println(mostCalories)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func puzzle2() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elf := 0
	var calories []int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			food, err := strconv.Atoi(line)
			_ = err
			elf += food
		} else {
			calories = append(calories, elf)
			elf = 0
		}
	}
	sort.Ints(calories)
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	fmt.Println(calories[0] + calories[1] + calories[2])
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
