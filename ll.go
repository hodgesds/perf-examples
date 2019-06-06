package main

import "math/rand"

type node struct {
	val        int64
	prev, next *node
}

type ll struct {
	head *node
}

func (l *ll) removeDuplicates() {
	set := map[int64]bool{}
	cur := l.head
	for cur != nil {
		seen, found := set[cur.val]
		if !found {
			set[cur.val] = true
			cur = cur.next
			continue
		}
		// If the value has been seen, then delete it.
		if seen {
			cur.prev.next = cur.next
		}
		cur = cur.next
	}
}

func testLL(count int) *ll {
	l := &ll{
		head: &node{val: 1},
	}
	cur := l.head
	for i := 0; i < count; i++ {
		newNode := &node{
			val:  rand.Int63(),
			prev: cur,
		}
		cur = newNode
	}
	return l
}
