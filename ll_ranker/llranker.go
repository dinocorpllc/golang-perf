package ll_ranker

import "github.com/dinocorpllc/golang-perf"

type LinkedListNode struct {
	key  string
	val  uint64
	next *LinkedListNode
}

// LinkedListRanker implements the synthetic workload using a singly linked list
type LinkedListRanker struct {
	root *LinkedListNode
}

func (ll *LinkedListRanker) TopN(n uint64) []uint64 {
	cur := ll.root
	acc := ranker.NewSliceAccumulator(n)
	for cur != nil {
		acc.Offer(cur.val)
		cur = cur.next
	}
	return acc.Get()
}

func (ll *LinkedListRanker) Insert(key string, value uint64) {
	node := &LinkedListNode{
		next: ll.root,
		key:  key,
		val:  value,
	}
	ll.root = node
}

func (ll *LinkedListRanker) Get(key string) uint64 {
	cur := ll.root
	for cur != nil {
		if cur.key == key {
			return cur.val
		}
		cur = cur.next
	}
	return 0
}

func (ll *LinkedListRanker) Delete(key string) bool {
	var prev, cur *LinkedListNode

	if ll.root == nil {
		return true
	} else if ll.root.key == key {
		ll.root = ll.root.next
		return true
	}

	prev = ll.root
	cur = ll.root.next
	for cur != nil {
		if cur.key == key {
			prev.next = cur.next
			return true
		}
		cur = cur.next
	}
	return true
}
