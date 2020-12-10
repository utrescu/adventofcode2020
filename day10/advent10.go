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
func readLines(path string) ([]adapter, int, error) {
	max := 0
	file, err := os.Open(path)
	if err != nil {
		return nil, max, err
	}
	defer file.Close()

	var lines []adapter
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		valor, _ := stringToInt(scanner.Text())
		if valor > max {
			max = valor
		}
		lines = append(lines, adapter{valor, false})
	}
	return lines, max, scanner.Err()
}

type adapter struct {
	valor int
	used  bool
}

func locate(jolts int, adapters []adapter) (int, int) {
	max := jolts + 4
	pos := -1
	fmt.Println("Buscant ", jolts)
	for value := range adapters {
		candidat := adapters[value].valor
		if !adapters[value].used &&
			candidat >= jolts &&
			candidat <= jolts+3 {

			if candidat < max {
				max = candidat
				pos = value
				fmt.Println("...Candidat ", candidat)
			}
		}
	}

	return pos, adapters[pos].valor
}

func packJolts(max int, adapters []adapter) (int, int, error) {

	used := 0
	numAdapters := len(adapters)
	var differences [4]int
	actualJolts := 0

	for used < numAdapters {
		index, newJolts := locate(actualJolts, adapters)
		fmt.Println("-------------- Diferènce - ", newJolts-actualJolts)
		differences[newJolts-actualJolts]++

		adapters[index].used = true
		used = used + 1

		actualJolts = newJolts
	}

	for actualJolts < max {
		differences[3]++
		actualJolts += 3
	}

	fmt.Println("cas 1:", differences[0], differences[1], differences[2], differences[3])
	return differences[1] * differences[3], 0, nil
}

func main() {
	numbers, max, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	var result1, result2, fails = packJolts(max+3, numbers)
	if fails != nil {
		panic(fails.Error())
	}
	fmt.Println("Cas 1: ", result1)
	fmt.Println("Cas 2: ", result2)
}
