package treeUtils

import (
	"container/heap"
	"sort"
	"strings"
)

const nonElementRune = -1

type Heap []HuffTree

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].weight() < h[j].weight() }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	*h = append(*h, x.(HuffTree))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type HuffTree struct {
	left    *HuffTree
	right   *HuffTree
	element rune
	value   int
}

func (h HuffTree) weight() int {
	return h.value
}

func BuildPrefixTable(t *HuffTree, tble map[rune]string) {
	builder := strings.Builder{}

	var TraverseTree func(root *HuffTree, table map[rune]string, b *strings.Builder)
	TraverseTree = func(root *HuffTree, table map[rune]string, b *strings.Builder) {
		if root.left == nil {
			table[root.element] = b.String()
			return
		}

		(*b).WriteString("0")
		TraverseTree(root.left, tble, b)
		buf := (*b).String()
		buf = strings.TrimSuffix(buf, "0")
		(*b).Reset()
		(*b).WriteString(buf)
		(*b).WriteString("1")
		TraverseTree(root.right, tble, b)
		buf = (*b).String()
		buf = strings.TrimSuffix(buf, "1")
		(*b).Reset()
		(*b).WriteString(buf)
	}

	TraverseTree(t, tble, &builder)
}

func BuildHuffmanTree(m map[rune]int) HuffTree {
	arr := []HuffTree{}
	for key, value := range m {
		arr = append(arr, HuffTree{nil, nil, key, value})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].element < arr[j].element
	})

	h := &Heap{}
	heap.Init(h)
	for _, v := range arr {
		heap.Push(h, v)
	}

	for h.Len() > 1 {
		x := (*h)[0]
		heap.Pop(h)
		y := (*h)[0]
		heap.Pop(h)

		total := x.weight() + y.weight()
		heap.Push(h, HuffTree{&x, &y, nonElementRune, total})
	}

	return (*h)[0]
}
