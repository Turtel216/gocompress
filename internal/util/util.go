package util

// Utility function to get the last value of a map
func GetLastValueOfMap(_map map[string]int) int {
	var last_value int
	for _, value := range _map {
		last_value = value
	}

	return last_value
}
