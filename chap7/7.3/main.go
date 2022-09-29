package main

import (
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
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

func (t *tree) String() string {
	var v string
	if t == nil {
		return v
	}
	v = v + t.left.String()
	v = v + strconv.Itoa(t.value)
	v = v + t.right.String()
	return v

}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t == nil {
		return values
	}

	values = appendValues(values, t.left)
	fmt.Printf("adding %d \n", t.value)
	values = append(values, t.value)
	fmt.Printf("adding %d \n", values)
	values = appendValues(values, t.right)
	return values
}

func main() {
	t := add(nil, 3)
	fmt.Println("values %s", t.String)
	t = add(t, 2)
	fmt.Println("values %s", t.String)
	t = add(t, 1)
	fmt.Println("values %s", t.String)
	t = add(t, 4)
	fmt.Println("values %s", t.String)

	fmt.Printf("....first.. %d \n", t.value)
	r := make([]int, 0)
	r = appendValues(r, t)
	fmt.Printf("len : %d\n", len(r))
	fmt.Printf("%v \n", r)
	array()
}

func array() {
	a := [3]uint8{'a', 'b', 'c'}
	fmt.Printf("length %d %v", len(a), a[0])
	a = [...]uint8{'a', 'b', 'c'}
	// this is slice,
	c := []int8{1, 2, 3}
	c = append(c, 4)
	fmt.Printf("\nlen %d and cap %d %v", len(c), cap(c), c)
}
