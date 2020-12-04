package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([][]fields, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var passports [][]fields
	var passport []fields
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), " ")
		if len(line) == 0 {
			// Nou password
			passports = append(passports, passport)
			passport = make([]fields, 0)
			fmt.Println("")
		} else {
			for _, part := range strings.Split(line, " ") {
				fmt.Println(part)
				field := strings.Split(part, ":")
				passport = append(passport, fields{field[0], field[1]})

			}
		}

	}
	return passports, scanner.Err()
}

type fields struct {
	field string
	value string
}

func main() {
	fieldsRequired := []string{
		"byr", "iyr", "eyr", "hgt",
		"hcl", "ecl", "pid",
	}
	fieldOptional := "cid"

	lines, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	var valids1 = passwordValids(
		lines,
		fieldsRequired,
		fieldOptional)

	fmt.Println("Cas 1: ", valids1)

}

func passportContains(passport []fields, value string) bool {
	for _, field := range passport {
		if field.field == value {
			return true
		}
	}
	return false
}

func passportContainsFields(passport []fields, requireds []string) bool {
	var contains = 0
	for _, requiredField := range requireds {
		if passportContains(passport, requiredField) {
			contains = contains + 1
		}
	}
	if contains == len(requireds) {
		return true
	}
	return false
}

func passwordValids(passports [][]fields, required []string, optional string) int {

	valids := 0

	for _, passport := range passports {
		if passportContainsFields(passport, required) {
			valids = valids + 1
		}
	}
	return valids
}
