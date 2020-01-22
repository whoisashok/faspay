package str

func StringContains(slices []string, comparizon string) bool {
	for _, a := range slices {
		if a == comparizon {
			return true
		}
	}

	return false
}
