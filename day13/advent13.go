package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToInt(str string) (uint64, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.ParseUint(nonFractionalPart[0], 10, 64)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (uint64, []uint64, error) {
	file, err := os.Open(path)
	if err != nil {
		return ^uint64(0), nil, err
	}
	defer file.Close()

	var lines []uint64
	scanner := bufio.NewScanner(file)

	var timestamp uint64
	i := 0
	for scanner.Scan() {
		if i == 0 {
			timestamp, _ = stringToInt(scanner.Text())
			i++
		} else {
			for _, text := range strings.Split(scanner.Text(), ",") {
				number, err := stringToInt(text)
				if err == nil {
					lines = append(lines, number)
				}
			}
		}
	}
	return timestamp, lines, scanner.Err()
}

func main() {
	timestamp, busos, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	correctes1 := firstBus(timestamp, busos)

	fmt.Println("Part 1: ", correctes1)
}

func firstBus(timestamp uint64, busos []uint64) uint64 {

	bestBus := ^uint64(0)
	timeWait := ^uint64(0)

	for _, bus := range busos {
		nextAt := ((timestamp / bus) + 1) * bus
		wait := nextAt - timestamp
		if wait < timeWait {
			bestBus = bus
			timeWait = wait
		}
	}

	return bestBus * timeWait
}
