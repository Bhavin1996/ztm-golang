package main

import "fmt"

const (
	Low = iota
	Medium
	High
)

type PriorityQueqe[P comparable, v any] struct {
	items      map[P][]v
	priorities []P
}

func NewPriorityQueue[P comparable, v any](priorities []P) PriorityQueqe[P, v] {
	return PriorityQueqe[P, v]{items: make(map[P][]v), priorities: priorities}
}

func (pq *PriorityQueqe[P, v]) Add(priority P, value v) {
	pq.items[priority] = append(pq.items[priority], value)
}

func (pq *PriorityQueqe[P, v]) Next() (v, bool) {
	for i := 0; i < len(pq.priorities); i++ {
		priority := pq.priorities[i]
		items := pq.items[priority]
		if len(items) > 0 {
			next := items[0]
			pq.items[priority] = items[1:]
			return next, true
		}
	}
	var empty v
	return empty, false
}

func main() {
	queue := NewPriorityQueue[int, string]([]int{High, Medium, Low})
	queue.Add(Low, "L-1")
	queue.Add(High, "H-1")

	fmt.Println(queue.Next())

	queue.Add(Medium, "M-1")
	queue.Add(High, "H-2")
	queue.Add(High, "H-3")

	fmt.Println(queue.Next())

}
