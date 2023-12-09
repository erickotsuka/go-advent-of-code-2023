package solutions

import (
	"fmt"
	"slices"
	"strings"

	"github.com/erickotsuka/go-advent-of-code-2023/utils"
)

type adjacentNodes struct {
	left string
	right string
}

func Day8Part1() {
	instructionsAndMap := strings.Split(utils.ReadInput(8), "\n\n")

	instructions := instructionsAndMap[0]
	nodesMap := instructionsAndMap[1]

	nodeToAdjacentsMap := make(map[string]adjacentNodes)

	nodesMapLines := strings.Split(nodesMap, "\n")

	for _, line := range nodesMapLines {
		nodeNameAndAdjacents := strings.Split(line, " = ")
		nodeName := nodeNameAndAdjacents[0]

		leftAndRightNodes := strings.Split(nodeNameAndAdjacents[1], ", ")
		nodeLeft := leftAndRightNodes[0][1:] 
		nodeRight := leftAndRightNodes[1][:len(leftAndRightNodes[1])-1]

		nodeToAdjacentsMap[nodeName] = adjacentNodes{left: nodeLeft, right: nodeRight}
	}

	initialNode := "AAA"
	finalNode := "ZZZ"
	currentNode := initialNode

	numberOfSteps :=0
	numberOfInstructions := len(instructions)

	for numberOfSteps = 0; currentNode != finalNode; numberOfSteps++ {
		instruction := instructions[numberOfSteps % numberOfInstructions]
		if instruction == 'L' {
			currentNode = nodeToAdjacentsMap[currentNode].left		 
		} else {
			currentNode = nodeToAdjacentsMap[currentNode].right	
		}
	}

	fmt.Println(numberOfSteps)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func Day8Part2() {
	instructionsAndMap := strings.Split(utils.ReadInput(8), "\n\n")

	instructions := instructionsAndMap[0]
	nodesMap := instructionsAndMap[1]

	nodeToAdjacentsMap := make(map[string]adjacentNodes)

	nodesMapLines := strings.Split(nodesMap, "\n")
	var startingNodes []string
	var endingNodes []string
	result := 1

	for _, line := range nodesMapLines {
		nodeNameAndAdjacents := strings.Split(line, " = ")
		nodeName := nodeNameAndAdjacents[0]
		lastChar := nodeName[len(nodeName)-1]
		if lastChar == 'A' {
			if len(startingNodes) == 0 {
				startingNodes = make([]string, 1)
				startingNodes[0] = nodeName
			} else {
				startingNodes = append(startingNodes, nodeName)
			}
		} else if lastChar == 'Z' {
			if len(endingNodes) == 0 {
				endingNodes = make([]string, 1)
				endingNodes[0] = nodeName
			} else {
				endingNodes = append(endingNodes, nodeName)
			}
		} 

		leftAndRightNodes := strings.Split(nodeNameAndAdjacents[1], ", ")
		nodeLeft := leftAndRightNodes[0][1:] 
		nodeRight := leftAndRightNodes[1][:len(leftAndRightNodes[1])-1]

		nodeToAdjacentsMap[nodeName] = adjacentNodes{left: nodeLeft, right: nodeRight}
	}

	startEndNodesStepsMatrix := make([][]int, len(startingNodes))

	for startingNodeIndex := 0; startingNodeIndex < len(startingNodes); startingNodeIndex++ {
		startEndNodesStepsMatrix[startingNodeIndex] = make([]int, len(endingNodes))
		for endingNodeIndex := 0; endingNodeIndex < len(endingNodes); endingNodeIndex++ {
			initialNode := startingNodes[startingNodeIndex] 
			finalNode := endingNodes[endingNodeIndex] 
			currentNode := initialNode

			numberOfSteps :=0
			numberOfInstructions := len(instructions)

			var visitedNodes []string

			for numberOfSteps = 0; currentNode != finalNode; numberOfSteps++ {
				instructionIndex  := numberOfSteps % numberOfInstructions
				instruction := instructions[instructionIndex]
				visitedNode := currentNode + string(fmt.Sprint(instructionIndex))
				if slices.Contains(visitedNodes, visitedNode) {
					numberOfSteps = -1
					break
				}
				visitedNodes = append(visitedNodes, visitedNode)
				if instruction == 'L' {
					currentNode = nodeToAdjacentsMap[currentNode].left		 
				} else {
					currentNode = nodeToAdjacentsMap[currentNode].right	
				}
			}

			startEndNodesStepsMatrix[startingNodeIndex][endingNodeIndex] = numberOfSteps
		}
	}

	numberOfStepsList := make([]int, len(startingNodes))
	index := 0

	for _, row := range startEndNodesStepsMatrix {
		for _, value := range row {
			if value != -1 {
				numberOfStepsList[index] = value
				index++
			}	
		}
	}
	
	result = lcm(numberOfStepsList[0], numberOfStepsList[1], numberOfStepsList[2:]...)

	fmt.Println(result)
}
