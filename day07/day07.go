package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	name     string
	parent   *Node
	children []*Node
	size     int
}

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
	rootDir := &Node{name: "/", size: 0}
	currentDir := rootDir
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$ cd") {
			dir := strings.Split(line, " ")[2]
			if dir == ".." {
				currentDir = currentDir.parent
			} else {
				currentDir = findDir(currentDir, dir)
			}
		} else if strings.HasPrefix(line, "dir") {
			newDir := Node{name: strings.Split(line, " ")[1], parent: currentDir, size: 0}
			currentDir.children = append(currentDir.children, &newDir)
		} else if _, err := strconv.Atoi(strings.Split(line, " ")[0]); err == nil {
			size, _ := strconv.Atoi(strings.Split(line, " ")[0])
			currentDir.size += size
		}
	}
	sum := 0
	var sizes []int
	sizesp := &sizes
	getSizes(rootDir, sizesp)
	for _, s := range sizes {
		if s < 100000 {
			sum += s
		}
	}
	fmt.Println(sum)
}
func getSizes(node *Node, sizes *[]int) int {
	sum := node.size
	for i := 0; i < len(node.children); i++ {
		sum += getSizes(node.children[i], sizes)
	}
	*sizes = append(*sizes, sum)
	return sum
}
func findDir(dir *Node, name string) *Node {
	for i := 0; i < len(dir.children); i++ {
		if dir.children[i].name == name {
			return dir.children[i]
		}
	}
	return nil
}

func puzzle2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rootDir := &Node{name: "/", size: 0}
	currentDir := rootDir
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$ cd") {
			dir := strings.Split(line, " ")[2]
			if dir == ".." {
				currentDir = currentDir.parent
			} else {
				currentDir = findDir(currentDir, dir)
			}
		} else if strings.HasPrefix(line, "dir") {
			newDir := Node{name: strings.Split(line, " ")[1], parent: currentDir, size: 0}
			currentDir.children = append(currentDir.children, &newDir)
		} else if _, err := strconv.Atoi(strings.Split(line, " ")[0]); err == nil {
			size, _ := strconv.Atoi(strings.Split(line, " ")[0])
			currentDir.size += size
		}
	}
	var sizes []int
	sizesp := &sizes
	getSizes(rootDir, sizesp)
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] < sizes[j]
	})
	totalSize := sizes[len(sizes)-1]
	unused := 70000000 - totalSize
	required := 30000000 - unused
	index := 0
	for _, v := range sizes {
		if v >= required {
			break
		}
		index++
	}
	fmt.Println(sizes[index])
}
