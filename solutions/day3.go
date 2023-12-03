package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/erickotsuka/go-advent-of-code-2023/utils"
)

func isDigit(char rune) bool {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	return slices.Contains(digits, char)
}

func isSymbol(char rune) bool {
	return char != '.' && !isDigit(char)
}

func Day3Part1() {
	lines := strings.Split(utils.ReadInput(3), "\n")
	rowQuantity := len(lines)
	columnsQuantity := len(lines[0])

	schematicMatrix := make([][]rune, rowQuantity)
	result := 0

	for rowNumber, line := range lines {
		schematicMatrix[rowNumber] = []rune(line)
	}

	for i := 0; i < len(schematicMatrix); i++ {
		var numberStr string
		hasAdjacentSymbol := false

		for j := 0; j < len(schematicMatrix[i]); j++ {
			if schematicMatrix[i][j] == '.' || isSymbol(schematicMatrix[i][j]) {
				if len(numberStr) > 0 {
					if hasAdjacentSymbol {
						convertedNumber, err := strconv.Atoi(numberStr)
						if err != nil {
							panic(err)
						}
						result += convertedNumber
					}
					numberStr = ""
					hasAdjacentSymbol = false
				}
				continue
			}
			numberStr += string(schematicMatrix[i][j])
			if !hasAdjacentSymbol {
				hasAdjacentSymbol = (i > 0 && (isSymbol(schematicMatrix[i-1][j]) || (j > 0 && isSymbol(schematicMatrix[i-1][j-1]) || (j+1 < columnsQuantity && isSymbol(schematicMatrix[i-1][j+1]))))) || (j > 0 && isSymbol(schematicMatrix[i][j-1])) || (i+1 < rowQuantity && (isSymbol(schematicMatrix[i+1][j]) || (j > 0 && isSymbol(schematicMatrix[i+1][j-1])) || (j+1 < columnsQuantity && isSymbol(schematicMatrix[i+1][j+1])))) || (j+1 < columnsQuantity && isSymbol(schematicMatrix[i][j+1]))
			}
		}

		if len(numberStr) > 0 {
			if hasAdjacentSymbol {
				convertedNumber, err := strconv.Atoi(numberStr)
				if err != nil {
					panic(err)
				}
				result += convertedNumber
			}
		}
	}

	fmt.Println(result)
}

func getAdjacentNumbersInRow(row []rune, asteriskColumn int) []string {
	var adjacentNumbers []string
	var numberStr string
	isAdjacentToAsterisk := false
	for column := 0; column < len(row); column++ {
		if isDigit(row[column]) {
			numberStr += string(row[column])
			if !isAdjacentToAsterisk {
				isAdjacentToAsterisk = column == asteriskColumn-1 || column == asteriskColumn || column == asteriskColumn+1
			}
			continue
		}

		if len(numberStr) > 0 && isAdjacentToAsterisk {
			adjacentNumbers = append(adjacentNumbers, numberStr)
		}

		numberStr = ""
		isAdjacentToAsterisk = false
	}
	if len(numberStr) > 0 && isAdjacentToAsterisk {
		adjacentNumbers = append(adjacentNumbers, numberStr)
	}
	return adjacentNumbers
}

func Day3Part2() {
	lines := strings.Split(utils.ReadInput(3), "\n")
	rowQuantity := len(lines)

	schematicMatrix := make([][]rune, rowQuantity)
	result := 0

	for rowNumber, line := range lines {
		schematicMatrix[rowNumber] = []rune(line)
	}

	for i := 0; i < len(schematicMatrix); i++ {
		for j := 0; j < len(schematicMatrix[i]); j++ {
			if schematicMatrix[i][j] != '*' {
				continue
			}

			var adjacentNumbers []string

			if i > 0 {
				adjacentNumbers = append(adjacentNumbers, getAdjacentNumbersInRow(schematicMatrix[i-1], j)...)
			}
			if i+1 < rowQuantity {
				adjacentNumbers = append(adjacentNumbers, getAdjacentNumbersInRow(schematicMatrix[i+1], j)...)
			}
			adjacentNumbers = append(adjacentNumbers, getAdjacentNumbersInRow(schematicMatrix[i], j)...)

			if len(adjacentNumbers) == 2 {
				number1, err1 := strconv.Atoi(adjacentNumbers[0])
				number2, err2 := strconv.Atoi(adjacentNumbers[1])
				if err1 != nil || err2 != nil {
					panic("error")
				}
				result += number1 * number2
			}
		}
	}

	fmt.Println(result)
}
