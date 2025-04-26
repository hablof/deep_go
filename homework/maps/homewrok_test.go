package main

import (
	"cmp"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type OrderedMap[K cmp.Ordered, V any] struct {
	m   map[K]V
	bst *node[K, V]
}

type node[K cmp.Ordered, V any] struct {
	key   K
	value V
	left  *node[K, V]
	right *node[K, V]
}

func (n *node[K, V]) forEach(action func(K, V)) {
	if n == nil {
		return
	}

	n.left.forEach(action)
	action(n.key, n.value)
	n.right.forEach(action)
}

func (n *node[K, V]) search(key K) (foundNode, parent *node[K, V]) {
	for n != nil {
		if n.key == key {
			return n, parent
		}

		if key < n.key {
			parent = n
			n = n.left
		} else {
			parent = n
			n = n.right
		}
	}

	return nil, parent
}

func (n *node[K, V]) insert(key K, val V) {
	for n != nil {
		if key < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}

	n = &node[K, V]{key: key, value: val}
}

func (n *node[K, V]) erase(key K) {
	target, parent := n.search(key)
	if target == nil {
		return
	}
	if target.right == nil {
		if parent == nil {
			n = target.left
		} else {
			if target == parent.left {
				parent.left = target.left
			} else {
				parent.right = target.left
			}
		}
	} else {
		leftmost := target.right
		parent = nil
		for leftmost.left != nil {
			parent = leftmost
			leftmost = leftmost.left
		}
		if parent != nil {
			parent.left = leftmost.right
		} else {
			target.right = leftmost.right
		}
		target.key = leftmost.key
		target.value = leftmost.value
	}
}

func NewOrderedMap[K cmp.Ordered, V any]() OrderedMap[K, V] {
	return OrderedMap[K, V]{
		m:   map[K]V{},
		bst: nil,
	}
}

func (m *OrderedMap[K, V]) Insert(key K, value V) {
	if _, ok := m.m[key]; ok {
		m.m[key] = value
		return //
	}

	m.m[key] = value
	_, parent := m.bst.search(key)
	if parent == nil {
		m.bst = &node[K, V]{key: key}
		return
	}

	if key < parent.key {
		parent.left = &node[K, V]{key: key}
		return
	}

	parent.right = &node[K, V]{key: key}
}

func (m *OrderedMap[K, V]) Erase(key K) {
	if _, ok := m.m[key]; !ok {
		return
	}

	delete(m.m, key)
	m.bst.erase(key)
}

func (m *OrderedMap[K, V]) Contains(key K) bool {
	_, ok := m.m[key]
	return ok
}

func (m *OrderedMap[K, V]) Size() int {
	return len(m.m)
}

func (m *OrderedMap[K, V]) ForEach(action func(K, V)) {
	m.bst.forEach(action)
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap[int, int]()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
