package main

import (
	"bufio"
	"errors"
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

const re = `(?m)^(nop|acc|jmp)\s([+-]\d*)$`

func parse(line []string) instruction {
	result := instruction{
		operation:     line[1],
		timesExecuted: 0,
	}
	if len(line) > 2 {
		result.value, _ = stringToInt(line[2])
	}
	return result
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]instruction, error) {
	reg := regexp.MustCompile(re)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		newInstruction := parse(reg.FindStringSubmatch(line))
		lines = append(lines, newInstruction)
	}
	return lines, scanner.Err()
}

type instruction struct {
	operation     string
	value         int
	timesExecuted int
}

type machine struct {
	index      int
	acumulator int
}

func execute(instruccio instruction, state machine) machine {

	switch {
	case instruccio.operation == "acc":
		state.acumulator = state.acumulator + instruccio.value
		state.index = state.index + 1
	case instruccio.operation == "jmp":
		state.index = state.index + instruccio.value
	default:
		state.index = state.index + 1

	}

	return state
}

func boot(linies []instruction) (int, error) {
	state := machine{
		index:      0,
		acumulator: 0,
	}

	// reset execution times (for part2)
	for i := range linies {
		linies[i].timesExecuted = 0
	}

	for state.index < len(linies) {
		// execute line
		linies[state.index].timesExecuted = linies[state.index].timesExecuted + 1

		state = execute(linies[state.index], state)
		if state.index >= len(linies) {
			break
		}
		if linies[state.index].timesExecuted > 0 {
			return state.acumulator, errors.New("Loop detected")
		}
	}

	return state.acumulator, nil
}

// ---------------   PART 2

func repairInstructions(linies []instruction) (int, error) {
	// Part 2
	for index, instruccio := range linies {
		old := instruccio.operation
		if instruccio.operation == "acc" {
			continue
		}
		if instruccio.operation == "jmp" {
			linies[index].operation = "nop"
		} else if instruccio.operation == "nop" {
			linies[index].operation = "jmp"
		}
		resultat, err := boot(linies)
		if err == nil {
			return resultat, nil
		}
		linies[index].operation = old
	}

	return -1, errors.New("No solution")
}

func main() {
	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1, _ := boot(linies)

	fmt.Println("Part 1: ", correctes1)

	correctes2, err := repairInstructions(linies)
	if err != nil {
		panic("Part 2 has no solution")
	}

	fmt.Println("Part 2: ", correctes2)
}
