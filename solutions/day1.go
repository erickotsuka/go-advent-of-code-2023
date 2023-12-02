package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/erickotsuka/go-advent-of-code-2023/utils"
)

func Day1Part1() {
	lines := strings.Split(utils.ReadInput(1), "\r\n")
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	result := 0
	for _, v := range lines {
		var numberStr string
		for _, char := range v {
			if slices.Contains(digits, char) {
				numberStr = numberStr + string(char)
			}
		}
		numberStr = string(numberStr[0]) + string(numberStr[len(numberStr)-1])
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		result += number
	}
	fmt.Println(result)
}

func Day1Part2() {
	lines := strings.Split(utils.ReadInput(1), "\r\n")
	digitWords := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	textToNumberMap := map[string]int{
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	result := 0
	for _, v := range lines {
		var positions []int
		positionToDigitMap := make(map[int]int)
		for _, digitWord := range digitWords {
			firstIndex := strings.Index(v, digitWord)
			lastIndex := strings.LastIndex(v, digitWord)
			if firstIndex != -1 {
				positionToDigitMap[firstIndex] = textToNumberMap[digitWord]
				positions = append(positions, firstIndex)
			}
			if lastIndex != firstIndex {
				positionToDigitMap[lastIndex] = textToNumberMap[digitWord]
				positions = append(positions, lastIndex)
			}
		}
		firstDigitPosition := slices.Min(positions)
		lastDigitPosition := slices.Max(positions)
		result += positionToDigitMap[firstDigitPosition]*10 + positionToDigitMap[lastDigitPosition]
	}
	fmt.Println(result)
}
