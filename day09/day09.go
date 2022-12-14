package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	tailX, tailY, headX, headY := 0, 0, 0, 0
	visited := make(map[string]bool)
	visited["0,0"] = true
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")[0]
		amount, _ := strconv.Atoi(strings.Split(line, " ")[1])
		if command == "R" {
			for i := 1; i <= amount; i++ {
				headX++
				if !isAdjacent(tailX, tailY, headX, headY) {
					tailX = headX - 1
					tailY = headY
					visited[strconv.Itoa(tailX)+","+strconv.Itoa(tailY)] = true
				}
			}
		} else if command == "L" {
			for i := 1; i <= amount; i++ {
				headX--
				if !isAdjacent(tailX, tailY, headX, headY) {
					tailX = headX + 1
					tailY = headY
					visited[strconv.Itoa(tailX)+","+strconv.Itoa(tailY)] = true
				}
			}
		} else if command == "U" {
			for i := 1; i <= amount; i++ {
				headY++
				if !isAdjacent(tailX, tailY, headX, headY) {
					tailX = headX
					tailY = headY - 1
					visited[strconv.Itoa(tailX)+","+strconv.Itoa(tailY)] = true
				}
			}
		} else { //D
			for i := 1; i <= amount; i++ {
				headY--
				if !isAdjacent(tailX, tailY, headX, headY) {
					tailX = headX
					tailY = headY + 1
					visited[strconv.Itoa(tailX)+","+strconv.Itoa(tailY)] = true
				}
			}
		}
	}
	fmt.Println(len(visited))
}

func isAdjacent(tailX int, tailY int, headX int, headY int) bool {
	if math.Abs(float64(headX)-float64(tailX)) > 1 || math.Abs(float64(headY)-float64(tailY)) > 1 {
		return false
	}
	return true
}

func puzzle2() {
	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	tailX := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	tailY := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	tailxp, tailyp := &tailX, &tailY

	visited := make(map[string]bool)
	visited["0,0"] = true
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")[0]
		amount, _ := strconv.Atoi(strings.Split(line, " ")[1])
		move(command, amount, tailxp, tailyp, visited)
	}
	fmt.Println(len(visited))
}
func FindSign(n int) int {
	if n == 0 {
		return 0
	} else if n > 0 {
		return 1
	} else {
		return -1
	}
}
func move(command string, amount int, tailX *[10]int, tailY *[10]int, visited map[string]bool) {
	if command == "R" {
		for a := 1; a <= amount; a++ {
			for i := 0; i < 9; i++ {
				if i == 0 {
					tailX[i]++
				}
				if !isAdjacent(tailX[i+1], tailY[i+1], tailX[i], tailY[i]) {
					offX := FindSign(tailX[i+1] - tailX[i])
					offY := FindSign(tailY[i+1] - tailY[i])
					tailX[i+1] += offX
					tailY[i+1] += offY
					if i == 8 {
						visited[strconv.Itoa(tailX[i+1])+","+strconv.Itoa(tailY[i+1])] = true
					}
				}
			}
		}
	} else if command == "L" {
		for a := 1; a <= amount; a++ {
			for i := 0; i < 9; i++ {
				if i == 0 {
					tailX[i]--
				}
				if !isAdjacent(tailX[i+1], tailY[i+1], tailX[i], tailY[i]) {
					offX := FindSign(tailX[i+1] - tailX[i])
					offY := FindSign(tailY[i+1] - tailY[i])
					tailX[i+1] += offX
					tailY[i+1] += offY
					if i == 8 {
						visited[strconv.Itoa(tailX[i+1])+","+strconv.Itoa(tailY[i+1])] = true
					}
				}
			}
		}
	} else if command == "U" {
		for a := 1; a <= amount; a++ {
			for i := 0; i < 9; i++ {
				if i == 0 {
					tailY[i]++
				}
				if !isAdjacent(tailX[i+1], tailY[i+1], tailX[i], tailY[i]) {
					offX := FindSign(tailX[i+1] - tailX[i])
					offY := FindSign(tailY[i+1] - tailY[i])
					tailX[i+1] += offX
					tailY[i+1] += offY
					if i == 8 {
						visited[strconv.Itoa(tailX[i+1])+","+strconv.Itoa(tailY[i+1])] = true
					}
				}
			}
		}
	} else { //D
		for a := 1; a <= amount; a++ {
			for i := 0; i < 9; i++ {
				if i == 0 {
					tailY[i]--
				}
				if !isAdjacent(tailX[i+1], tailY[i+1], tailX[i], tailY[i]) {
					offX := FindSign(tailX[i+1] - tailX[i])
					offY := FindSign(tailY[i+1] - tailY[i])
					tailX[i+1] += offX
					tailY[i+1] += offY
					if i == 8 {
						visited[strconv.Itoa(tailX[i+1])+","+strconv.Itoa(tailY[i+1])] = true
					}
				}
			}
		}
	}

}
