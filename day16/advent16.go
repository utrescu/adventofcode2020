package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (map[int]string, []int, []int, error) {
	var re = regexp.MustCompile(`(?m)(\d+)-(\d+)`)
	valids := make(map[int]string)
	ticket := make([]int, 0)
	nearby := make([]int, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, nil, err
	}
	defer file.Close()

	var lines []string
	group := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			group++
		} else {
			switch group {
			case 0:
				match := re.FindAllStringSubmatch(line, -1)
				for i := 0; i < len(match); i++ {
					start, _ := stringToInt(match[i][1])
					end, _ := stringToInt(match[i][2])
					for n := start; n <= end; n++ {
						valids[n] = "si"
					}
				}
			case 1:
				separa := strings.Split(line, ",")
				for _, value := range separa {
					number, err := stringToInt(value)
					if err == nil {
						ticket = append(ticket, number)
					}
				}
			case 2:
				for _, value := range strings.Split(line, ",") {
					number, err := stringToInt(value)
					if err == nil {
						nearby = append(nearby, number)
					}
				}
			}
		}

		lines = append(lines, scanner.Text())
	}
	return valids, ticket, nearby, scanner.Err()
}

func main() {
	valids, ticket, others, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1 := soluciona(valids, ticket, others)

	fmt.Println("Part 1: ", correctes1)
}

func soluciona(valids map[int]string, ticket []int, others []int) int {
	suma := 0
	for _, value := range others {
		if valids[value] != "si" {
			suma += value
		}
	}
	return suma
}
