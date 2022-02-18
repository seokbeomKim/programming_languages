package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, node *tree) []int {
	if node != nil {
		values = appendValues(values, node.left)
		values = append(values, node.value)
		values = appendValues(values, node.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	val := []int{1, 2, 3, 4, 10, 8, 2, 4}
	Sort(val)
	fmt.Printf("%v\n", val)
}
