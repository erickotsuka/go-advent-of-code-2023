package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erickotsuka/go-advent-of-code-2023/utils"
)

type gameRound map[string]int

func Day2Part1()  {
	lines := strings.Split(utils.ReadInput(2), "\n")
	result := 0
	for _, gameRecord := range lines {
		gameIdAndResults := strings.Split(gameRecord, ": ")
		gameNameAndId := strings.Split(gameIdAndResults[0], " ")
		gameId, err := strconv.Atoi(gameNameAndId[1])
		if err != nil {
			panic(err)
		}
		gameRounds := strings.Split(gameIdAndResults[1], "; ")
		possible := true
		for _, round := range gameRounds {
			cubeSets := strings.Split(round, ", ")	
			cubesGameRound := make(gameRound)
			for _, cubeSet := range cubeSets {
				quantityAndColor := strings.Split(cubeSet, " ")	
				quantity, error := strconv.Atoi(quantityAndColor[0])
				if error != nil {
					panic(error)
				}
				cubesGameRound[quantityAndColor[1]] = quantity 
			}
			possible = cubesGameRound["red"] <= 12 && cubesGameRound["green"] <= 13 && cubesGameRound["blue"] <= 14
			if !possible {
				break
			}
		}
		if !possible {
			continue
		}
		result += gameId
	}
	fmt.Println(result)
}

func Day2Part2() {
	lines := strings.Split(utils.ReadInput(2), "\n")
	result := 0
	for _, gameRecord := range lines {
		gameIdAndResults := strings.Split(gameRecord, ": ")
		gameRounds := strings.Split(gameIdAndResults[1], "; ")
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for _, round := range gameRounds {
			cubeSets := strings.Split(round, ", ")	
			cubesGameRound := make(gameRound)
			for _, cubeSet := range cubeSets {
				quantityAndColor := strings.Split(cubeSet, " ")	
				quantity, error := strconv.Atoi(quantityAndColor[0])
				if error != nil {
					panic(error)
				}
				cubesGameRound[quantityAndColor[1]] = quantity 
			}
			if cubesGameRound["red"] > maxRed {
				maxRed = cubesGameRound["red"]
			}
			if cubesGameRound["green"] > maxGreen {
				maxGreen = cubesGameRound["green"]
			}
			if cubesGameRound["blue"] > maxBlue {
				maxBlue = cubesGameRound["blue"]
			}
		}

		result += maxRed * maxGreen * maxBlue
	}
	fmt.Println(result)
}
