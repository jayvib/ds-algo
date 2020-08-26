package hashtable

func FindPair(ints []int, value int) ([2]int, bool) {
	m := make(map[int]interface{})
	size := len(ints)
	for i := 0; i < size; i++ {
		d := diff(value, ints, i)
		if _, ok := m[d]; ok {
			return [2]int{d, ints[i]}, true
		}
		m[ints[i]] = true
	}
	return [2]int{}, false
}

func diff(value int, ints []int, i int) int {
	return value - ints[i]
}
