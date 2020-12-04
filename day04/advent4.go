package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
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
		} else {
			for _, part := range strings.Split(line, " ") {
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

type passportFields struct {
	field      string
	validation string
}

func main() {
	fieldsRequired := []passportFields{
		{"byr", "^(19[2-9][0-9]|200[0-2])$"},
		{"iyr", "^20(1[0-9]|20)$"},
		{"eyr", "^20(2[0-9]|30)$"},
		{"hgt", "^(1[5-8][0-9]cm$|19[0-3]cm$)|(59|6[0-9]|7[0-6])in$"},
		{"hcl", "^#[0-9a-f]{6}$"},
		{"ecl", "^amb|blu|brn|gry|grn|hzl|oth$"},
		{"pid", "^[0-9]{9}$"},
	}
	fieldOptional := "cid"

	lines, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	valids1, valids2 := passwordValids(
		lines,
		fieldsRequired,
		fieldOptional)

	fmt.Println("Cas 1: ", valids1)
	fmt.Println("Cas 2: ", valids2)

}

// -- PART 1

func passportContains(passport []fields, value string) bool {
	for _, field := range passport {
		if field.field == value {
			return true
		}
	}
	return false
}

// --- PART 2: validate fields ----

func getField(passport []fields, name string) (fields, error) {
	for _, field := range passport {
		if field.field == name {
			return field, nil
		}
	}
	return fields{}, errors.New("Not found field")
}

func passportValidates(passport []fields, requireds []passportFields) bool {
	for _, required := range requireds {
		passportField, _ := getField(passport, required.field)

		match, _ := regexp.MatchString(required.validation, passportField.value)
		if !match {
			return false
		}
	}
	return true
}

// -------------

func passportIsCorrect(passport []fields, requireds []passportFields) (bool, bool) {
	var contains = 0
	for _, requiredField := range requireds {
		if passportContains(passport, requiredField.field) {
			contains = contains + 1
		}
	}

	if contains == len(requireds) {
		// ok, now Validate part 2
		return true, passportValidates(passport, requireds)
	}
	return false, false
}

func passwordValids(passports [][]fields, required []passportFields, optional string) (int, int) {

	valids1 := 0
	valids2 := 0

	for _, passport := range passports {
		one, two := passportIsCorrect(passport, required)
		if one {
			valids1 = valids1 + 1
			if two {
				valids2 = valids2 + 1
			}
		}
	}
	return valids1, valids2
}
