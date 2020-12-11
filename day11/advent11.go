package main

import (
	"bufio"
	"errors"
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

func occuped(y int, x int, seats [][]string) (int, error) {
	if x < 0 || y < 0 || x >= len(seats[0]) || y >= len(seats) {
		return 0, errors.New("Fora de pantalla")
	}
	switch seats[y][x] {
	case "#":
		return 1, errors.New("Seient ocupat")
	case "L":
		return 0, errors.New("Seient buit")
	}
	return 0, nil
}

func peopleAround(y int, x int, distance int, seats [][]string) int {
	dxs := []int{-1, 0, +1}
	dys := []int{-1, 0, +1}
	people := 0

	for _, dy := range dys {
		for _, dx := range dxs {
			if dx == 0 && dy == 0 {
				continue
			}
			y0 := y
			x0 := x
			var err error
			sum := 0
			distanceActual := 0
			for err == nil && distanceActual < distance {
				y0 = y0 + dy
				x0 = x0 + dx
				distanceActual++
				sum, err = occuped(y0, x0, seats)
			}
			if err != nil {
				people = people + sum
			}
		}
	}

	return people
}

func stepSeats(occuped int, tolerance int, distance int, seats [][]string) (int, int, [][]string) {
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
				if peopleAround(y, x, distance, seats) >= tolerance {
					simbol = "L"
					changed++
					occuped--
				}
			case simbol == "L":
				if peopleAround(y, x, distance, seats) == 0 {
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

func fillSeats(seats [][]string, tolerance int, distance int) (int, int) {
	changed := -1
	occuped := 0
	for {
		occuped, changed, seats = stepSeats(occuped, tolerance, distance, seats)
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

	valids, changes := fillSeats(lines, 4, 1)

	fmt.Println("Cas 1: ", valids, "(", changes, ")")

	valids2, changes2 := fillSeats(lines, 5, len(lines))
	fmt.Println("Cas 2: ", valids2, "(", changes2, ")")

}
