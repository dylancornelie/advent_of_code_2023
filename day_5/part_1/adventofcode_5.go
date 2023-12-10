package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal("Error : Unable to open the file")
	}
	var seeds []int64
	seedToSoilMap := [][]int64{}
	soilToFertilizerMap := [][]int64{}
	fertilizerToWaterMap := [][]int64{}
	waterToLightMap := [][]int64{}
	lightToTemperatureMap := [][]int64{}
	temperatureToHumidityMap := [][]int64{}
	humidityToLocationMap := [][]int64{}
	fileInput := string(b)
	separatedFileInput := strings.Split(fileInput, "\n\n")
	for i, input := range separatedFileInput {
		if i == 0 {
			splittedSeeds := strings.Split(strings.Split(input, ":")[1], " ")
			for _, splittedSeed := range splittedSeeds {
				trimedSeed := strings.Trim(splittedSeed, " ")
				if len(trimedSeed) == 0 {
					continue
				}
				trimSeedAsInt, err := strconv.ParseInt(trimedSeed, 10, 64)
				if err != nil {
					log.Fatal("Error : Unable to transform " + trimedSeed + " into an integer")
				}
				seeds = append(seeds, int64(trimSeedAsInt))
			}
			continue
		}
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
	fmt.Printf("seeds : %v\n", seeds)
	fmt.Printf("seedToSoilMap : %v\n", seedToSoilMap)
	fmt.Printf("soilToFertilizerMap : %v\n", soilToFertilizerMap)
	fmt.Printf("fertilizerToWaterMap : %v\n", fertilizerToWaterMap)
	fmt.Printf("waterToLightMap : %v\n", waterToLightMap)
	fmt.Printf("lightToTemperatureMap : %v\n", lightToTemperatureMap)
	fmt.Printf("temperatureToHumidityMap : %v\n", temperatureToHumidityMap)
	fmt.Printf("humidityToLocationMap : %v\n", humidityToLocationMap)
	var finalSeed int64
	var finalLocation int64 = math.MaxInt64
	for _, seed := range seeds {
		actualValue := seed
		for i := 0; i < 7; i++ {
			switch i {
			case 0:
				actualValue = foundMatch(actualValue, seedToSoilMap)
			case 1:
				actualValue = foundMatch(actualValue, soilToFertilizerMap)
			case 2:
				actualValue = foundMatch(actualValue, fertilizerToWaterMap)
			case 3:
				actualValue = foundMatch(actualValue, waterToLightMap)
			case 4:
				actualValue = foundMatch(actualValue, lightToTemperatureMap)
			case 5:
				actualValue = foundMatch(actualValue, temperatureToHumidityMap)
			case 6:
				actualValue = foundMatch(actualValue, humidityToLocationMap)
			default:
				log.Fatal("Error : Invalid index received")
			}
		}
		if actualValue < finalLocation {
			finalLocation = actualValue
			finalSeed = seed
		}
	}
	fmt.Printf("finalLocation : %v\n", finalLocation)
	fmt.Printf("finalSeed : %v\n", finalSeed)
}

func foundMatch(source int64, mappings [][]int64) int64 {
	for _, mapping := range mappings {
		mapSource := mapping[1]
		mapRange := mapping[2]
		if source < mapSource || source > mapSource+(mapRange-1) {
			continue
		}
		offset := source - mapSource
		return mapping[0] + offset
	}
	return source
}

func processMap(mapping string) [][]int64 {
	result := [][]int64{}
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
		result = append(result, dataAsInteger[:])
	}
	return result
}
