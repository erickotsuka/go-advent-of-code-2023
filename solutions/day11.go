package solutions

import (
	"fmt"
	"slices"
	"strings"

	"github.com/erickotsuka/go-advent-of-code-2023/utils"
)

func getRowsWithoutGalaxies(lines []string) []int {
	rowsWithoutGalaxies := make([]int, 0)
	for rowNumber, line := range lines {
		if !strings.Contains(line, "#") {
			rowsWithoutGalaxies = append(rowsWithoutGalaxies, rowNumber)
		}
	}
	return rowsWithoutGalaxies
}

func getColumnsWithoutGalaxies(lines []string) []int {
	columnsWithoutGalaxies := make([]int, 0)
	for columnNumber := 0; columnNumber < len(lines[0]); columnNumber++ {
		rowNumber := 0
		for rowNumber = 0; rowNumber < len(lines); rowNumber++ {
			if lines[rowNumber][columnNumber] == '#' {
				break
			}
		}

		if rowNumber == len(lines) && lines[rowNumber-1][columnNumber] != '#' {
			columnsWithoutGalaxies = append(columnsWithoutGalaxies, columnNumber)
		}
	}
	return columnsWithoutGalaxies
}

func getExpandedUniverse(lines []string) []string {
	rowsWithoutGalaxies := getRowsWithoutGalaxies(lines)
	columnsWithoutGalaxies := getColumnsWithoutGalaxies(lines)
	expandedUniverse := make([]string, 0)
	for rowNumber, line := range lines {
		expandedUniverseLine := ""
		for columnNumber := 0; columnNumber < len(line); columnNumber++ {
			expandedUniverseLine = expandedUniverseLine + string(line[columnNumber])
			if slices.Contains(columnsWithoutGalaxies, columnNumber) {
				expandedUniverseLine = expandedUniverseLine + string(line[columnNumber])
			}
		}
		expandedUniverse = append(expandedUniverse, expandedUniverseLine)
		if slices.Contains(rowsWithoutGalaxies, rowNumber) {
			expandedUniverse = append(expandedUniverse, expandedUniverseLine)
		}
	}
	return expandedUniverse
}

type position struct {
	row int
	col int
}

func getGalaxyPositions(universe []string) []position {
	galaxyPositions := make([]position, 0)
	for rowNumber, row := range universe {
		for columnNumber, column := range row {
			if column == '#' {
				galaxyPositions = append(galaxyPositions, position{row: rowNumber, col: columnNumber})
			}
		}
	}
	return galaxyPositions
}

func Day11Part1() {
	lines := strings.Split(utils.ReadInput(11), "\n")
	expandedUniverse := getExpandedUniverse(lines)
	galaxyPositions := getGalaxyPositions(expandedUniverse)
	result := 0
	for index, position := range galaxyPositions[:len(galaxyPositions)-1] {
		for j := index + 1; j < len(galaxyPositions); j++ {
			horizontalDistance := position.row - galaxyPositions[j].row
			if horizontalDistance < 0 {
				horizontalDistance = -horizontalDistance
			}
			verticalDistance := position.col - galaxyPositions[j].col
			if verticalDistance < 0 {
				verticalDistance = -verticalDistance
			}
			distance := horizontalDistance + verticalDistance
			result += distance
		}
	}
	fmt.Println(result)
}
