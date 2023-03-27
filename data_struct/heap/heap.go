package heap

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func main() {

	h := &IntHeap{2, 1, 5, 13, 9, 0}
	heap.Init(h)
	fmt.Println(h) //&[0 1 2 13 9 5]
	heap.Push(h, 6)
	fmt.Printf("堆顶元素: %d\n", (*h)[0]) // 0
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	} //0 1 2 5 6 9 13
}
