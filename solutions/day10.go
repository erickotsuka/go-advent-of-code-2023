package solutions

import (
	"fmt"
	"slices"
	"strings"

	"github.com/erickotsuka/go-advent-of-code-2023/utils"
)

type pipe struct {
	format byte
	row int
	col int
	distanceFromStart int
}


type direction int

const (
	north direction = iota
	south
	west
	east
)

var directionToPossibleConnectFormatsMap = map[direction][]byte{
	north: {'|', '7', 'F'},
	south: {'|', 'L', 'J'},
	west: {'-', 'L', 'F'},
	east: {'-', 'J', '7'},
}

var formatToPossibleDirectionsMap = map[byte][]direction{
	'|': {north, south},
	'-': {west, east},
	'L': {north, east},
	'J': {north, west},
	'7': {west, south},
	'F': {east, south},
}


func getStartingPipeFormat(startingRow int, startingCol int, lines []string) byte {
	connectedDirections := make([]direction, 0)
	if startingRow > 0 {
		pipeFormatNorth := lines[startingRow - 1][startingCol]
		if slices.Contains(directionToPossibleConnectFormatsMap[north], pipeFormatNorth) {
			connectedDirections = append(connectedDirections, north)
		}
	}
	if startingRow < len(lines) - 1 {
		pipeFormatSouth := lines[startingRow + 1][startingCol]
		if slices.Contains(directionToPossibleConnectFormatsMap[south], pipeFormatSouth) {
			connectedDirections = append(connectedDirections, south)
		}
	}
	if startingCol > 0 {
		pipeFormatWest := lines[startingRow][startingCol - 1]
		if slices.Contains(directionToPossibleConnectFormatsMap[west], pipeFormatWest) {
			connectedDirections = append(connectedDirections, west)
		}
	}
	if startingCol < len(lines[startingRow]) - 1 {
		pipeFormatEast := lines[startingRow][startingCol + 1]
		if slices.Contains(directionToPossibleConnectFormatsMap[east], pipeFormatEast) {
			connectedDirections = append(connectedDirections, east)
		}
	}
	if len(connectedDirections) != 2 {
		panic("More than 2 connected directions in S")
	}

	if connectedDirections[0] == north && connectedDirections[1] == south {
		return '|'
	}
	if connectedDirections[0] == north && connectedDirections[1] == west {
		return 'J'
	}
	if connectedDirections[0] == north && connectedDirections[1] == east {
		return 'L'
	}
	if connectedDirections[0] == south && connectedDirections[1] == west {
		return '7'
	}
	if connectedDirections[0] == south && connectedDirections[1] == east {
		return 'F'
	}

	return '-'
}

func dfs(lines []string, checkedPipes []pipe, rootPipe pipe) {
	connectedPipes := make([]pipe, 0)
	possibleConnectedDirections := formatToPossibleDirectionsMap[rootPipe.format]
	if rootPipe.row > 0 && slices.Contains(possibleConnectedDirections, north) {
		pipeFormatNorth := lines[rootPipe.row -1][rootPipe.col]	
		if slices.Contains(directionToPossibleConnectFormatsMap[north], pipeFormatNorth) { 
			pipeAbove := pipe {
				row: rootPipe.row - 1,
				col: rootPipe.col,
				distanceFromStart: rootPipe.distanceFromStart + 1,
				format: pipeFormatNorth,
			}
			connectedPipes = append(connectedPipes, pipeAbove)
		}
	}
	if rootPipe.col > 0 && slices.Contains(possibleConnectedDirections, west) {
		pipeFormatWest := lines[rootPipe.row][rootPipe.col - 1]
		if slices.Contains(directionToPossibleConnectFormatsMap[west], pipeFormatWest) { 
			pipeWest := pipe {
				row: rootPipe.row,
				col: rootPipe.col - 1,
				distanceFromStart: rootPipe.distanceFromStart + 1,
				format: pipeFormatWest,
			}
			connectedPipes = append(connectedPipes, pipeWest)
		}
	}
	if rootPipe.row < len(lines) - 1 && slices.Contains(possibleConnectedDirections, south) {
		pipeFormatSouth := lines[rootPipe.row + 1][rootPipe.col]
		if slices.Contains(directionToPossibleConnectFormatsMap[south], pipeFormatSouth) {
			pipeSouth := pipe {
				row: rootPipe.row + 1,
				col: rootPipe.col,
				distanceFromStart: rootPipe.distanceFromStart + 1,
				format: pipeFormatSouth,
			}
			connectedPipes = append(connectedPipes, pipeSouth)
		}
	}
	if rootPipe.col < len(lines[rootPipe.row]) - 1 && slices.Contains(possibleConnectedDirections, east) {
		pipeFormatEast := lines[rootPipe.row][rootPipe.col + 1]
		if slices.Contains(directionToPossibleConnectFormatsMap[east], pipeFormatEast) {
			pipeEast := pipe {
				row: rootPipe.row,
				col: rootPipe.col + 1,
				distanceFromStart: rootPipe.distanceFromStart + 1,
				format: pipeFormatEast,
			}
			connectedPipes = append(connectedPipes, pipeEast)
		}
	}

	for _, connectedPipe := range connectedPipes {
		dfs(lines, checkedPipes, connectedPipe)	
	}
	checkedPipes = append(checkedPipes, rootPipe)
}

func Day10Part1() {
	lines := strings.Split(utils.ReadInput(10), "\n")	
	startingLine, startingColumn := 0, 0
	for lineIndex, line := range lines {
		start := strings.Index(line, "S") 
		if start != -1 {
			startingLine = lineIndex
			startingColumn = start
			break
		}
	}

	fmt.Println(string(getStartingPipeFormat(startingLine, startingColumn, lines)))

}
