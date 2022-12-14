package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	var trees []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trees = append(trees, line)
	}
	visible := 0
	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[0])-1; j++ {
			if isVisible(trees, i, j) {
				visible++
			}
		}
	}
	fmt.Println(visible + (len(trees)*4 - 4))
}

func isVisible(trees []string, i int, j int) bool {
	visible := true
	tree, _ := strconv.Atoi(string(trees[i][j]))
	for l := i - 1; l >= 0; l-- {
		currentTree, _ := strconv.Atoi(string(trees[l][j]))
		if currentTree >= tree {
			visible = false
			break
		}
	}
	if !visible {
		visible = true
		for r := i + 1; r < len(trees); r++ {
			currentTree, _ := strconv.Atoi(string(trees[r][j]))
			if currentTree >= tree {
				visible = false
				break
			}
		}
	}
	if !visible {
		visible = true
		for n := j + 1; n < len(trees); n++ {
			currentTree, _ := strconv.Atoi(string(trees[i][n]))
			if currentTree >= tree {
				visible = false
				break
			}
		}
	}
	if !visible {
		visible = true
		for s := j - 1; s >= 0; s-- {
			currentTree, _ := strconv.Atoi(string(trees[i][s]))
			if currentTree >= tree {
				visible = false
				break
			}
		}
	}
	return visible
}

func puzzle2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var trees []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trees = append(trees, line)
	}
	score := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[0]); j++ {
			currentScore := getScore(trees, i, j)
			if currentScore > score {
				score = currentScore
			}
		}
	}
	fmt.Println(score)

}
func getScore(trees []string, i int, j int) int {
	score1, score2, score3, score4 := 0, 0, 0, 0
	tree, _ := strconv.Atoi(string(trees[i][j]))
	for l := i - 1; l >= 0; l-- {
		currentTree, _ := strconv.Atoi(string(trees[l][j]))
		score1++
		if currentTree >= tree {
			break
		}
	}
	for r := i + 1; r < len(trees); r++ {
		currentTree, _ := strconv.Atoi(string(trees[r][j]))
		score2++
		if currentTree >= tree {
			break
		}
	}
	for n := j + 1; n < len(trees); n++ {
		currentTree, _ := strconv.Atoi(string(trees[i][n]))
		score3++
		if currentTree >= tree {
			break
		}
	}
	for s := j - 1; s >= 0; s-- {
		currentTree, _ := strconv.Atoi(string(trees[i][s]))
		score4++
		if currentTree >= tree {
			break
		}
	}
	return score1 * score2 * score3 * score4
}
