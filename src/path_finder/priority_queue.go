package pathfinder

type PriorityElement struct {
	Priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Nodes.  The
// PriorityQueue is used to track open nodes by rank.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest priority since we don't know the size of the path so we use lower than here.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// No verirication at compile time, this is pretty bad...
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	no := x.(*Node)
	no.Index = n
	*pq = append(*pq, no)
}

// No verirication at compile time, this is pretty bad...
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	no := old[n-1]
	no.Index = -1
	*pq = old[0 : n-1]
	return no
}
