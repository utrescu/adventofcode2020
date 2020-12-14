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

// Part 1 ------------------------------

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

// Part 2 -----------------

func calculateMemory(pos int, mask []rune, original []int, generateds [][]int) [][]int {
	if pos >= len(mask) {
		return generateds
	}

	switch mask[pos] {
	case '0':
		for i := range generateds {
			generateds[i] = append(generateds[i], original[pos])
		}
	case '1':
		for i := range generateds {
			generateds[i] = append(generateds[i], 1)
		}
	case 'X':
		for i := range generateds {
			zero := make([]int, len(generateds[i]))
			copy(zero, generateds[i])
			generateds[i] = append(generateds[i], 1)

			zero = append(zero, 0)
			generateds = append(generateds, zero)
		}
	}
	return calculateMemory(pos+1, mask, original, generateds)
}

func getFloatingRegister(mask []rune, numRegister int) []int {
	result := make([]int, 0)
	binValues := toBinari(numRegister)
	generateds := make([][]int, 1)

	registers := calculateMemory(0, mask, binValues, generateds)
	for _, numbers := range registers {
		result = append(result, toInt(numbers))
	}

	return result
}

func bitmaskvaluecount2(lines []instruction) int {
	var mask []rune
	registers := make(map[int]int)

	for _, line := range lines {
		if line.mask != nil {
			mask = line.mask
		} else {
			// L'enunciat no deixa clar que ara no s'ha de gravar a memòria
			// de la mateixa forma que a la part 1 ... (m'ha fet perdre molt de temps)
			newValue := line.value // processa(mask, line.value)

			floatingRegisters := getFloatingRegister(mask, line.register)
			for _, floatingRegister := range floatingRegisters {
				registers[floatingRegister] = newValue
			}
		}
	}

	suma := 0
	for _, value := range registers {
		suma += value
	}
	return suma
}

// --------------------------

func main() {
	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1 := bitmaskvaluecount(linies)
	fmt.Println("Part 1: ", correctes1)

	correctes2 := bitmaskvaluecount2(linies)
	fmt.Println("Part 2: ", correctes2)
}
