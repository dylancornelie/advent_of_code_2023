package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Hand struct {
	cards                   [5]string
	bid                     int
	maxOccurence            int
	maxOccurenceCards       map[string]bool
	secondMaxOccurence      int
	secondMaxOccurenceCards map[string]bool
	sortScoreNumber1        int
	sortScoreNumber2        int
	sortScoreNumber3        int
	sortScoreNumber4        int
	sortScoreNumber5        int
}

func main() {

	buffer, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("Error : Unable to open the file")
	}
	fileInput := string(buffer)
	hands := []Hand{}
	for _, input := range strings.Split(fileInput, "\n") {
		splittedInput := strings.Split(input, " ")
		bidValue, err := strconv.Atoi(splittedInput[1])
		if err != nil {
			log.Fatal("Error : Unable to convert " + splittedInput[1] + " to an int value")
		}
		cards := [5]string{}
		for i, card := range strings.Split(splittedInput[0], "") {
			cards[i] = card
		}
		hands = append(hands, Hand{cards: cards, bid: bidValue})
	}

	cardScore := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}

	handTypeToHands := map[string][]Hand{
		"FiveOfAKind":  {},
		"FourOfAKind":  {},
		"FullHouse":    {},
		"ThreeOfAKind": {},
		"TwoPair":      {},
		"OnePair":      {},
		"HighCard":     {},
	}
	for _, hand := range hands {
		cardsOccurence := map[string]int{}
		maxOccurenceCards := map[string]bool{}
		secondMaxOccurenceCards := map[string]bool{}
		maxOccurence := 0
		secondMaxOccurence := 0
		for cardsIndex, card := range hand.cards {
			// If not exists returns 0
			previousOccurence := cardsOccurence[card]
			cardsOccurence[card] = previousOccurence + 1
			switch cardsIndex + 1 {
			case 1:
				hand.sortScoreNumber1 = cardScore[card]
			case 2:
				hand.sortScoreNumber2 = cardScore[card]
			case 3:
				hand.sortScoreNumber3 = cardScore[card]
			case 4:
				hand.sortScoreNumber4 = cardScore[card]
			case 5:
				hand.sortScoreNumber5 = cardScore[card]
			}
		}
		for card, occurence := range cardsOccurence {
			if occurence > maxOccurence {
				secondMaxOccurence = maxOccurence
				secondMaxOccurenceCards = maxOccurenceCards
				maxOccurence = occurence
				maxOccurenceCards = map[string]bool{card: true}
			} else if occurence == maxOccurence {
				maxOccurenceCards[card] = true
			} else if occurence < maxOccurence && occurence == secondMaxOccurence {
				secondMaxOccurenceCards[card] = true
			} else if occurence < maxOccurence && occurence > secondMaxOccurence {
				secondMaxOccurence = occurence
				secondMaxOccurenceCards = map[string]bool{card: true}
			}
		}
		hand.maxOccurence = maxOccurence
		hand.maxOccurenceCards = maxOccurenceCards
		hand.secondMaxOccurence = secondMaxOccurence
		hand.secondMaxOccurenceCards = secondMaxOccurenceCards
		if hand.maxOccurence == 5 {
			hands := handTypeToHands["FiveOfAKind"]
			hands = append(hands, hand)
			handTypeToHands["FiveOfAKind"] = hands
		} else if hand.maxOccurence == 4 {
			hands := handTypeToHands["FourOfAKind"]
			hands = append(hands, hand)
			handTypeToHands["FourOfAKind"] = hands
		} else if hand.maxOccurence == 3 && hand.secondMaxOccurence == 2 {
			hands := handTypeToHands["FullHouse"]
			hands = append(hands, hand)
			handTypeToHands["FullHouse"] = hands
		} else if hand.maxOccurence == 3 && hand.secondMaxOccurence == 1 {
			hands := handTypeToHands["ThreeOfAKind"]
			hands = append(hands, hand)
			handTypeToHands["ThreeOfAKind"] = hands
		} else if hand.maxOccurence == 2 && len(hand.maxOccurenceCards) != 1 {
			hands := handTypeToHands["TwoPair"]
			hands = append(hands, hand)
			handTypeToHands["TwoPair"] = hands
		} else if hand.maxOccurence == 2 && len(hand.maxOccurenceCards) == 1 {
			hands := handTypeToHands["OnePair"]
			hands = append(hands, hand)
			handTypeToHands["OnePair"] = hands
		} else if hand.maxOccurence == 1 && len(hand.maxOccurenceCards) == 5 {
			hands := handTypeToHands["HighCard"]
			hands = append(hands, hand)
			handTypeToHands["HighCard"] = hands
		}
	}
	fmt.Printf("Hand in FiveOfAKind = %v\n", len(handTypeToHands["FiveOfAKind"]))
	fmt.Printf("Hand in FourOfAKind = %v\n", len(handTypeToHands["FourOfAKind"]))
	fmt.Printf("Hand in FullHouse = %v\n", len(handTypeToHands["FullHouse"]))
	fmt.Printf("Hand in ThreeOfAKind = %v\n", len(handTypeToHands["ThreeOfAKind"]))
	fmt.Printf("Hand in TwoPair = %v\n", len(handTypeToHands["TwoPair"]))
	fmt.Printf("Hand in OnePair = %v\n", len(handTypeToHands["OnePair"]))
	fmt.Printf("Hand in HighCard = %v\n", len(handTypeToHands["HighCard"]))
	rank := 1
	totalWinnings := 0
	for i := 0; i < len(handTypeToHands); i++ {
		handType := ""
		switch i + 1 {
		case 7:
			handType = "FiveOfAKind"
		case 6:
			handType = "FourOfAKind"
		case 5:
			handType = "FullHouse"
		case 4:
			handType = "ThreeOfAKind"
		case 3:
			handType = "TwoPair"
		case 2:
			handType = "OnePair"
		case 1:
			handType = "HighCard"
		}
		hands := handTypeToHands[handType]
		// fmt.Printf("Number of hand with handType %v = %v\n", handType, len(hands))
		if len(hands) > 1 {
			handsSortedByWeight := orderHandsByCardsWeight(hands, 1)
			for _, hand := range handsSortedByWeight {
				fmt.Printf("\nHand is %v with rank = %v\n", hand.cards, rank)
				totalWinnings += hand.bid * rank
				rank += 1
			}
		} else if len(hands) == 1 {
			fmt.Printf("\nHand is %v with rank = %v\n", hands[0].cards, rank)
			totalWinnings += hands[0].bid * rank
			rank += 1
		}
	}
	fmt.Printf("\nFinal response : %v\n", totalWinnings)
}

func orderHandsByCardsWeight(hands []Hand, sortDepth int) []Hand {
	handsOrderedByCardsWeight := []Hand{}
	minWeightFound := math.MaxInt
	handMaxWeightIndex := []int{}
	for {
		for handIndex, hand := range hands {
			switch sortDepth {
			case 1:
				if hand.sortScoreNumber1 < minWeightFound {
					handMaxWeightIndex = []int{handIndex}
					minWeightFound = hand.sortScoreNumber1
				} else if hand.sortScoreNumber1 == minWeightFound {
					handMaxWeightIndex = append(handMaxWeightIndex, handIndex)
				}
			case 2:
				if hand.sortScoreNumber2 < minWeightFound {
					handMaxWeightIndex = []int{handIndex}
					minWeightFound = hand.sortScoreNumber2
				} else if hand.sortScoreNumber2 == minWeightFound {
					handMaxWeightIndex = append(handMaxWeightIndex, handIndex)
				}
			case 3:
				if hand.sortScoreNumber3 < minWeightFound {
					handMaxWeightIndex = []int{handIndex}
					minWeightFound = hand.sortScoreNumber3
				} else if hand.sortScoreNumber3 == minWeightFound {
					handMaxWeightIndex = append(handMaxWeightIndex, handIndex)
				}
			case 4:
				if hand.sortScoreNumber4 < minWeightFound {
					handMaxWeightIndex = []int{handIndex}
					minWeightFound = hand.sortScoreNumber4
				} else if hand.sortScoreNumber4 == minWeightFound {
					handMaxWeightIndex = append(handMaxWeightIndex, handIndex)
				}
			case 5:
				if hand.sortScoreNumber5 < minWeightFound {
					handMaxWeightIndex = []int{handIndex}
					minWeightFound = hand.sortScoreNumber5
				} else if hand.sortScoreNumber5 == minWeightFound {
					handMaxWeightIndex = append(handMaxWeightIndex, handIndex)
				}
			}
		}
		if len(handMaxWeightIndex) == 0 || len(hands) == 0 {
			return handsOrderedByCardsWeight
		} else if len(handMaxWeightIndex) == 1 {
			handsOrderedByCardsWeight = append(handsOrderedByCardsWeight, hands[handMaxWeightIndex[0]])
			hands[handMaxWeightIndex[0]] = hands[len(hands)-1]
			hands = hands[:len(hands)-1]
			// Reset states
			handMaxWeightIndex = []int{}
			minWeightFound = math.MaxInt
		} else if sortDepth == 5 && len(handMaxWeightIndex) > 1 {
			log.Fatal("We should never arrive at this point, it means we can not sort")
			// for offset, maxWeightIndex := range handMaxWeightIndex {
			// 	handsOrderedByCardsWeight = append(handsOrderedByCardsWeight, hands[maxWeightIndex-offset])
			// 	hands = append(hands[:maxWeightIndex-offset], hands[maxWeightIndex-offset+1:]...)
			// }
			// Reset states
			// handMaxWeightIndex = []int{}
			// minWeightFound = math.MaxInt
		} else {
			tempHands := []Hand{}
			for offset, maxWeightIndex := range handMaxWeightIndex {
				tempHands = append(tempHands, hands[maxWeightIndex-offset])
				hands = append(hands[:maxWeightIndex-offset], hands[maxWeightIndex-offset+1:]...)
			}
			tempResult := orderHandsByCardsWeight(tempHands, sortDepth+1)
			handsOrderedByCardsWeight = append(handsOrderedByCardsWeight, tempResult...)
			handMaxWeightIndex = []int{}
			minWeightFound = math.MaxInt
		}
	}
}
