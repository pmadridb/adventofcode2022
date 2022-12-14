package main

import (
	"fmt"
)

type Monkey struct {
	items       []int
	operation   string
	operand     int
	divisible   int
	falseResult int
	trueResult  int
}

func main() {
	puzzle1()
	puzzle2()
}

func puzzle1() {
	var monkeys [8]Monkey
	counter := [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	monkeys[0].items = []int{50, 70, 89, 75, 66, 66}
	monkeys[0].operation = "por"
	monkeys[0].operand = 5
	monkeys[0].divisible = 2
	monkeys[0].trueResult = 2
	monkeys[0].falseResult = 1

	monkeys[1].items = []int{85}
	monkeys[1].operation = "por"
	monkeys[1].operand = 0
	monkeys[1].divisible = 7
	monkeys[1].trueResult = 3
	monkeys[1].falseResult = 6

	monkeys[2].items = []int{66, 51, 71, 76, 58, 55, 58, 60}
	monkeys[2].operation = "mas"
	monkeys[2].operand = 1
	monkeys[2].divisible = 13
	monkeys[2].trueResult = 1
	monkeys[2].falseResult = 3

	monkeys[3].items = []int{79, 52, 55, 51}
	monkeys[3].operation = "mas"
	monkeys[3].operand = 6
	monkeys[3].divisible = 3
	monkeys[3].trueResult = 6
	monkeys[3].falseResult = 4

	monkeys[4].items = []int{69, 92}
	monkeys[4].operation = "por"
	monkeys[4].operand = 17
	monkeys[4].divisible = 19
	monkeys[4].trueResult = 7
	monkeys[4].falseResult = 5

	monkeys[5].items = []int{71, 76, 73, 98, 67, 79, 99}
	monkeys[5].operation = "mas"
	monkeys[5].operand = 8
	monkeys[5].divisible = 5
	monkeys[5].trueResult = 0
	monkeys[5].falseResult = 2

	monkeys[6].items = []int{82, 76, 69, 69, 57}
	monkeys[6].operation = "mas"
	monkeys[6].operand = 7
	monkeys[6].divisible = 11
	monkeys[6].trueResult = 7
	monkeys[6].falseResult = 4

	monkeys[7].items = []int{65, 79, 86}
	monkeys[7].operation = "mas"
	monkeys[7].operand = 5
	monkeys[7].divisible = 17
	monkeys[7].trueResult = 5
	monkeys[7].falseResult = 0

	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			cop := make([]int, len(monkeys[j].items))
			copy(cop, monkeys[j].items)
			monkeys[j].items = monkeys[j].items[:0]
			for k := 0; k < len(cop); k++ {
				res0 := 0
				if monkeys[j].operation == "mas" {
					if monkeys[j].operand == 0 {
						res0 = cop[k] + cop[k]
					} else {
						res0 = cop[k] + monkeys[j].operand
					}
				} else {
					if monkeys[j].operand == 0 {
						res0 = cop[k] * cop[k]
					} else {
						res0 = cop[k] * monkeys[j].operand
					}
				}
				res1 := res0 / 3
				counter[j]++
				if res1%monkeys[j].divisible == 0 {
					monkeys[monkeys[j].trueResult].items = append(monkeys[monkeys[j].trueResult].items, int(res1))
				} else {
					monkeys[monkeys[j].falseResult].items = append(monkeys[monkeys[j].falseResult].items, int(res1))
				}
			}
		}

	}

	fmt.Println(counter)
}

func puzzle2() {
	var monkeys [8]Monkey
	counter := [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	monkeys[0].items = []int{50, 70, 89, 75, 66, 66}
	monkeys[0].operation = "por"
	monkeys[0].operand = 5
	monkeys[0].divisible = 2
	monkeys[0].trueResult = 2
	monkeys[0].falseResult = 1

	monkeys[1].items = []int{85}
	monkeys[1].operation = "por"
	monkeys[1].operand = 0
	monkeys[1].divisible = 7
	monkeys[1].trueResult = 3
	monkeys[1].falseResult = 6

	monkeys[2].items = []int{66, 51, 71, 76, 58, 55, 58, 60}
	monkeys[2].operation = "mas"
	monkeys[2].operand = 1
	monkeys[2].divisible = 13
	monkeys[2].trueResult = 1
	monkeys[2].falseResult = 3

	monkeys[3].items = []int{79, 52, 55, 51}
	monkeys[3].operation = "mas"
	monkeys[3].operand = 6
	monkeys[3].divisible = 3
	monkeys[3].trueResult = 6
	monkeys[3].falseResult = 4

	monkeys[4].items = []int{69, 92}
	monkeys[4].operation = "por"
	monkeys[4].operand = 17
	monkeys[4].divisible = 19
	monkeys[4].trueResult = 7
	monkeys[4].falseResult = 5

	monkeys[5].items = []int{71, 76, 73, 98, 67, 79, 99}
	monkeys[5].operation = "mas"
	monkeys[5].operand = 8
	monkeys[5].divisible = 5
	monkeys[5].trueResult = 0
	monkeys[5].falseResult = 2

	monkeys[6].items = []int{82, 76, 69, 69, 57}
	monkeys[6].operation = "mas"
	monkeys[6].operand = 7
	monkeys[6].divisible = 11
	monkeys[6].trueResult = 7
	monkeys[6].falseResult = 4

	monkeys[7].items = []int{65, 79, 86}
	monkeys[7].operation = "mas"
	monkeys[7].operand = 5
	monkeys[7].divisible = 17
	monkeys[7].trueResult = 5
	monkeys[7].falseResult = 0

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			cop := make([]int, len(monkeys[j].items))
			copy(cop, monkeys[j].items)
			monkeys[j].items = monkeys[j].items[:0]
			for k := 0; k < len(cop); k++ {
				res0 := 0
				if monkeys[j].operation == "mas" {
					if monkeys[j].operand == 0 {
						res0 = cop[k] + cop[k]
					} else {
						res0 = cop[k] + monkeys[j].operand
					}
				} else {
					if monkeys[j].operand == 0 {
						res0 = cop[k] * cop[k]
					} else {
						res0 = cop[k] * monkeys[j].operand
					}
				}
				res1 := res0 % 9699690
				counter[j]++
				if res1%monkeys[j].divisible == 0 {
					monkeys[monkeys[j].trueResult].items = append(monkeys[monkeys[j].trueResult].items, int(res1))
				} else {
					monkeys[monkeys[j].falseResult].items = append(monkeys[monkeys[j].falseResult].items, int(res1))
				}
			}
		}

	}

	fmt.Println(counter)
}
