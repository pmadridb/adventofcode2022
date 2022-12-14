package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	x, cycle, strength := 1, 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")[0]
		if command != "noop" {
			amount, _ := strconv.Atoi(strings.Split(line, " ")[1])
			cycle++
			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				strength += cycle * x
			}
			cycle++
			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				strength += cycle * x
			}
			x += amount
		} else {
			cycle++
			continue
		}
	}
	fmt.Println(strength)
}

func puzzle2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	x, cycle := 2, 1
	line1, line2, line3, line4, line5, line6 := "", "", "", "", "", ""
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")[0]
		if command != "noop" {
			amount, _ := strconv.Atoi(strings.Split(line, " ")[1])
			if cycle <= 40 {
				if cycle >= x-1 && cycle <= x+1 {
					line1 = line1 + "#"
				} else {
					line1 = line1 + "."
				}
			} else if cycle <= 80 {
				if cycle-40 >= x-1 && cycle-40 <= x+1 {
					line2 = line2 + "#"
				} else {
					line2 = line2 + "."
				}
			} else if cycle <= 120 {
				if cycle-80 >= x-1 && cycle-80 <= x+1 {
					line3 = line3 + "#"
				} else {
					line3 = line3 + "."
				}
			} else if cycle <= 160 {
				if cycle-120 >= x-1 && cycle-120 <= x+1 {
					line4 = line4 + "#"
				} else {
					line4 = line4 + "."
				}
			} else if cycle <= 200 {
				if cycle-160 >= x-1 && cycle-160 <= x+1 {
					line5 = line5 + "#"
				} else {
					line5 = line5 + "."
				}
			} else {
				if cycle-200 >= x-1 && cycle-200 <= x+1 {
					line6 = line6 + "#"
				} else {
					line6 = line6 + "."
				}
			}
			cycle++
			if cycle <= 40 {
				if cycle >= x-1 && cycle <= x+1 {
					line1 = line1 + "#"
				} else {
					line1 = line1 + "."
				}
			} else if cycle <= 80 {
				if cycle-40 >= x-1 && cycle-40 <= x+1 {
					line2 = line2 + "#"
				} else {
					line2 = line2 + "."
				}
			} else if cycle <= 120 {
				if cycle-80 >= x-1 && cycle-80 <= x+1 {
					line3 = line3 + "#"
				} else {
					line3 = line3 + "."
				}
			} else if cycle <= 160 {
				if cycle-120 >= x-1 && cycle-120 <= x+1 {
					line4 = line4 + "#"
				} else {
					line4 = line4 + "."
				}
			} else if cycle <= 200 {
				if cycle-160 >= x-1 && cycle-160 <= x+1 {
					line5 = line5 + "#"
				} else {
					line5 = line5 + "."
				}
			} else {
				if cycle-200 >= x-1 && cycle-200 <= x+1 {
					line6 = line6 + "#"
				} else {
					line6 = line6 + "."
				}
			}
			cycle++
			x += amount
		} else {
			cycle++
			if cycle <= 40 {
				if cycle >= x-1 && cycle <= x+1 {
					line1 = line1 + "#"
				} else {
					line1 = line1 + "."
				}
			} else if cycle <= 80 {
				if cycle-40 >= x-1 && cycle-40 <= x+1 {
					line2 = line2 + "#"
				} else {
					line2 = line2 + "."
				}
			} else if cycle <= 120 {
				if cycle-80 >= x-1 && cycle-80 <= x+1 {
					line3 = line3 + "#"
				} else {
					line3 = line3 + "."
				}
			} else if cycle <= 160 {
				if cycle-120 >= x-1 && cycle-120 <= x+1 {
					line4 = line4 + "#"
				} else {
					line4 = line4 + "."
				}
			} else if cycle <= 200 {
				if cycle-160 >= x-1 && cycle-160 <= x+1 {
					line5 = line5 + "#"
				} else {
					line5 = line5 + "."
				}
			} else {
				if cycle-200 >= x-1 && cycle-200 <= x+1 {
					line6 = line6 + "#"
				} else {
					line6 = line6 + "."
				}
			}
			continue
		}
	}
	fmt.Println(line1)
	fmt.Println(line2)
	fmt.Println(line3)
	fmt.Println(line4)
	fmt.Println(line5)
	fmt.Println(line6)
}
