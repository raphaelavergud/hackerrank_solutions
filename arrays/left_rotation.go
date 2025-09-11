package arrays

func rotateLeft(d int32, arr []int32) []int32 {
	// d is number of left rotations to do
	arrayLength := len(arr)

	if arrayLength < 2 {
		return arr
	}
	// why use modulo? because if d is 5 and len(arr) is also 5, then same arr again.
	// if d is 7 and arr is 5, i can save 5 rotations, and do 2 instead.

	// modulo, as a reminder, is '5 mod 2 = 1, wel 2 ind 5 2 mol ihnepasst und daenn abr 1 uebrigbliebt.'
	// also ka i ma 2 mol die rotation spara weil daenn eh nur wieda da selbe array ussakummt, stattdessen einfach 1 left rotation maha.
	numberOfRotations := int(d) % arrayLength

	// der index goht vo 0 bis numberOfRotations
	elementsToMove := arr[0:numberOfRotations]

	// der index nimmt da rest, falls numberOfRotations 2 isch, daenn ab index 2 (weils bei 0 startet)
	remainingElements := arr[numberOfRotations:arrayLength]

	// ... to unpack 'elementsToMove' because else it tries to append a list as an item into the other list (compiler error)
	rotatedArray := append(remainingElements, elementsToMove...)

	return rotatedArray
}
