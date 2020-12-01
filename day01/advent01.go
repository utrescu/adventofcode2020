package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func stringArrayToInt(stringArray []string) ([]int, error) {
	var result []int
	for _, value := range stringArray {
		numero, err := stringToInt(value)
		if err != nil {
			return nil, err
		}
		result = append(result, numero)
	}
	return result, nil
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
	filestrings, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	numbers, err := stringArrayToInt(filestrings)

	var result2, result3, fails = sumenDosMilVint(numbers)
	if fails != nil {
		panic(fails.Error())
	}
	fmt.Println("Cas 1: ", result2[0], result2[1], "=", result2[0]*result2[1])
	fmt.Println("Cas 2: ", result3[0], result3[1], result3[2], "=", result3[0]*result3[1]*result3[2])
}

func sumenDosMilVint(numbers []int) (result2 []int, result3 []int, err error) {
	found := 0
	for _, value1 := range numbers {
		for _, value2 := range numbers {
			if value1+value2 == 2020 {
				result2 = append(result2, value1, value2)
				found = found + 1
			}
			if len(result3) == 0 {
				for _, value3 := range numbers {
					if value1+value2+value3 == 2020 {
						result3 = append(result3, value1, value2, value3)
						found = found + 1
						break
					}
				}
			}
			if found == 2 {
				return result2, result3, nil
			}
		}

	}
	return result2, result3, errors.New("Cap número correcte")
}
