package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type SourceInterval struct {
	sourceStart int64
	sourceRange int64
}

type mapXtoY struct {
	destinationStart int64
	sourceStart      int64
	mapRange         int64
}

func main() {
	buffer, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal("Error : Unable to open the file")
	}
	fileInput := string(buffer)
	separatedFileInput := strings.Split(fileInput, "\n\n")
	var seeds []SourceInterval
	var seedToSoilMap []mapXtoY
	var soilToFertilizerMap []mapXtoY
	var fertilizerToWaterMap []mapXtoY
	var waterToLightMap []mapXtoY
	var lightToTemperatureMap []mapXtoY
	var temperatureToHumidityMap []mapXtoY
	var humidityToLocationMap []mapXtoY
	for i, input := range separatedFileInput {
		// Parse the seeds
		if i == 0 {
			splittedSeeds := strings.Split(strings.Split(input, ":")[1], " ")
			temp := []int64{}
			counter := 0
			for _, splittedSeed := range splittedSeeds {
				trimedSeed := strings.Trim(splittedSeed, " ")
				if len(trimedSeed) == 0 {
					continue
				}
				trimSeedAsInt, err := strconv.ParseInt(trimedSeed, 10, 64)
				if err != nil {
					log.Fatal("Error : Unable to transform " + trimedSeed + " into an integer")
				}
				counter++
				temp = append(temp, trimSeedAsInt)
				if counter != 0 && counter%2 == 0 {
					seeds = append(seeds, SourceInterval{sourceStart: temp[0], sourceRange: temp[1]})
					temp = []int64{}
				}
			}
			continue
		}
		// Parse the rest
		switch i {
		case 1:
			seedToSoilMap = processMap(input)
		case 2:
			soilToFertilizerMap = processMap(input)
		case 3:
			fertilizerToWaterMap = processMap(input)
		case 4:
			waterToLightMap = processMap(input)
		case 5:
			lightToTemperatureMap = processMap(input)
		case 6:
			temperatureToHumidityMap = processMap(input)
		case 7:
			humidityToLocationMap = processMap(input)
		default:
			log.Fatal("Error : Unexpected map received")
		}
	}
	// fmt.Printf("seedToSoilMap : %v\n", seedToSoilMap)
	// fmt.Printf("soilToFertilizerMap : %v\n", soilToFertilizerMap)
	// fmt.Printf("fertilizerToWaterMap : %v\n", fertilizerToWaterMap)
	// fmt.Printf("waterToLightMap : %v\n", waterToLightMap)
	// fmt.Printf("lightToTemperatureMap : %v\n", lightToTemperatureMap)
	// fmt.Printf("temperatureToHumidityMap : %v\n", temperatureToHumidityMap)
	// fmt.Printf("humidityToLocationMap : %v\n", humidityToLocationMap)
	// fmt.Printf("Seeds = %v\n", seeds)
	minLocations := []SourceInterval{}
	for _, seed := range seeds {
		sourceValue := []SourceInterval{seed}
		for i := 1; i <= 7; i++ {
			switch i {
			case 1:
				sourceValue = foundMatch(sourceValue, seedToSoilMap)
			case 2:
				sourceValue = foundMatch(sourceValue, soilToFertilizerMap)
			case 3:
				sourceValue = foundMatch(sourceValue, fertilizerToWaterMap)
			case 4:
				sourceValue = foundMatch(sourceValue, waterToLightMap)
			case 5:
				sourceValue = foundMatch(sourceValue, lightToTemperatureMap)
			case 6:
				sourceValue = foundMatch(sourceValue, temperatureToHumidityMap)
			case 7:
				sourceValue = foundMatch(sourceValue, humidityToLocationMap)
			default:
				log.Fatal("Error : Unexpected map received")
			}
		}
		tempMin := SourceInterval{sourceStart: math.MaxInt64, sourceRange: 0}
		for _, locations := range sourceValue {
			if locations.sourceStart < tempMin.sourceStart {
				tempMin = locations
			}
		}
		minLocations = append(minLocations, tempMin)
	}
	finalResult := SourceInterval{sourceStart: math.MaxInt64, sourceRange: 0}
	for _, location := range minLocations {
		if finalResult.sourceStart > location.sourceStart {
			finalResult = location
		}
	}
	fmt.Printf("Result = %v\n", finalResult)
}

func foundMatch(sources []SourceInterval, mappings []mapXtoY) []SourceInterval {
	resultsInterval := []SourceInterval{}
	intervalsToProcess := sources
	for _, mapping := range mappings {
		tempIntervalToProcess := []SourceInterval{}
		for _, sourceInterval := range intervalsToProcess {
			// [mappingInner...mappingOutter[
			mappingInner := mapping.sourceStart
			mappingOutter := mapping.sourceStart + mapping.mapRange
			// [sourceInner...sourceOutter[
			sourceInner := sourceInterval.sourceStart
			sourceOutter := sourceInterval.sourceStart + sourceInterval.sourceRange
			if sourceInner < mappingInner && sourceOutter <= mappingOutter && sourceOutter >= mappingInner {
				tempIntervalToProcess = append(tempIntervalToProcess, SourceInterval{sourceStart: sourceInner, sourceRange: mappingInner - sourceInner})
				resultsInterval = append(resultsInterval, SourceInterval{sourceStart: mapping.destinationStart, sourceRange: sourceOutter - mappingInner})
			} else if sourceInner >= mappingInner && sourceInner <= mappingOutter && sourceOutter > mappingOutter {
				offset := sourceInner - mappingInner
				resultsInterval = append(resultsInterval, SourceInterval{sourceStart: mapping.destinationStart + offset, sourceRange: mappingOutter - sourceInner})
				tempIntervalToProcess = append(tempIntervalToProcess, SourceInterval{sourceStart: mappingOutter, sourceRange: sourceOutter - mappingOutter})
			} else if sourceInner >= mappingInner && sourceOutter <= mappingOutter {
				offset := sourceInner - mappingInner
				resultsInterval = append(resultsInterval, SourceInterval{sourceStart: mapping.destinationStart + offset, sourceRange: sourceOutter - sourceInner})
			} else if sourceInner < mappingInner && sourceOutter > mappingInner {
				tempIntervalToProcess = append(tempIntervalToProcess, SourceInterval{sourceStart: sourceInner, sourceRange: mappingInner - sourceInner})
				resultsInterval = append(resultsInterval, SourceInterval{sourceStart: mapping.destinationStart, sourceRange: mappingOutter - mappingInner})
				tempIntervalToProcess = append(tempIntervalToProcess, SourceInterval{sourceStart: mappingOutter, sourceRange: sourceOutter - mappingOutter})
			} else {
				tempIntervalToProcess = append(tempIntervalToProcess, sourceInterval)
			}
		}
		intervalsToProcess = tempIntervalToProcess
	}
	return append(resultsInterval, intervalsToProcess...)
}

func processMap(mapping string) []mapXtoY {
	result := []mapXtoY{}
	splittedInput := strings.Split(mapping, "\n")
	dataToProcess := splittedInput[1:]
	for _, data := range dataToProcess {
		splittedData := strings.Split(data, " ")
		var dataAsInteger [3]int64
		for index, dataToTransform := range splittedData {
			trimmedDataToTransform := strings.Trim(dataToTransform, " ")
			transformedData, err := strconv.ParseInt(trimmedDataToTransform, 10, 64)
			if err != nil {
				log.Fatal("Error : Unable to transform " + trimmedDataToTransform + " into an integer")
			}
			dataAsInteger[index] = transformedData
		}
		result = append(result, mapXtoY{destinationStart: dataAsInteger[0], sourceStart: dataAsInteger[1], mapRange: dataAsInteger[2]})
	}
	return result
}
