package main

import (
	"fmt"
)

func soluciona1(numbers []int, end int) int {
	play := make(map[int]int)

	for i, number := range numbers[:len(numbers)-1] {

		play[number] = i + 1
	}
	turn := len(numbers)
	next := numbers[len(numbers)-1]

	for turn < end {

		lastTurn, found := play[next]
		if !found {
			play[next] = turn
			next = 0
		} else {
			play[next] = turn
			next = turn - lastTurn
		}
		turn++
	}

	return next
}

func main() {

	startingNumbers := []int{1, 0, 16, 5, 17, 4}

	correctes1 := soluciona1(startingNumbers, 2020)
	fmt.Println("Part 1: ", correctes1)

	correctes2 := soluciona1(startingNumbers, 30000000)
	fmt.Println("Part 2: ", correctes2)

}
