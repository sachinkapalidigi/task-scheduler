package collection

type PriorityTask struct {
	Priority int64
	Task     func(interface{})
}

type PriorityTaskHeap []PriorityTask

func (h PriorityTaskHeap) Len() int           { return len(h) }
func (h PriorityTaskHeap) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h PriorityTaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PriorityTaskHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(PriorityTask))
}

func (h *PriorityTaskHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h PriorityTaskHeap) Peek() PriorityTask {
	return h[0]
}

func (h PriorityTaskHeap) Empty() bool {
	return len(h) == 0
}
