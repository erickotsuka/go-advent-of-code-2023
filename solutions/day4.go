package solutions

import (
	"fmt"
	"slices"
	"strings"

	"github.com/erickotsuka/go-advent-of-code-2023/utils"
)

func Day4Part1() {
	lines := strings.Split(utils.ReadInput(4), "\n")
	result := 0

	for _, line := range lines {
		labelNumbers := strings.Split(line, ": ")
		winningAndCardNumbers := strings.Split(labelNumbers[1], " | ")	
		var winningNumbers []string
		for i := 0; i < len(winningAndCardNumbers[0]); i += 3 {
			winningNumbers = append(winningNumbers, winningAndCardNumbers[0][i:i+2])
		}	
		quantityOfMatches := 0
		for i := 0; i < len(winningAndCardNumbers[1]); i += 3 {
			cardNumber := winningAndCardNumbers[1][i:i+2]
			if slices.Contains(winningNumbers, cardNumber) {
				quantityOfMatches += 1
			} 
		}
		if quantityOfMatches > 0 {
			result += 1 << (quantityOfMatches - 1)		
		}
	}

	fmt.Println(result)
}

func Day4Part2() {
	lines := strings.Split(utils.ReadInput(4), "\n")
	result := len(lines)

	copiesOfCardNumber := make(map[int]int)
	copiesOfCardNumber[0] = 0

	for cardNumber, line := range lines {
		labelNumbers := strings.Split(line, ": ")
		winningAndCardNumbers := strings.Split(labelNumbers[1], " | ")	
		var winningNumbers []string
		for i := 0; i < len(winningAndCardNumbers[0]); i += 3 {
			winningNumbers = append(winningNumbers, winningAndCardNumbers[0][i:i+2])
		}	
		quantityOfMatches := 0
		for i := 0; i < len(winningAndCardNumbers[1]); i += 3 {
			cardNumber := winningAndCardNumbers[1][i:i+2]
			if slices.Contains(winningNumbers, cardNumber) {
				quantityOfMatches += 1
			} 
		}
		for i := cardNumber + 1; i <= cardNumber + quantityOfMatches; i++ {
			result += copiesOfCardNumber[cardNumber] + 1
			copiesOfCardNumber[i] += copiesOfCardNumber[cardNumber] + 1
		}
	}

	fmt.Println(result)
}
