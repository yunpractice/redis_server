package main

import (
	"math/rand"
)

type SkipListNode struct {
	next *SkipListNode
	pre  *SkipListNode

	forward []*SkipListNode

	key   interface{}
	value interface{}
}

type CMP func(a interface{}, b interface{}) int

type SkipList struct {
	head     *SkipListNode
	tail     *SkipListNode
	len      int
	maxlevel int
	cmp      CMP
}

func (sl *SkipList) Length() int {
	return sl.len
}

func (sl *SkipList) Empty() bool {
	return sl.len == 0
}

func (sl *SkipList) Forward(current *SkipListNode, path []*SkipListNode, key interface{}) *SkipListNode {
	if sl.len == 0 {
		return nil
	}
	depth := len(current.forward) - 1

	for i := depth; i >= 0; i-- {
		for current.forward[i] != nil && sl.cmp(current.key, key) < 0 {
			current = current.forward[i]
		}
		if path != nil {
			path[i] = current
		}
	}
	return current.next
}

func (sl *SkipList) GetRandomLevel() int {
	n := 0
	for ; n <= sl.maxlevel && rand.Float32() < 0.25; n++ {

	}
	return n
}

func (sl *SkipList) CreateNode(ishead bool, istail bool, key interface{}, value interface{}) *SkipListNode {
	depth := 0
	if ishead {
		depth = sl.maxlevel
	} else if !istail {
		depth = sl.GetRandomLevel()
	}

	node := &SkipListNode{
		key:     key,
		value:   value,
		forward: make([]*SkipListNode, depth),
	}
	return node
}

func (sl *SkipList) Add(key interface{}, value interface{}) bool {
	if sl.len == 0 {
		sl.head = sl.CreateNode(true, false, key, value)
		sl.tail = sl.head
		sl.len = 1
		return true
	}

	// 插入头部
	if sl.cmp(key, sl.head.key) < 0 {
		old := sl.head.key
		oldvalue := sl.head.value
		sl.head.key = key
		sl.head.value = value
		key = old
		value = oldvalue
		// 插入尾部
	} else if sl.len > 1 && sl.cmp(key, sl.tail.key) > 0 {
		old := sl.tail.key
		oldvalue := sl.tail.value
		sl.tail.key = key
		sl.tail.value = value
		key = old
		value = oldvalue
	}

	if sl.len == 1 {
		if sl.cmp(sl.head.key, key) == 0 {

		} else {
			newnode := sl.CreateNode(false, true, key, value)
			for i := sl.maxlevel - 1; i >= 0; i-- {
				sl.head.forward[i] = newnode
				newnode.pre = sl.tail
				sl.tail.next = newnode
				sl.tail = newnode
			}
			sl.len++
		}

		return true
	}

	path := make([]*SkipListNode, sl.maxlevel)
	node := sl.Forward(sl.head, path, key)
	if node != nil && node.key == key {
		return true
	}

	newnode := sl.CreateNode(false, false, key, value)
	for i := len(newnode.forward) - 1; i >= 0; i-- {
		newnode.forward[i] = path[i].forward[i]
		path[i].forward[i] = newnode
	}
	newnode.pre = node.pre
	node.pre.next = newnode
	sl.len++
	return true
}

func (sl *SkipList) Get(key interface{}) interface{} {
	if sl.len == 0 {
		return nil
	}

	node := sl.Forward(sl.head, nil, key)
	if node == nil || sl.cmp(node.key, key) != 0 {
		return nil
	}

	return node.value
}

func (sl *SkipList) DeleteKey(key interface{}) {
	path := make([]*SkipListNode, sl.maxlevel)

	node := sl.Forward(sl.head, path, key)

	if node == nil || node.key != key {
		return
	}

	sl.len--

	if sl.len == 0 {
		sl.head = nil
		sl.tail = nil
		return
	} else if sl.len == 1 {
		if node == sl.tail {
			node.key = sl.head.key
			node.value = sl.head.value

		}
		sl.head = sl.tail
		sl.head.next = nil
		sl.head.pre = nil
		return
	}

	pre := node.pre
	next := node.next
	if node == sl.head {
		for i := 0; i < len(next.forward); i++ {
			if node.forward[i] == next {
				node.forward[i] = next.forward[i]
			}
		}
		next.forward = node.forward
		sl.head = next
		next.pre = nil
	} else {
		for i := sl.maxlevel - 1; i >= 0; i-- {
			if path[i].forward[i] == node {
				if node == sl.tail {
					path[i].forward[i] = pre
				} else {
					path[i].forward[i] = node.forward[i]
				}
			}
		}
		pre.next = next
		if next != nil {
			next.pre = pre
		}
		if node == sl.tail {
			sl.tail = pre
			pre.forward = node.forward
		}
	}
}

func (sl *SkipList) DeleteValue(key interface{}, value interface{}) {

}
