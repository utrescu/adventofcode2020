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

func moveTo(actual position, move string) position {
	x, y := actual.x, actual.y
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
	return position{x, y}
}

func moveToTile(moves []string) position {
	pos := position{0, 0}

	for _, move := range moves {
		pos = moveTo(pos, move)
	}
	return pos
}

func part1(tiles [][]string) (int, map[position]bool) {

	switchedTiles := map[position]bool{}

	for _, tile := range tiles {
		position := moveToTile(tile)
		if value, ok := switchedTiles[position]; !ok {
			switchedTiles[position] = true
		} else {
			switchedTiles[position] = !value
		}

	}

	return countBlacks(switchedTiles), switchedTiles
}

func getNumBlackNeighbors(cell position, mapa map[position]bool) int {
	sum := 0
	for _, move := range []string{"e", "w", "se", "sw", "ne", "nw"} {
		newPosition := moveTo(cell, move)
		if v, ok := mapa[newPosition]; ok {
			if v {
				sum++
			}
		}
	}

	return sum
}

func getNewNeighbors(cell position, mapa map[position]bool) []position {
	newCells := make([]position, 0)

	for _, move := range []string{"e", "w", "se", "sw", "ne", "nw"} {
		newPosition := moveTo(cell, move)
		if _, ok := mapa[newPosition]; !ok {
			newCells = append(newCells, newPosition)
		}
	}
	return newCells
}

func part2(mapa map[position]bool, days int) int {
	day := 0
	for day < days {
		newDay := map[position]bool{}
		blackcells := []position{}
		// faig les que tinc
		for cellposition, value := range mapa {
			veinsnegres := getNumBlackNeighbors(cellposition, mapa)
			newValue := value

			if value {
				if veinsnegres == 0 || veinsnegres > 2 {
					// Black to white
					newValue = !value
				}
				blackcells = append(blackcells, cellposition)
			} else {
				if veinsnegres == 2 {
					newValue = !value
				}
			}

			newDay[cellposition] = newValue
		}

		// per cada negra anterior
		// 	- veins blancs es tornen negres?
		for _, blackcell := range blackcells {
			for _, cell := range getNewNeighbors(blackcell, mapa) {
				veinsnegres := getNumBlackNeighbors(cell, mapa)
				if veinsnegres == 2 {
					newDay[cell] = true
				}
			}
		}

		// Copiar mapes

		mapa = newDay

		day++
	}

	return countBlacks(mapa)
}

func countBlacks(mapa map[position]bool) int {
	blacks := 0
	for _, v := range mapa {
		if v {
			blacks++
		}
	}
	return blacks
}

func main() {
	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1, mapa := part1(linies)
	fmt.Println("Part 1: ", correctes1)

	correctes2 := part2(mapa, 100)
	fmt.Println("Part 2: ", correctes2)

}
