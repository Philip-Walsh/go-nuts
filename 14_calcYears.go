package kata

func CalculateYears(years int) (result [3]int) {
	var cYears = 0
	var dYears = 0

	for i := 0; i <= years; i++ {
		if i == 1 {
			cYears += 15
			dYears += 15
		} else if i == 2 {
			cYears += 9
			dYears += 9
		} else if i > 2 {
			cYears += 4
			dYears += 5
		}
	}
	return [3]int{years, cYears, dYears}
}
