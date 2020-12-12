package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

type airplane struct {
	direccio direction
	posicio  direction
}

type move struct {
	action string
	value  int
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]move, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []move
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		value, _ := stringToInt(line[1:])
		lines = append(lines, move{string(line[0]), value})
	}
	return lines, scanner.Err()
}

func mou(actual direction, value int, dir direction) direction {
	actual.y += (dir.y * value)
	actual.x += (dir.x * value)
	return actual
}

type direction struct {
	x int
	y int
}

func step(actual airplane, moviment move, waypoint bool) airplane {
	directions := map[string]direction{"N": {0, -1}, "S": {0, 1}, "E": {1, 0}, "W": {-1, 0}}

	switch moviment.action {
	case "N", "S", "E", "W":
		if waypoint {
			actual.direccio = mou(actual.direccio, moviment.value, directions[moviment.action])
		} else {
			actual.posicio = mou(actual.posicio, moviment.value, directions[moviment.action])
		}
	case "F":

		actual.posicio = mou(actual.posicio, moviment.value, actual.direccio)

	case "L":
		for i := 0; i < moviment.value/90; i++ {
			nova := actual.direccio.x
			actual.direccio.x = actual.direccio.y
			actual.direccio.y = -nova
		}
	case "R":
		for i := 0; i < moviment.value/90; i++ {
			nova := actual.direccio.x
			actual.direccio.x = -actual.direccio.y
			actual.direccio.y = nova
		}

	default:
		panic("Horror")
	}

	return actual
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func goAt(moves []move, actual airplane, waypoint bool) int {

	for _, mou := range moves {
		actual = step(actual, mou, waypoint)
	}
	return abs(actual.posicio.x) + abs(actual.posicio.y)
}

func main() {

	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	actual := airplane{
		direccio: direction{1, 0},
		posicio:  direction{0, 0},
	}

	distance1 := goAt(linies, actual, false)

	fmt.Println("Part 1: ", distance1)

	waypoint := airplane{
		direccio: direction{10, -1},
		posicio:  direction{0, 0},
	}

	distance2 := goAt(linies, waypoint, true)
	fmt.Println("Part 2: ", distance2)
}
