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

type position struct {
	angle    float64
	direccio int
	east     int
	north    int
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

func mou(actual position, moviment string, value int) position {
	switch moviment {
	case "N":
		actual.north -= value
	case "S":
		actual.north += value
	case "E":
		actual.east += value
	case "W":
		actual.east -= value
	default:
		panic("Il·legal moviment")
	}
	return actual
}

func step(actual position, moviment move) position {
	directions := []string{"E", "N", "W", "S"}

	switch moviment.action {
	case "L":
		pas := moviment.value / 90
		actual.direccio = (actual.direccio + pas) % len(directions)
	case "R":
		pas := moviment.value / 90
		for pas > 0 {
			if actual.direccio-1 < 0 {
				actual.direccio = len(directions) - 1
			} else {
				actual.direccio--
			}
			pas--
		}

	case "F":
		actual = mou(actual, directions[actual.direccio], moviment.value)
	default:
		actual = mou(actual, moviment.action, moviment.value)
	}

	return actual
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func goAt(moves []move) int {
	actual := position{
		angle:    0,
		direccio: 0,
		east:     0,
		north:    0,
	}
	i := 1
	for _, mou := range moves {
		fmt.Println(i)
		actual = step(actual, mou)
		i++
	}
	return abs(actual.north) + abs(actual.east)
}

func main() {
	correctes2 := 0

	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	distance1 := goAt(linies)

	fmt.Println("Part 1: ", distance1)
	fmt.Println("Part 2: ", correctes2)
}
