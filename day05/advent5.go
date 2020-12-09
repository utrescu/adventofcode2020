package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	id1, myID := calculaBigID(linies)

	fmt.Println("Part 1: ", id1)

	fmt.Println("Part 2: ", myID)
}

// -- PART 1

func decode(code string, values int) int {

	min := 0
	max := values

	for _, letter := range code {
		if letter == 'F' || letter == 'L' {
			max = max - (max-min)/2
		} else {
			min = min + (max-min)/2
		}
	}
	return min
}

// --- PART 2

func findForat(ids []int) int {

	for pos, currentID := range ids {
		if pos+2 < len(ids) {
			if currentID+2 == ids[pos+1] {
				return currentID + 1
			}
		}
	}
	panic("No solution")
}

func calculaBigID(lines []string) (int, int) {
	big := 0
	var ids []int

	for _, line := range lines {
		fila, col := decode(line[:7], 128), decode(line[7:], 8)

		id := fila*8 + col
		ids = append(ids, id)
		if id > big {
			big = id
		}
	}

	// PART 2
	sort.Ints(ids)
	myID := findForat(ids)

	return big, myID
}
