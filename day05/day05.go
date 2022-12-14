package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() string {
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element
}

func main() {
	puzzle1()
	puzzle2()
}

func initStacks() [9]Stack {
	var stacks [9]Stack
	stacks[0].Push("T")
	stacks[0].Push("D")
	stacks[0].Push("W")
	stacks[0].Push("Z")
	stacks[0].Push("V")
	stacks[0].Push("P")
	stacks[1].Push("L")
	stacks[1].Push("S")
	stacks[1].Push("W")
	stacks[1].Push("V")
	stacks[1].Push("F")
	stacks[1].Push("J")
	stacks[1].Push("D")
	stacks[2].Push("Z")
	stacks[2].Push("M")
	stacks[2].Push("L")
	stacks[2].Push("S")
	stacks[2].Push("V")
	stacks[2].Push("T")
	stacks[2].Push("B")
	stacks[2].Push("H")
	stacks[3].Push("R")
	stacks[3].Push("S")
	stacks[3].Push("J")
	stacks[4].Push("C")
	stacks[4].Push("Z")
	stacks[4].Push("B")
	stacks[4].Push("G")
	stacks[4].Push("F")
	stacks[4].Push("M")
	stacks[4].Push("L")
	stacks[4].Push("W")
	stacks[5].Push("Q")
	stacks[5].Push("W")
	stacks[5].Push("V")
	stacks[5].Push("H")
	stacks[5].Push("Z")
	stacks[5].Push("R")
	stacks[5].Push("G")
	stacks[5].Push("B")
	stacks[6].Push("V")
	stacks[6].Push("J")
	stacks[6].Push("P")
	stacks[6].Push("C")
	stacks[6].Push("B")
	stacks[6].Push("D")
	stacks[6].Push("N")
	stacks[7].Push("P")
	stacks[7].Push("T")
	stacks[7].Push("B")
	stacks[7].Push("Q")
	stacks[8].Push("H")
	stacks[8].Push("G")
	stacks[8].Push("Z")
	stacks[8].Push("R")
	stacks[8].Push("C")
	return stacks
}

func puzzle1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stacks := initStacks()
	re := regexp.MustCompile("[0-9]+")

	scanner := bufio.NewScanner(file)

	for i := 0; i <= 9; i++ {
		scanner.Scan()
	}

	for scanner.Scan() {
		line := scanner.Text()
		result := re.FindAllString(line, -1)
		amount, _ := strconv.Atoi(result[0])
		from, _ := strconv.Atoi(result[1])
		to, _ := strconv.Atoi(result[2])
		for i := 1; i <= amount; i++ {
			stacks[to-1].Push(stacks[from-1].Pop())
		}
	}
	output := ""
	for i := 0; i < len(stacks); i++ {
		output = output + stacks[i].Pop()
	}
	fmt.Println(output)
}

func puzzle2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stacks := initStacks()
	re := regexp.MustCompile("[0-9]+")

	scanner := bufio.NewScanner(file)

	for i := 0; i <= 9; i++ {
		scanner.Scan()
	}

	for scanner.Scan() {
		line := scanner.Text()
		result := re.FindAllString(line, -1)
		amount, _ := strconv.Atoi(result[0])
		from, _ := strconv.Atoi(result[1])
		to, _ := strconv.Atoi(result[2])
		var tmp []string
		for i := 1; i <= amount; i++ {
			tmp = append(tmp, stacks[from-1].Pop())
		}
		for i := len(tmp) - 1; i >= 0; i-- {
			stacks[to-1].Push(tmp[i])
		}
	}
	output := ""
	for i := 0; i < len(stacks); i++ {
		output = output + stacks[i].Pop()
	}
	fmt.Println(output)
}
