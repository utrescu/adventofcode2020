package main

import (
	"bufio"
	"fmt"
	"os"
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

	id1 := calculaBigID(linies)

	fmt.Println("Part 1: ", id1)

	fmt.Println(decodeLine("BFFFBBFRRR"))
}

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

func decodeLine(line string) (int, int) {

	return decode(line[:7], 128), decode(line[7:], 8)

}

func calculaBigID(lines []string) int {
	big := 0

	for _, line := range lines {
		fila, col := decodeLine(line)

		id := fila*8 + col
		if id > big {
			big = id
		}
	}

	return big
}
