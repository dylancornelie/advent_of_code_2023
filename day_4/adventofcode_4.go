package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const lineLength = 117

// const lineCount = 6
const lineCount = 201

func main() {
	// const isPart2 = true
	fileNamePath := "./input.txt"
	file, err := os.Open(fileNamePath)
	if err != nil {
		log.Fatal("Error : Unable to open the file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buffer := make([]byte, lineLength)
	scanner.Buffer(buffer, lineLength)
	var scratchCards []string
	var finalResult float64 = 0
	var finalNumberOfCards int = 0
	for scanner.Scan() {
		scratchCards = append(scratchCards, scanner.Text())
	}
	for index := range scratchCards {
		result, numberOfCard := solve(&scratchCards, index)
		finalResult += result
		finalNumberOfCards += numberOfCard
		finalNumberOfCards += 1
	}

	fmt.Println(finalNumberOfCards)
	fmt.Println(finalResult)
}

func solve(scratchCards *[]string, scratchCardsIndex int) (float64, int) {

	input := (*scratchCards)[scratchCardsIndex]
	splitedInput := strings.Split(input, ":")
	if len(splitedInput) != 2 {
		log.Fatal("Error : Unable to split the string " + input + " on \":\" character")
	}
	splitGame := strings.Split(splitedInput[1], "|")
	var winningCards, currentCards string = splitGame[0], splitGame[1]
	var winningCardsSet map[int8]bool = fromCardsStringToSet(winningCards)
	var currentCardsSet map[int8]bool = fromCardsStringToSet(currentCards)
	var finalResult float64 = 0
	var finalNumberOfCards int = 0
	intersection := setIntersection[int8](&winningCardsSet, &currentCardsSet)
	for range intersection {
		switch finalResult {
		case 0:
			finalResult = 1
		default:
			finalResult *= 2
		}
	}
	finalNumberOfCards += len(intersection)
	i := 1
	for range intersection {
		newIndex := scratchCardsIndex + i
		if newIndex <= lineCount {
			_, numberOfCard := solve(scratchCards, newIndex)
			finalNumberOfCards += numberOfCard
		}
		i++
	}
	return finalResult, finalNumberOfCards
}

func fromCardsStringToSet(cards string) map[int8]bool {
	result := make(map[int8]bool)
	for _, winningCard := range strings.Split(cards, " ") {
		trimedWinningCard := strings.Trim(winningCard, " ")
		if len(trimedWinningCard) == 0 {
			continue
		}
		winningCardAsInt, err := strconv.ParseInt(trimedWinningCard, 10, 8)
		if err != nil {
			log.Fatal("Error : Unable to convert " + trimedWinningCard + " into an integer")
		}
		result[int8(winningCardAsInt)] = true
	}
	return result
}

func setIntersection[T int8](A, B *map[T]bool) map[T]bool {
	unionResult := make(map[T]bool)
	for key := range *A {
		if (*B)[key] {
			unionResult[key] = true
		}
	}
	return unionResult
}
