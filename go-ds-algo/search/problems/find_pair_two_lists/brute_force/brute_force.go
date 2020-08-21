package brute_force

func FindPair(setA, setB []int, value int) (out [][2]int, ok bool) {
	for i := 0; i < len(setA); i++ {
		for j := 0; j < len(setB); j++ {
			x, y := setA[i], setB[j]
			if (x + y) == value {
				out = append(out, [2]int{x, y})
				ok = true
			}
		}
	}
	return
}
