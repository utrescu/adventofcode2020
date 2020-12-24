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

func parse(line string) []string {

	result := make([]string, 0)
	direction := ""
	for _, caracter := range line {
		switch {
		case caracter == 'e' || caracter == 'w':
			direction += string(caracter)
			result = append(result, direction)
			direction = ""
		default:
			direction += string(caracter)
		}
	}
	return result
}

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
		line := scanner.Text()
		newInstruction := parse(line)
		lines = append(lines, newInstruction)
	}
	return lines, scanner.Err()
}

type position struct {
	x, y int
}

func moveToTile(moves []string) position {
	x := 0
	y := 0
	for _, move := range moves {
		switch move {
		case "e":
			x++
		case "w":
			x--
		case "se":
			x++
			y--
		case "sw":
			y--
		case "ne":
			y++
		case "nw":
			x--
			y++
		default:
			panic("Incorrect move")
		}
	}
	return position{x, y}
}

func part1(tiles [][]string) int {

	switchedTiles := map[position]bool{}

	for _, tile := range tiles {
		position := moveToTile(tile)
		if value, ok := switchedTiles[position]; !ok {
			switchedTiles[position] = true
		} else {
			switchedTiles[position] = !value
		}

	}

	black := 0
	for _, switched := range switchedTiles {
		if switched {
			black++
		}
	}

	return black
}

func main() {
	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1 := part1(linies)
	fmt.Println("Part 1: ", correctes1)

}
