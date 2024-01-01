package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	buffer, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("Error : Unable to open the file")
	}
	fileInput := string(buffer)
	result := 0

	for _, historyAsString := range strings.Split(fileInput, "\n") {
		if len(historyAsString) == 0 {
			continue
		}
		historyValues := []int{}
		for _, splittedHistoryValue := range strings.Split(historyAsString, " ") {
			historyValue, err := strconv.Atoi(splittedHistoryValue)
			if err != nil {
				log.Fatal("Error : Unable to convert " + splittedHistoryValue + " to an int value")
			}
			historyValues = append(historyValues, historyValue)
		}
		result += findNextValueOfHistory(historyValues, true)
	}
	fmt.Println(result)
}

func findNextValueOfHistory(history []int, part2 bool) int {
	extrapolationArray := [][]int{}
	if part2 {
		reversedArray := []int{}
		for i := len(history) - 1; i >= 0; i-- {
			reversedArray = append(reversedArray, history[i])
		}
		extrapolationArray = append(extrapolationArray, reversedArray)
	} else {
		extrapolationArray = append(extrapolationArray, history)
	}

	depth := 0
	notAllZero := true
	for notAllZero {
		notAllZero = false
		temp_array := []int{}
		for index := range extrapolationArray[depth] {
			if index == 0 {
				continue
			}
			difference := extrapolationArray[depth][index] - extrapolationArray[depth][index-1]
			if difference != 0 {
				notAllZero = true
			}
			if part2 {
				difference *= -1
			}
			temp_array = append(temp_array, difference)
		}
		extrapolationArray = append(extrapolationArray, temp_array)
		depth += 1
	}
	max_depth := depth
	var valueToInsert int = 0
	for depth != -1 {
		lenNthExtrapArray := len(extrapolationArray[depth])
		// Check if there already is a previous computed value to insert or not (for initialisation)
		if depth == max_depth {
			valueToInsert = 0
		} else {
			if part2 {
				valueToInsert = extrapolationArray[depth][lenNthExtrapArray-1] + -valueToInsert
			} else {
				valueToInsert = extrapolationArray[depth][lenNthExtrapArray-1] + valueToInsert
			}
		}
		// Insert value
		extrapolationArray[depth] = append(extrapolationArray[depth], valueToInsert)
		depth -= 1
	}
	checkArray(extrapolationArray, max_depth, part2)
	return extrapolationArray[depth+1][len(extrapolationArray[depth+1])-1]
}

func checkArray(array [][]int, maxDepth int, part2 bool) {

	for depth := 0; depth < maxDepth; depth++ {

		lenNthExtrapArray := len(array[depth])
		if part2 {
			if array[depth][lenNthExtrapArray-2]-array[depth][lenNthExtrapArray-1] != array[depth+1][lenNthExtrapArray-2] {
				fmt.Println("Aïe !")
				printArray(array)
				return
			}
		} else {

			if array[depth][lenNthExtrapArray-1]-array[depth][lenNthExtrapArray-2] != array[depth+1][lenNthExtrapArray-2] {
				fmt.Println("Aïe !")
				printArray(array)
				return
			}
		}
	}

}

func printArray(array [][]int) {
	for i := range array {
		fmt.Printf("{%v} : ", i)
		for j := range array[i] {
			fmt.Printf("%v ", array[i][j])
		}
		fmt.Println()
	}
}
