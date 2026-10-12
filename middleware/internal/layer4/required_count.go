package layer4

func RequiredCount(threshold int, eligible int) (int, bool) {
	if threshold < 0 {
		return (-threshold*eligible + 99) / 100, true
	}
	return threshold, threshold <= eligible
}

func RequiredFavourCount(threshold int, eligible int) int {
	needed, _ := RequiredCount(threshold, eligible)
	return min(needed, eligible)
}
