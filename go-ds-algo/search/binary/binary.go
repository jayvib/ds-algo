package binary

import "sort"

func Search(data []int, value int) bool {
	size := len(data)
	var low, mid, high int
	low = 0
	high = size-1
	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		} else if data[mid] > value {
			high = mid-1
		} else if data[mid] < value {
			low = mid+1
		}
	}
	return false
}

type Keyer interface {
	Key() int
}

type Product struct {
	ID int
	Name string
}

func (p *Product) Key() int {
	return p.ID
}

func SearchProducts(products []Keyer, id int) (item Keyer, found bool) {

	// Check first if sorted
	if !isSorted(products) {
		sort.Sort(keyers(products))
	}

	size := len(products)
	low := 0
	high := size-1
	mid := 0

	for low <= high {
		mid = low + (high-low)/2
		if products[mid].Key() == id {
			return products[mid], true
		} else if products[mid].Key() > id {
			high = mid-1
		} else if products[mid].Key() < id {
			low = mid+1
		}
	}
	return
}

func isSorted(product []Keyer) bool {
	return sort.IsSorted(keyers(product))
}

type keyers []Keyer
func(k keyers) Len() int { return len(k) }
func(k keyers) Less(i, j int) bool { return k[i].Key() < k[j].Key()}
func(k keyers) Swap(i, j int) { k[i], k[j] = k[j], k[i]}


