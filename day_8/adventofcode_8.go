package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type NodeHeadings struct {
	left  string
	right string
}

const IS_PART_2 = true

func main() {
	buffer, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("Error : Unable to open the file")
	}
	fileInput := string(buffer)
	leftRightInstructions := []string{}
	nodeMap := map[string]NodeHeadings{}
	currentNodes := []string{}
	for _, input := range strings.Split(fileInput, "\n") {
		if len(input) == 0 {
			continue
		}
		if len(leftRightInstructions) == 0 {
			leftRightInstructions = strings.Split(input, "")
			continue
		}
		splittedInput := strings.Split(input, "=")
		nodeName := strings.Trim(splittedInput[0], " ")
		if strings.Compare(nodeName[len(nodeName)-1:], "A") == 0 {
			currentNodes = append(currentNodes, nodeName)
		}
		nodeHeadingsStrings := strings.Split(strings.ReplaceAll(strings.ReplaceAll(splittedInput[1], "(", ""), ")", ""), ",")
		nodeHeadings := NodeHeadings{left: "", right: ""}
		for _, nodeHeadingString := range nodeHeadingsStrings {
			if len(nodeHeadings.left) == 0 {
				nodeHeadings.left = strings.Trim(nodeHeadingString, " ")
			} else {
				nodeHeadings.right = strings.Trim(nodeHeadingString, " ")
			}
		}
		nodeMap[nodeName] = nodeHeadings
	}
	currentNode := "AAA"
	step := 0
	leftRightInstructionLength := len(leftRightInstructions)
	fmt.Printf("CurrentNodes = %v\n", currentNodes)
	steps := []int{}
	if IS_PART_2 {
		for _, currentNode := range currentNodes {
			step := 0
			for {
				if strings.Compare(currentNode[len(currentNode)-1:], "Z") == 0 {
					break
				}
				if leftRightInstructions[step%leftRightInstructionLength] == "L" {
					currentNode = nodeMap[currentNode].left
				} else {
					currentNode = nodeMap[currentNode].right
				}
				step += 1
			}
			steps = append(steps, step)
		}
		// All path don't have the same length...
		fmt.Printf("Reached XXZ in %v steps.\n", LCM(steps[0], steps[1], steps[2:]...))

	} else {
		for {
			if currentNode == "ZZZ" {
				fmt.Printf("Reached ZZZ in %v steps.\n", step)
				break
			}
			if leftRightInstructions[step%leftRightInstructionLength] == "L" {
				// fmt.Println("L")
				currentNode = nodeMap[currentNode].left
			} else {
				// fmt.Println("R")
				currentNode = nodeMap[currentNode].right
			}
			step += 1
		}
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
