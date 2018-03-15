package algorithms

func findSmallest(in []int) (smallest [2]int) {
	smallest[0] = in[0]
	smallest[1] = 0
	for i, ii := range in {
		if ii < smallest[0] {
			smallest[0] = ii
			smallest[1] = i
		}
	}
	return
}

// SelectionSort is
func SelectionSort(in []int) (out []int) {
	length := len(in)
	for i := 0; i < length; i++ {
		smallest := findSmallest(in)
		out = append(out, smallest[0])
		in = append(in[:smallest[1]], in[smallest[1]+1:]...)
	}
	return
}
