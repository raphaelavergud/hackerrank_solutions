package arrays

// this was so difficult and i did not like it very much
// i asked for help for this one and i still kind of dont understand it

// explanation of the problem:
// follow a set of instructions (queries) to place numbers into different lists. kind of like putting items on
// different shelves. they key was 'lastAnswer', which changes every time you look up an item, which then determines
// which list to use for the following instructions.

func dynamicArray(n int32, queries [][]int32) []int32 {
	sequences := make([][]int32, n)
	var lastAnswer int32 = 0
	var answers []int32

	queryIndex := 0
	for queryIndex < len(queries) {
		currentQuery := queries[queryIndex]
		queryType := currentQuery[0]
		xValue := currentQuery[1]
		yValue := currentQuery[2]

		sequenceIndex := (xValue ^ lastAnswer) % n

		if queryType == 1 {
			sequences[sequenceIndex] = append(sequences[sequenceIndex], yValue)

		} else {
			sequenceToQuery := sequences[sequenceIndex]

			sizeOfSequence := int32(len(sequenceToQuery))
			elementIndex := yValue % sizeOfSequence

			valueAtIndex := sequenceToQuery[elementIndex]
			lastAnswer = valueAtIndex

			answers = append(answers, lastAnswer)
		}

		queryIndex = queryIndex + 1
	}

	return answers
}
