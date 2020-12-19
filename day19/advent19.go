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

type rule struct {
	hasMore bool
	value   string
	rule    int
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (map[int]string, []string, error) {
	messages := make([]string, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	linies := make(map[int]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linia := scanner.Text()
		if strings.Contains(linia, ":") {
			values := strings.Split(linia, ": ")
			ruleNumber, _ := stringToInt(values[0])
			linies[ruleNumber] = values[1]
		} else {
			if linia != "" {
				messages = append(messages, linia)
			}
		}
	}
	return linies, messages, scanner.Err()
}

/// Part 1 ...

func validate(messages []string, rules map[int]string) int {

	valides1 := 0
	translated := make(map[int]string)
	reg := translate(0, rules, translated)

	var re = regexp.MustCompile("^" + reg + "$")
	// fmt.Println(reg)
	for _, message := range messages {

		match := re.FindAllString(message, -1)
		if len(match) != 0 {
			valides1++
		}

	}

	return valides1
}

func translate(actual int, rules map[int]string, translated map[int]string) string {
	value, ok := translated[actual]
	if ok {
		return value
	}

	currentRule, _ := rules[actual]
	if strings.Contains(currentRule, "\"") {
		return strings.Replace(currentRule, "\"", "", -1)
	}

	result := "("
	for _, part := range strings.Split(currentRule, " ") {
		// si és un número hem de buscar-ne un de nou
		digit, err := strconv.Atoi(part)
		if err != nil {
			if part == "|" {
				result += "|"
			}
		} else {
			result += translate(digit, rules, translated)
		}
	}
	result += ")"

	translated[actual] = result

	return result
}

// Part 2 ---------------

func main() {
	linies, messages, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1 := validate(messages, linies)

	fmt.Println("Part 1: ", correctes1)
}
