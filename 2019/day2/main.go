package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var codeArray []int
var data string

func main() {
	file, err := os.Open("data.txt") // What value is left at position 0 at final state?
	// file, err := os.Open("testData.txt") // final state - 30,1,1,4,2,5,6,0,99
	if err != nil {
		fmt.Println("File reading error")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = scanner.Text()
		codeArray = convertStringToInt(data)
	}

	fmt.Println("Part1 answer:")
	fmt.Println(getResultPart1(codeArray))

	// Part 2 (BRUTE FORCE - Bad solution, so TODO: find good solution)
	output := 19690720 // after calculation it is value at position 0
	start := time.Now()
	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			codeArray = convertStringToInt(data)
			codeArray[1] = n
			codeArray[2] = v
			// just try goroutines
			go checkAnswer(codeArray, n, v, output, start)
		}
	}
}

func checkAnswer(array []int, n int, v int, output int, start time.Time) {
	// fmt.Printf("I'm goroutine %d %d \n", n, v)
	if getResultPart1(array) == output {
		answerPart2 := 100*n + v
		fmt.Println("Part2 answer:")
		fmt.Println(answerPart2)
		duration := time.Since(start)

		fmt.Println("Duration in milliseconds:")
		fmt.Println(duration.Milliseconds())
		return
	}
}

func convertStringToInt(data string) []int {
	var array []int
	strCodeArray := strings.Split(data, ",")
	for _, val := range strCodeArray {
		v, _ := strconv.Atoi(val)
		array = append(array, v)
	}
	return array
}

// return value at position 0
func getResultPart1(array []int) int {
	stepsCount := int(len(array) / 4)

	i := 0
	for i < stepsCount {
		array = calculateOperation(array[i*4:i*4+4], array)
		i++
	}

	return array[0]
}

func calculateOperation(operation []int, array []int) []int {
	operationType := operation[0]
	firstPosition := operation[1]
	secondPosition := operation[2]
	resultPosition := operation[3]

	switch operationType {
	// addict
	case 1:
		array[resultPosition] = array[firstPosition] + array[secondPosition]
	//  multiply
	case 2:
		array[resultPosition] = array[firstPosition] * array[secondPosition]
	case 99:
		return array
	default:
		return array
	}
	return array
}
