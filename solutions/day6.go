package solutions

import "fmt"

type race struct {
	time uint
	recordDistance uint
}

func Day6Part1() {
	races := [4]race {
		{time: 41, recordDistance: 249},
		{time: 77, recordDistance: 1362},
		{time: 70, recordDistance: 1127},
		{time: 96, recordDistance: 1011},
	}
	
	result := 1

	for _, race := range races {
		numberOfWinningPossibilities := 0
		for holdingTime := uint(0); holdingTime <= race.time; holdingTime++ {
			remainingTime := race.time - holdingTime
			distance := holdingTime * remainingTime
			if distance > race.recordDistance {
				numberOfWinningPossibilities++
			} else if numberOfWinningPossibilities > 0 {
				break
			}
		}

		result *= numberOfWinningPossibilities
	}

	fmt.Println(result)
}

func Day6Part2() {
	race := race{time: 41777096, recordDistance: 249136211271011}

	result := 0
	for holdingTime := uint(0); holdingTime <= race.time; holdingTime++ {
		remainingTime := race.time - holdingTime
		distance := holdingTime * remainingTime
		if distance > race.recordDistance {
			result++
		} else if result > 0 {
			break
		}
	}

	fmt.Println(result)
}
