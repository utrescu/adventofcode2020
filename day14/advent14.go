package main

import (
	"bufio"
	"fmt"
	"math"
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
func readLines(path string) ([]instruction, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line instruction
		line.create(scanner.Text())
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}

type instruction struct {
	mask     []rune
	register int
	value    int
}

func (i *instruction) create(line string) {
	var re1 = regexp.MustCompile(`(?m)^mask = (.+)$`)
	var re2 = regexp.MustCompile(`(?m)^mem\[(\d+)\] = (\d+)$`)

	match := re1.FindStringSubmatch(line)
	if len(match) == 2 {
		i.mask = []rune(match[1])
	} else {
		match = re2.FindStringSubmatch(line)
		i.mask = nil
		i.register, _ = stringToInt(match[1])
		i.value, _ = stringToInt(match[2])
	}
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func toBinari(number int) []int {
	value := make([]int, 36)

	i := 0
	for number > 1 {
		bit := number % 2
		number /= 2
		value[i] = bit
		i++
	}
	value[i] = number

	return reverse(value)
}

func toInt(bits []int) int {
	value := 0
	index := int(math.Pow(2, 35))
	for _, bit := range bits {
		value += bit * index
		index /= 2
	}
	return value
}

func processa(mask []rune, value int) int {
	binValues := toBinari(value)
	for i := range binValues {
		switch mask[i] {
		case '0':
			binValues[i] = 0
		case '1':
			binValues[i] = 1
		}
	}
	return toInt(binValues)
}

func bitmaskvaluecount(lines []instruction) int {
	var mask []rune
	registers := make(map[int]int)

	for _, line := range lines {
		if line.mask != nil {
			mask = line.mask
		} else {
			registers[line.register] = processa(mask, line.value)
		}
	}

	suma := 0
	for _, value := range registers {
		suma += value
	}
	return suma
}

func main() {
	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1 := bitmaskvaluecount(linies)

	fmt.Println("Part 1: ", correctes1)
}
