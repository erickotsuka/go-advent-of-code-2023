package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erickotsuka/go-advent-of-code-2023/utils"
)

type interval struct {
	begin uint
	end uint
}

func Day5Part1() {
	sections := strings.Split(utils.ReadInput(5), "\n\n")
	seeds := strings.Split(sections[0], " ")[1:]

	var intervalToIntervalMaps []map[interval]interval

	for _, section := range sections[1:] {
		intervalToIntervalMap := make(map[interval]interval)
		entries := strings.Split(section, "\n")[1:]
		for _, entry := range entries {
			values := strings.Split(entry, " ")
			destinationRangeStart, err1 := strconv.Atoi(values[0])
			sourceRangeStart, err2 := strconv.Atoi(values[1])
			rangeSize, err3 := strconv.Atoi(values[2])
			
			if err1 != nil || err2 != nil || err3 != nil {
				panic("error atoi seed to soil")
			}

			destinationInterval := interval{begin: uint(destinationRangeStart), end: uint(destinationRangeStart + rangeSize - 1)}
			sourceInterval := interval{begin: uint(sourceRangeStart), end: uint(sourceRangeStart + rangeSize - 1)}
			intervalToIntervalMap[sourceInterval] = destinationInterval
		}
		intervalToIntervalMaps = append(intervalToIntervalMaps, intervalToIntervalMap)
	}

	const maxUint = ^uint(0) 

	smallestLocationNumber := maxUint
	
	for _, seed := range seeds {
		seedNumber, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}

		sourceNumber := uint(seedNumber)
		destinationNumber := sourceNumber

		for _, intervalToIntervalMap := range intervalToIntervalMaps {
			for sourceInterval, destinationInterval := range intervalToIntervalMap {
				if sourceInterval.begin <= sourceNumber && sourceNumber <= sourceInterval.end {
					distance := sourceNumber - sourceInterval.begin
					destinationNumber = destinationInterval.begin + distance
					sourceNumber = destinationNumber
					break
				}	
			}
		}

		if destinationNumber < smallestLocationNumber {
			smallestLocationNumber = destinationNumber	
		}

	}

	fmt.Println(smallestLocationNumber)
}

func Day5Part2() {
	sections := strings.Split(utils.ReadInput(5), "\n\n")
	seedIntervalFields := strings.Split(sections[0], " ")[1:]

	var seedIntervals []interval

	for i := 0; i < len(seedIntervalFields); i +=2 {
		intervalStart, err1 := strconv.Atoi(seedIntervalFields[i])
		intervalSize, err2 := strconv.Atoi(seedIntervalFields[i + 1])
		if err1 != nil || err2 != nil {
			panic("error atoi seeds")
		}
		seedInterval := interval{begin: uint(intervalStart), end: uint(intervalStart + intervalSize - 1)}
		seedIntervals = append(seedIntervals, seedInterval)
	}

	var intervalToIntervalMaps []map[interval]interval

	for _, section := range sections[1:] {
		intervalToIntervalMap := make(map[interval]interval)
		entries := strings.Split(section, "\n")[1:]
		for _, entry := range entries {
			values := strings.Split(entry, " ")
			destinationRangeStart, err1 := strconv.Atoi(values[0])
			sourceRangeStart, err2 := strconv.Atoi(values[1])
			rangeSize, err3 := strconv.Atoi(values[2])
			
			if err1 != nil || err2 != nil || err3 != nil {
				panic("error atoi entries loop")
			}

			destinationInterval := interval{begin: uint(destinationRangeStart), end: uint(destinationRangeStart + rangeSize - 1)}
			sourceInterval := interval{begin: uint(sourceRangeStart), end: uint(sourceRangeStart + rangeSize - 1)}
			intervalToIntervalMap[sourceInterval] = destinationInterval
		}
		intervalToIntervalMaps = append(intervalToIntervalMaps, intervalToIntervalMap)
	}

	const maxUint = ^uint(0) 

	smallestLocationNumber := maxUint

	sourceIntervals := seedIntervals
	var destinationIntervals []interval

	for _, intervalToIntervalMap := range intervalToIntervalMaps {
		for sourceInterval, destinationInterval := range intervalToIntervalMap {
			for _, checkingSourceInterval := range sourceIntervals {
				// TODO: finish this	
			}
		}	
	}
	
			

	fmt.Println(smallestLocationNumber)
}