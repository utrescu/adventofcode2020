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

type tiquet struct {
	data []int
}

func (c tiquet) hasValuesNotListedIn(valids map[int]bool) (bool, []int) {
	notListed := make([]int, 0)
	result := false
	for _, num := range c.data {
		_, found := valids[num]
		if !found {
			notListed = append(notListed, num)
			result = true
		}
	}
	return result, notListed
}

func generateTicket(line string) ([]int, error) {
	separa := strings.Split(line, ",")
	data := make([]int, 0)
	for _, value := range separa {
		number, err := stringToInt(value)
		if err != nil {
			return nil, err
		}
		data = append(data, number)
	}
	return data, nil
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (map[string]map[int]bool, tiquet, []tiquet, error) {
	var re = regexp.MustCompile(`(?m)(\d+)-(\d+)`)
	valids := make(map[string]map[int]bool)
	personal := tiquet{nil}
	nearby := make([]tiquet, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, personal, nil, err
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
				name := strings.Split(line, ":")[0]
				match := re.FindAllStringSubmatch(line, -1)
				for i := 0; i < len(match); i++ {
					actual := valids[name]
					if actual == nil {
						actual = make(map[int]bool, 0)
					}
					start, _ := stringToInt(match[i][1])
					end, _ := stringToInt(match[i][2])
					for n := start; n <= end; n++ {
						actual[n] = true
					}
					valids[name] = actual
				}
			case 1:
				if line != "your ticket:" {
					data, err := generateTicket(line)
					if err != nil {
						panic("ticket fails")
					}
					personal.data = data
				}
			case 2:
				if line != "nearby tickets:" {
					data, err := generateTicket(line)
					if err != nil {
						panic("nearby ticket fails")
					}
					nearby = append(nearby, tiquet{data})
				}
			}
		}

		lines = append(lines, scanner.Text())
	}
	return valids, personal, nearby, scanner.Err()
}

func soluciona(valids map[int]bool, others []tiquet) (int, []tiquet) {
	suma := 0

	corrects := make([]tiquet, 0)
	for _, ticket := range others {
		failed, results := ticket.hasValuesNotListedIn(valids)
		if failed {
			for _, result := range results {
				suma = suma + result
			}
		} else {
			corrects = append(corrects, ticket)
		}
	}
	return suma, corrects
}

func isCandidate(noteNumbers map[int]bool, actual int) bool {
	_, ok := noteNumbers[actual]
	return ok
}

func locatePossibleCandidates(noteNumbers map[int]bool, tickets []tiquet, posicio int) bool {
	valid := true

	for _, ticket := range tickets {
		valid = valid && isCandidate(noteNumbers, ticket.data[posicio])
		if valid == false {
			return false
		}
	}
	return valid

}

func locate(llista []int, value int) int {
	for i := range llista {
		if llista[i] == value {
			return i
		}
	}
	return -1
}

func remove(key string, value int, candidates map[string][]int) {
	for k, a := range candidates {
		if k != key {
			i := locate(a, value)
			if i != -1 {
				a[i] = a[len(a)-1]
				a = a[:len(a)-1]
				candidates[k] = a
			}
		}
	}
}

func purge(candidates map[string][]int) bool {
	for _, v := range candidates {
		if len(v) > 1 {
			return true
		}
	}
	return false
}

func soluciona2(notes map[string]map[int]bool, personal tiquet, tickets []tiquet) int {
	candidates := make(map[string][]int)

	tickets = append(tickets, personal)

	for noteName, noteNumbers := range notes {
		possibles := make([]int, 0)
		for i := 0; i < len(personal.data); i++ {
			if locatePossibleCandidates(noteNumbers, tickets, i) {
				possibles = append(possibles, i)
			}
		}
		candidates[noteName] = possibles
	}

	for purge(candidates) {
		for k, v := range candidates {
			if len(v) == 1 {
				remove(k, v[0], candidates)
			}
		}
	}

	// Solution
	solution := 1
	for k, v := range candidates {
		if strings.HasPrefix(k, "departure") {
			pos := v[0]
			solution *= personal.data[pos]
		}
	}
	return solution
}

func main() {
	valids, personal, others, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	numberList := make(map[int]bool)
	for _, validnumbers := range valids {
		for validnumber := range validnumbers {
			numberList[validnumber] = true
		}
	}
	correctes1, tiquetsCorrectes := soluciona(numberList, others)
	fmt.Println("Part 1: ", correctes1)

	correctes2 := soluciona2(valids, personal, tiquetsCorrectes)
	fmt.Println("Part 2: ", correctes2)
}
