package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		caracters := strings.Split(scanner.Text(), "")
		lines = append(lines, caracters)
	}
	return lines, scanner.Err()
}

type path struct {
	X, Y int
}

func main() {
	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1 := processaMapa(linies, 3, 1)

	fmt.Println("Part 1: ", correctes1)

	var camins = []path{
		path{1, 1},
		path{3, 1},
		path{5, 1},
		path{7, 1},
		path{1, 2},
	}

	correctes2 := 1
	for _, cami := range camins {
		correctes2 = correctes2 * processaMapa(linies, cami.X, cami.Y)
	}

	fmt.Println("Part 2: ", correctes2)
}

func processaMapa(linies [][]string, stepx int, stepy int) int {
	trees := 0
	x := 0
	y := 0

	height := len(linies)
	width := len(linies[0])

	for y < height {
		y = y + stepy
		if y < height {
			x = (x + stepx) % width
			if linies[y][x] == "#" {
				trees = trees + 1
			}
		}
	}

	return trees
}
