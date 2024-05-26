package kata

func SmallestIntegerFinder(numbers []int) (smol int) {
	smol = numbers[0]
	for _, i := range numbers {
		if i < smol {
			smol = i
		}
	}
	return smol
}
