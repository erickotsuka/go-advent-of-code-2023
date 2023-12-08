package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/erickotsuka/go-advent-of-code-2023/utils"
)

type handType int

const (
	fiveOfAKind handType = iota
	fourOfAKind
	fullHouse
	threeOfAKind
	twoPair
	onePair
	highCard
)

type handTypeBid struct {
	hand string
	handType handType 
	bid int
}

type handBid struct {
	hand string
	bid int
}

func Day7Part1() {
	// cards := []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	cardToStrengthMap := make(map[rune]rune, 13)

	cardToStrengthMap['2'] = '1'
	cardToStrengthMap['3'] = '2'
	cardToStrengthMap['4'] = '3'
	cardToStrengthMap['5'] = '4'
	cardToStrengthMap['6'] = '5'
	cardToStrengthMap['7'] = '6'
	cardToStrengthMap['8'] = '7'
	cardToStrengthMap['9'] = '8'
	cardToStrengthMap['T'] = '9'
	cardToStrengthMap['J'] = 'A'
	cardToStrengthMap['Q'] = 'B' 
	cardToStrengthMap['K'] = 'C'
	cardToStrengthMap['A'] = 'D'

	handToBidMap := make(map[string]int)
	handTypeToHandsMap := make(map[handType][]string)

	result := 0

	lines := strings.Split(utils.ReadInput(7), "\n")

	for _, line := range lines {
		handAndBid := strings.Split(line, " ")	
		hand := handAndBid[0]
		bid, err := strconv.Atoi(handAndBid[1])
		if err != nil {
			panic(err)
		}

		handToBidMap[hand] = bid

		cardCountMap := make(map[rune]int)
		for _, card := range hand {
			cardCountMap[card] = cardCountMap[card] + 1	
		}

		var handType handType

		switch len(cardCountMap) {
		case 1:
			handType = fiveOfAKind
		case 2:
			for _, count := range cardCountMap {
				if count == 1 || count == 4 {
					handType = fourOfAKind
				} else {
					handType = fullHouse
				}
				break
			}
		case 3:
			for _, count := range cardCountMap {
				if count == 1 {
					continue
				}
				if count == 3 {
					handType = threeOfAKind
					break
				} 
				if count == 2 {
					handType = twoPair
					break
				}
			}
		case 4:
			handType = onePair
		case 5:
			handType = highCard
		}

		hands, handTypeExists := handTypeToHandsMap[handType]
		if handTypeExists {
			hands = append(hands, hand)	
		} else {
			hands = []string{hand} 
		}
		handTypeToHandsMap[handType] = hands 

	}

	for handType, hands := range handTypeToHandsMap {
		handStrenghtStrings := make([]string, len(hands))
		handStrenghtStringToHandMap := make(map[string]string, len(hands))
		for index, hand := range hands {
			handStrengthString := ""
			for _, card := range hand {
				handStrengthString = handStrengthString + string(cardToStrengthMap[card])
			}
			handStrenghtStringToHandMap[handStrengthString] = hand
			handStrenghtStrings[index] = handStrengthString
		}
		slices.Sort(handStrenghtStrings)
		sortedHands := make([]string, len(hands))
		for index, handStrenghtString := range handStrenghtStrings {
			sortedHands[index] = handStrenghtStringToHandMap[handStrenghtString]
		}
		handTypeToHandsMap[handType] = sortedHands
	}

	weakestToStrongestHandTypes := []handType{ highCard, onePair, twoPair, threeOfAKind, fullHouse, fourOfAKind, fiveOfAKind }
	
	currentRank := 1
	for _, handType := range weakestToStrongestHandTypes {
		for _, hand := range handTypeToHandsMap[handType] {
			bid := handToBidMap[hand]
			result += bid * currentRank
			currentRank++
		}	
	}

	fmt.Println(result)

}
