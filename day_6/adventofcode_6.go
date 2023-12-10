package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     float64
	distance float64
}

const IS_PART_2 = true

func main() {
	buffer, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("Error : Unable to open the file")
	}
	fileInput := string(buffer)
	timeSlice := []float64{}
	distanceSlice := []float64{}
	computingTime := true
	for _, input := range strings.Split(fileInput, "\n") {
		dataAsStrings := strings.Split(strings.Split(input, ":")[1], " ")
		concatenatedNumber := ""
		for _, dataAsString := range dataAsStrings {
			cleanedData := strings.Trim(dataAsString, " ")
			if len(cleanedData) == 0 {
				continue
			}
			if IS_PART_2 == true {
				concatenatedNumber += dataAsString
			} else {
				dataAsInt, err := strconv.ParseFloat(cleanedData, 64)
				if err != nil {
					log.Fatal("Error : Unable to convert " + cleanedData + " to an int value")
				}
				if computingTime {
					timeSlice = append(timeSlice, dataAsInt)
				} else {
					distanceSlice = append(distanceSlice, dataAsInt)
				}
			}
		}
		if IS_PART_2 == true {
			dataAsInt, err := strconv.ParseFloat(concatenatedNumber, 64)
			if err != nil {
				log.Fatal("Error : Unable to convert " + concatenatedNumber + " to an int value")
			}
			if computingTime {
				timeSlice = append(timeSlice, dataAsInt)
			} else {
				distanceSlice = append(distanceSlice, dataAsInt)
			}
		}
		computingTime = false
	}
	races := []Race{}
	for index := range timeSlice {
		// + 1 because we need to beat the record and not equal it
		races = append(races, Race{time: timeSlice[index], distance: distanceSlice[index] + 1})
	}
	fmt.Printf("Races : %v\n", races)
	answers := [][2]float64{}
	for _, race := range races {
		answer := [2]float64{}
		// Δ = b² - 4ac = TIME² - (4*-1*-DISTANCE)
		delta := math.Pow(race.time, 2) - 4*race.distance
		fmt.Printf("Delta = %v\n", delta)
		switch {
		case delta == 0:
			// It means the record is an absolute record we can not do better
			// Should never happen in this puzzle then...
			answer[0] = race.time / -2
			answer[1] = 0
		case delta > 0:
			answer[0] = math.Ceil((-race.time + math.Sqrt(delta)) / -2)
			//
			answer[1] = math.Floor((-race.time-math.Sqrt(delta))/-2) + 1
		default:
			log.Fatal("Error : The function does not have solution in R")
		}
		answers = append(answers, answer)
	}
	fmt.Printf("Answers : %v\n", answers)
	result := float64(1)
	for _, answer := range answers {
		// If the record is unbeatable we can not win the puzzle
		// Let's say we don't count it for the sake of it
		if answer[1] != 0 {
			result *= (answer[1] - answer[0])
		}
	}
	fmt.Printf("Result : %f\n", result)
}
