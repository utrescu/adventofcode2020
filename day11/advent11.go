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

	var seats [][]string
	var lineseats []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineseats = make([]string, 0)
		line := strings.TrimRight(scanner.Text(), " ")
		for _, letter := range line {
			lineseats = append(lineseats, string(letter))
		}
		seats = append(seats, lineseats)
	}

	return seats, scanner.Err()
}

func occuped(y int, x int, seats [][]string) int {
	if x < 0 || y < 0 || x >= len(seats[0]) || y >= len(seats) {
		return 0
	}
	if seats[y][x] == "#" {
		return 1
	}
	return 0
}

func peopleAround(y int, x int, seats [][]string) int {
	xs := []int{x - 1, x, x + 1}
	ys := []int{y - 1, y, y + 1}
	people := 0

	for _, y0 := range ys {
		for _, x0 := range xs {
			if x == x0 && y == y0 {
				continue
			}
			people += occuped(y0, x0, seats)
		}
	}

	return people
}

func stepSeats(occuped int, seats [][]string) (int, int, [][]string) {
	numCols := len(seats[0])
	numRows := len(seats)
	var newseats [][]string
	changed := 0
	for y := 0; y < numRows; y++ {
		line := make([]string, 0)
		for x := 0; x < numCols; x++ {
			simbol := seats[y][x]
			switch {
			case simbol == "#":
				if peopleAround(y, x, seats) >= 4 {
					simbol = "L"
					changed++
					occuped--
				}
			case simbol == "L":
				if peopleAround(y, x, seats) == 0 {
					simbol = "#"
					changed++
					occuped++
				}
			default:

			}
			line = append(line, simbol)
		}
		newseats = append(newseats, line)
	}

	// printseats(newseats)

	return occuped, changed, newseats
}

func printseats(seats [][]string) {
	for _, line := range seats {
		for _, car := range line {
			fmt.Print(car)
		}
		fmt.Println()
	}
	fmt.Println()
}

func fillSeats(seats [][]string) (int, int) {
	changed := -1
	occuped := 0
	for {
		occuped, changed, seats = stepSeats(occuped, seats)
		if changed == 0 {
			break
		}
	}
	return occuped, changed
}

func main() {

	lines, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	valids, changes := fillSeats(lines)

	fmt.Println("Cas 1: ", valids, "(", changes, ")")

}
