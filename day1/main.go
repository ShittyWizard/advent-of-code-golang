package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	fmt.Println("Start solution of Day1 task of Advent-of-Code")
	file, err := os.Open("data.txt")
	// file, err := os.Open("testData.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Part 1
	answerPart1 := 0
	answerPart2 := 0
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		initialFuel := calculateFuel(mass)
		answerPart1 = answerPart1 + initialFuel
		answerPart2 = answerPart2 + initialFuel
		for initialFuel > 0 {
			initialFuel = calculateFuel(initialFuel)
			answerPart2 = answerPart2 + initialFuel
		}
	}
	fmt.Println("Part1 answer is:")
	fmt.Println(answerPart1)

	fmt.Println("Part2 answer is:")
	fmt.Println(answerPart2)
}

func calculateFuel(mass int) int {
	fuel := int(mass / 3) - 2
	if fuel >= 0 {
		return fuel
	}
	return 0
}

func calculateAdditionalFuel(mass int) int {
	fuel := int(mass / 3) - 2
	if fuel >= 0 {
		result := fuel + calculateAdditionalFuel(fuel)
		fmt.Println(result)
		return result
	}
	return 0
}
