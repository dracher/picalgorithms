package algorithms

import "fmt"

// BinarySearch is
func BinarySearch(list []int, item int) (int, error) {
	start := 0
	end := len(list) - 1
	for start <= end {
		mid := (start + end) / 2
		if list[mid] == item {
			return mid, nil
		}
		if list[mid] > item {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1, fmt.Errorf("%d not in %v", item, list)
}
