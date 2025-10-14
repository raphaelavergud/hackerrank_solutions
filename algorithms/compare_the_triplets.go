package algorithms

func compareTriplets(a []int32, b []int32) []int32 {
	var aliceScore int32 = 0
	var bobScore int32 = 0

	var categoryIndex int = 0

	for categoryIndex < 3 {

		var aliceRating int32 = a[categoryIndex]
		var bobRating int32 = b[categoryIndex]

		if aliceRating > bobRating {
			aliceScore = aliceScore + 1
		} else if bobRating > aliceRating {
			bobScore = bobScore + 1
		}
		categoryIndex = categoryIndex + 1
	}
	var finalScores []int32 = []int32{aliceScore, bobScore}

	return finalScores
}
