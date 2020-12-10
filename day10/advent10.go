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
func readLines(path string) ([]adapter, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []adapter
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		valor, _ := stringToInt(scanner.Text())
		lines = append(lines, adapter{valor, false})
	}
	return lines, scanner.Err()
}

type adapter struct {
	valor int
	used  bool
}

func locate(jolts int, adapters []adapter) (int, int) {
	max := jolts + 4
	pos := -1
	for value := range adapters {
		candidat := adapters[value].valor
		if !adapters[value].used &&
			candidat >= jolts &&
			candidat <= jolts+3 {

			if candidat < max {
				max = candidat
				pos = value
			}
		}
	}

	return pos, adapters[pos].valor
}

func packJolts(adapters []adapter) (int, int, error) {

	used := 0
	numAdapters := len(adapters)
	var differences [4]int
	actualJolts := 0

	for used < numAdapters {
		index, newJolts := locate(actualJolts, adapters)
		differences[newJolts-actualJolts]++

		adapters[index].used = true
		used = used + 1

		actualJolts = newJolts
	}

	// Last jump
	differences[3]++

	return differences[1] * differences[3], 0, nil
}

func main() {
	numbers, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	var result1, result2, fails = packJolts(numbers)
	if fails != nil {
		panic(fails.Error())
	}
	fmt.Println("Cas 1: ", result1)

	fmt.Println("Cas 2: ", result2)
}
