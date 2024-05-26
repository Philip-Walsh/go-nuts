package kata

func HowMuchILoveYou(i int) string {
	var opts = []string{
		"I love you",
		"a little",
		"a lot",
		"passionately",
		"madly",
		"not at all",
	}
	index := (i - 1) % len(opts)
	return opts[index]
}
