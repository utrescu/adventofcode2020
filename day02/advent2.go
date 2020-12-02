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

	correctes1, correctes2, err := processaPasswords1(linies)

	fmt.Println("Part 1: ", correctes1)
	fmt.Println("Part 2: ", correctes2)
}

func processaPasswords1(linies []string) (int, int, error) {

	passwords := 0
	passwords2 := 0

	for _, linia := range linies {

		if esCorrecte(linia) {
			passwords = passwords + 1
		}
		if esCorrecte2(linia) {
			passwords2 = passwords2 + 1
		}
	}

	return passwords, passwords2, nil
}

func esCorrecte(linia string) bool {
	parts := strings.Split(linia, " ")
	minimax := strings.Split(parts[0], "-")

	min, _ := stringToInt(minimax[0])
	max, _ := stringToInt(minimax[1])
	lletra := parts[1][0:1]

	suma := 0
	for _, candidata := range parts[2] {
		lletraCandidata := string(candidata)
		if lletraCandidata == lletra {
			suma = suma + 1
		}
	}

	if suma >= min && suma <= max {
		return true
	}

	return false
}

func esCorrecte2(linia string) bool {
	parts := strings.Split(linia, " ")
	minimax := strings.Split(parts[0], "-")

	min, _ := stringToInt(minimax[0])
	max, _ := stringToInt(minimax[1])
	lletra := parts[1][0:1]

	isMin := parts[2][min-1:min] == lletra
	isMax := parts[2][max-1:max] == lletra

	if isMin == isMax {
		return false
	}

	return true
}
