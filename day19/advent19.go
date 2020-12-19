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
func translate2(actual int, rules map[int]string, translated map[int]string, cut int) string {
	value, ok := translated[actual]
	if ok {
		return value
	}
	currentRule, _ := rules[actual]
	if strings.Contains(currentRule, "\"") {
		return strings.Replace(currentRule, "\"", "", -1)
	}
	result := "("

	switch actual {
	case 8:
		// Un o més 42
		result += translate2(42, rules, translated, cut) + "+"
	case 11:
		v42 := translate2(42, rules, translated, cut)
		v31 := translate2(31, rules, translated, cut)
		// un grapat de 42+ , (42 o 31)* i després 31+
		//
		// Els poso manualment perquè en algun moment he de parar el bucle infinit
		//  (4 dóna 417, 5 dóna 420, 6 dóna 422, 7 dóna 422).
		//  -- Per tant amb aquest cas podria deixar el 6 ...
		for i := 1; i < cut; i++ {
			if i > 1 {
				result += "|"
			}
			result += "("
			for j := 0; j < i; j++ {
				result += v42
			}
			for k := 0; k < i; k++ {
				result += v31
			}
			result += ")"
		}
		translated[11] = "(" + result + ")"
	default:

		for _, part := range strings.Split(currentRule, " ") {
			// si és un número hem de buscar-ne un de nou
			digit, err := strconv.Atoi(part)
			if err != nil {
				if part == "|" {
					result += "|"
				}
			} else {
				result += translate2(digit, rules, translated, cut)
			}
		}

	}
	result += ")"
	translated[actual] = result
	return result
}

// ------ principal

func evaluate(messages []string, reg string) int {
	valids := 0
	var re = regexp.MustCompile("^" + reg + "$")
	// fmt.Println(reg)
	for _, message := range messages {

		match := re.FindAllString(message, -1)
		if len(match) != 0 {
			valids++
		}
	}
	return valids
}

func validate(messages []string, rules map[int]string, part1 bool) int {

	translated := make(map[int]string)
	if part1 {
		return evaluate(messages, translate(0, rules, translated))

	} else {
		// Com que no sé quan podar vaig provant des de 5 fins que el
		// resultats es repeteixi
		oldValue := 0
		newValue := 1
		cutAt := 5

		for oldValue != newValue {
			translated := make(map[int]string)
			oldValue = newValue
			newValue = evaluate(messages, translate2(0, rules, translated, cutAt))
			fmt.Println(".... Cut", cutAt, "=", newValue)
			cutAt++
		}
		return newValue
	}

}

func main() {
	linies, messages, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1 := validate(messages, linies, true)
	fmt.Println("Part 1: ", correctes1)

	correctes2 := validate(messages, linies, false)
	fmt.Println("Part 2:", correctes2)
}
