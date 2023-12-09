package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/erickotsuka/go-advent-of-code-2023/utils"
)

func Day9Part1() {

	histories := strings.Split(utils.ReadInput(9), "\n")
	result := 0
	
	for _, history := range histories {
		valuesString := strings.Split(history, " ")
		values := make([]int, len(valuesString))
		for index, text := range valuesString {
			value, err := strconv.Atoi(text)
			if err != nil {
				panic(err)
			}
			values[index] = value
		}

		differencesMatrix := make([][]int, 1)
		differencesMatrix[0] = make([]int, len(values)-1)
		currentDifferencesRow := 0
		currentSequence := values
		for {
			allZeroes := true
			for i := 0; i < len(currentSequence)-1; i++ {
				difference := currentSequence[i+1] - currentSequence[i]
				differencesMatrix[currentDifferencesRow][i] = difference 
				allZeroes = allZeroes && (difference == 0)
			}
			if allZeroes {
				break
			}
			currentSequence = differencesMatrix[currentDifferencesRow]
			currentDifferencesRow++
			differencesMatrix = append(differencesMatrix, make([]int, len(currentSequence)-1)) 
		}
		for row := len(differencesMatrix) - 1; row >= 1; row-- {
			lastValueCurrentRow := differencesMatrix[row][len(differencesMatrix[row])-1]
			lastValueRowAbove := differencesMatrix[row-1][len(differencesMatrix[row-1])-1]
			differencesMatrix[row-1] = append(differencesMatrix[row-1], lastValueCurrentRow + lastValueRowAbove)
		}
		nextValue := differencesMatrix[0][len(differencesMatrix[0])-1] + values[len(values)-1]
		result += nextValue
	}
	fmt.Println(result)
}
 func Day9Part2() {
	histories := strings.Split(utils.ReadInput(9), "\n")
	result := 0
	
	for _, history := range histories {
		valuesString := strings.Split(history, " ")
		slices.Reverse(valuesString)
		values := make([]int, len(valuesString))
		for index, text := range valuesString {
			value, err := strconv.Atoi(text)
			if err != nil {
				panic(err)
			}
			values[index] = value
		}

		differencesMatrix := make([][]int, 1)
		differencesMatrix[0] = make([]int, len(values)-1)
		currentDifferencesRow := 0
		currentSequence := values
		for {
			allZeroes := true
			for i := 0; i < len(currentSequence)-1; i++ {
				difference := currentSequence[i] - currentSequence[i+1]
				differencesMatrix[currentDifferencesRow][i] = difference 
				allZeroes = allZeroes && (difference == 0)
			}
			if allZeroes {
				break
			}
			currentSequence = differencesMatrix[currentDifferencesRow]
			currentDifferencesRow++
			differencesMatrix = append(differencesMatrix, make([]int, len(currentSequence)-1)) 
		}
		for row := len(differencesMatrix) - 1; row >= 1; row-- {
			lastValueCurrentRow := differencesMatrix[row][len(differencesMatrix[row])-1]
			lastValueRowAbove := differencesMatrix[row-1][len(differencesMatrix[row-1])-1]
			differencesMatrix[row-1] = append(differencesMatrix[row-1], lastValueRowAbove - lastValueCurrentRow)
		}
		nextValue := values[len(values)-1] - differencesMatrix[0][len(differencesMatrix[0])-1]
		result += nextValue
	}
	fmt.Println(result)
}
