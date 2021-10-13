package calculate

type PriorityQueue []AxieChild

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Chance Descending, PartID Ascending
	if pq[i].Chance != pq[j].Chance {
		return pq[i].Chance > pq[j].Chance
	}
	if pq[i].Mouth != pq[j].Mouth {
		return pq[i].Mouth < pq[j].Mouth
	}
	if pq[i].Horn != pq[j].Horn {
		return pq[i].Horn < pq[j].Horn
	}
	if pq[i].Back != pq[j].Back {
		return pq[i].Back < pq[j].Back
	}
	return pq[i].Tail < pq[j].Tail
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	axieChild := x.(AxieChild)
	*pq = append(*pq, axieChild)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	axieChild := old[n-1]
	old[n-1] = AxieChild{} // avoid memory leak
	*pq = old[0 : n-1]
	return axieChild
}
