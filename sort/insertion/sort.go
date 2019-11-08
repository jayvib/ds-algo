package insertion

func SortInt(l []int) {
	for i := 1; i < len(l); i++ {
		position := i
		tmp := l[i]
		for position > 0 && l[position-1] > tmp {
			l[position] = l[position-1]
			position--
		}
		l[position] = tmp
	}
}
