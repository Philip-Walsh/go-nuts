package kata

func Past(h, m, s int) int {
	return (h * 3.6e6) + (m * 60000) + (s * 1000)
}
