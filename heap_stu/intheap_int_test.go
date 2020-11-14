package heap_stu

import (
	"container/heap"
	"fmt"
	"testing"
)


type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}



func TestHeap(t *testing.T) {
	heap1 := &myHeap{1, 3, 2, 4, 8, 5, 9}
	heap.Fix(heap1, 1)
	fmt.Println(heap1)
	heap1.Push(0)
	heap.Fix(heap1, len(*heap1) - 1)
	fmt.Println(heap1)
	heap1.Push(6)
	heap.Fix(heap1, len(*heap1) - 1)
	fmt.Println(heap1)

	heap1.Push(7)
	heap.Fix(heap1, len(*heap1) - 1)
	fmt.Println(heap1)

	heap1.Push(-1)
	heap.Fix(heap1, len(*heap1) - 1)
	fmt.Println(heap1)

}

func TestHeap2(t *testing.T) {
	heap1 := &myHeap{3, 2, 4, 8, 1, 9}
	fmt.Println(heap1)
	heap.Init(heap1)
	fmt.Println(heap1)
	heap.Fix(heap1, 1)
	fmt.Println(heap1)

}