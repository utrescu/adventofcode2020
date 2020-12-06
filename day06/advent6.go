package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]group, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var responses []group
	response := group{persons: 0, responses: make(map[string]int)}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), " ")
		if len(line) == 0 {
			// Nou password
			responses = append(responses, response)
			response = group{persons: 0, responses: make(map[string]int)}

		} else {
			response.persons = response.persons + 1
			for _, part := range line {
				letter := string(part)
				data, exists := response.responses[letter]
				if !exists {
					response.responses[letter] = 1
				} else {
					response.responses[letter] = data + 1
				}
			}
		}
	}

	responses = append(responses, response)
	return responses, scanner.Err()
}

type group struct {
	persons   int
	responses map[string]int
}

func main() {

	lines, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	valids1, valids2 := sumYesResponses(lines)

	fmt.Println("Cas 1: ", valids1)
	fmt.Println("Cas 2: ", valids2)
}

func sumYesResponses(lines []group) (int, int) {
	total1 := 0
	total2 := 0
	for _, value := range lines {
		total1 = total1 + len(value.responses)
		total2 = total2 + allRespond(value)
	}
	return total1, total2
}

func allRespond(grup group) int {
	suma := 0
	for _, v := range grup.responses {
		if grup.persons == v {
			suma = suma + 1
		}
	}
	return suma
}
