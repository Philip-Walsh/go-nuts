package kata

func QuarterOf(month int) int {
	if month > 6 {
		if month > 9 {
			return 4
		}
		return 3
	}
	if month > 3 {
		return 2
	}
	return 1
}
