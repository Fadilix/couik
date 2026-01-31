package game

func CountCorrect(results []bool) int {
	count := 0
	for _, current := range results {
		if current {
			count++
		}
	}
	return count
}

func CountIncorrect(results []bool) int {
	count := 0
	for _, current := range results {
		if !current {
			count++
		}
	}
	return count
}
