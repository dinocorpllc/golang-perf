package hybrid_ranker

import "github.com/dinocorpllc/golang-perf"

type DoublyLinkedListNode struct {
	key  string
	val  uint64
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

// LinkedListRanker implements the synthetic workload using a singly linked list
type HybridRanker struct {
	root *DoublyLinkedListNode
	index map[string]*DoublyLinkedListNode
}

func (ll *HybridRanker) TopN(n uint64) []uint64 {
	cur := ll.root
	acc := ranker.NewSliceAccumulator(n)
	for cur != nil {
		acc.Offer(cur.val)
		cur = cur.next
	}
	return acc.Get()
}

func (ll *HybridRanker) Insert(key string, value uint64) {
	node := &DoublyLinkedListNode{
		next: ll.root,
		prev: nil,
		key:  key,
		val:  value,
	}
	ll.root = node

	if node.next != nil {
		node.next.prev = node
	}

	ll.index[key] = node
}

func (ll *HybridRanker) Get(key string) uint64 {
	node, _ := ll.index[key]
	if node != nil {
		return node.val
	}
	return 0
}

func (ll *HybridRanker) Delete(key string) bool {
	node, _ := ll.index[key]
	if node == nil {
		return false
	}
	prev := node.prev
	next := node.next
	if prev != nil {
		prev.next = next
	}
	delete(ll.index, key)
	return true
}
