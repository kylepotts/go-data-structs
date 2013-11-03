package main

import "fmt"

type heap struct {
	array []int
	size  int
}

func newHeap(size int) *heap {
	h := new(heap)
	h.array = make([]int, size)
	h.size = 0
	return h

}

func (h *heap) less(i int, j int) bool {
	return h.array[i] < h.array[j]

}

func (h *heap) swap(i int, j int) {
	temp := h.array[i]
	h.array[i] = h.array[j]
	h.array[j] = temp

}

func (h *heap) up(j int) {
	for {
		i := j / 2
		if i == j || h.less(i, j) {
			break
		}
		h.swap(i, j)
		j = i

	}

}

func (h *heap) down(i int) {
	for {
		left := 2 * i
		if left > h.size {
			break
		}
		j := left
		if right := left + 1; right < h.size && !h.less(left, right) {
			j = right
		}

		if h.less(i, j) {
			break
		}

		h.swap(i, j)
		i = j

	}

}

func (h *heap) extractMin() int {
	val := h.array[1]
	h.swap(1, h.size)
	h.size--
	h.down(1)
	return val

}

func (h *heap) insert(val int) {
	h.size++
	h.array[h.size] = val
	h.up(h.size)

}

func (h *heap) printHeap() {
	for i := 1; i <= h.size; i++ {
		fmt.Printf("%d\n", h.array[i])
	}

}

func main() {
	h := newHeap(200)
	for i := 100; i > 0; i-- {
		h.insert(i)

	}

	for i := 100; i > 0; i-- {
		a := h.extractMin()
		fmt.Printf("%d\n", a)
	}

}
