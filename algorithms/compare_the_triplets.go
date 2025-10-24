package algorithms

func compareTriplets(a []int32, b []int32) []int32 {
	var aliceScore int32 = 0
	var bobScore int32 = 0

	var categoryIndex = 0

	for categoryIndex < 3 {

		var aliceRating = a[categoryIndex]
		var bobRating = b[categoryIndex]

		if aliceRating > bobRating {
			aliceScore = aliceScore + 1
		} else if aliceRating < bobRating {
			bobScore = bobScore + 1
		}
		categoryIndex = categoryIndex + 1
	}
	var finalScores = []int32{aliceScore, bobScore}

	return finalScores
}

// more readable: refactor

func compareTriplets2(a []int32, b []int32) []int32 {

	aliceScore := int32(0)
	bobScore := int32(0)

	for categoryIndex := 0; categoryIndex < 3; categoryIndex++ {

		aliceRating := a[categoryIndex]
		bobRating := b[categoryIndex]

		if aliceRating > bobRating {
			aliceScore = aliceScore + 1
			continue
		}
		if aliceRating < bobRating {
			bobScore = bobScore + 1
			continue
		}

	}
	var finalScores []int32 = []int32{aliceScore, bobScore}

	return finalScores
}

// most readable:

func compareTriplets3(a []int32, b []int32) []int32 {

	aliceScore := int32(0)
	bobScore := int32(0)

	for categoryIndex := 0; categoryIndex < 3; categoryIndex++ {

		aliceRating := a[categoryIndex]
		bobRating := b[categoryIndex]

		if aliceRating == bobRating {
			continue
		}

		switch result := aliceRating > bobRating; result {
		case true:
			aliceScore += 1
		case false:
			bobScore += 1
		}
	}
	var finalScores []int32 = []int32{aliceScore, bobScore}

	return finalScores
}
