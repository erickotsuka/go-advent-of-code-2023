package solutions

import (
	"fmt"
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
