package main

type IntegerHeap []int

func (h IntegerHeap) Len() int { return len(h) }
func (h IntegerHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntegerHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntegerHeap) Push(heapInt interface{}) {
	*h = append(*h, heapInt.(int))
}

func main() {

}
