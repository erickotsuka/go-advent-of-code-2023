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

func Day8Part2() {
	instructionsAndMap := strings.Split(utils.ReadInput(8), "\n\n")

	instructions := instructionsAndMap[0]
	nodesMap := instructionsAndMap[1]

	nodeToAdjacentsMap := make(map[string]adjacentNodes)

	nodesMapLines := strings.Split(nodesMap, "\n")
	var startingNodes []string
	var endingNodes []string

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

	fmt.Println(startingNodes)
	fmt.Println(endingNodes)

	startEndNodesStepsMatrix := make([][]int, len(startingNodes))

	for startingNodeIndex := 0; startingNodeIndex < len(startingNodes); startingNodeIndex++ {
		startEndNodesStepsMatrix[startingNodeIndex] = make([]int, len(endingNodes))
		for endingNodeIndex := 0; endingNodeIndex < len(endingNodes); endingNodeIndex++ {
			initialNode := startingNodes[startingNodeIndex] 
			finalNode := endingNodes[endingNodeIndex] 
			fmt.Println(initialNode + "->" + finalNode)
			currentNode := initialNode

			numberOfSteps :=0
			numberOfInstructions := len(instructions)

			var visitedNodes []string

			for numberOfSteps = 0; currentNode != finalNode; numberOfSteps++ {
				instruction := instructions[numberOfSteps % numberOfInstructions]
				visitedNode := currentNode + string(instruction)
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

	fmt.Println(startEndNodesStepsMatrix)
}
