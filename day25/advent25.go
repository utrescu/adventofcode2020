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

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := stringToInt(line)
		lines = append(lines, number)
	}
	return lines, scanner.Err()
}

func loopSize(key int) int {
	subjectNumber := 7
	i := 1
	value := subjectNumber
	for {
		i++
		value = (value * subjectNumber) % 20201227
		if value == key {
			return i
		}

	}
}

func transform(key int, loop int) int {
	value := 1
	for i := 0; i < loop; i++ {
		value = (value * key) % 20201227
	}
	return value
}

func part1(numbers []int) int {
	subjects := make([]int, 2)
	for i, number := range numbers {
		subjects[i] = loopSize(number)
	}

	return transform(numbers[0], subjects[1])
}

func main() {
	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1 := part1(linies)
	fmt.Println("Part 1: ", correctes1)

}
