package main

import "fmt"

// Node belongs to a multilevel list: Next links column heads left-to-right,
// Bottom links each column's sorted chain downwards.
type Node struct {
	Val    int
	Next   *Node
	Bottom *Node
}

// Flatten merges every sorted Bottom chain into one single sorted list
// linked via Bottom, folding columns right-to-left.
func Flatten(root *Node) *Node {
	if root == nil || root.Next == nil {
		return root
	}
	rest := Flatten(root.Next)
	root.Next = nil
	return mergeBottom(root, rest)
}

func mergeBottom(a, b *Node) *Node {
	dummy := &Node{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Bottom = a
			a = a.Bottom
		} else {
			tail.Bottom = b
			b = b.Bottom
		}
		tail = tail.Bottom
	}
	if a != nil {
		tail.Bottom = a
	} else {
		tail.Bottom = b
	}
	return dummy.Bottom
}

func column(vals ...int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Bottom = &Node{Val: v}
		cur = cur.Bottom
	}
	return dummy.Bottom
}

func flatToSlice(head *Node) []int {
	out := []int{}
	for n := head; n != nil; n = n.Bottom {
		out = append(out, n.Val)
	}
	return out
}

func main() {
	a := column(5, 7, 8, 30)
	b := column(10, 20)
	c := column(19, 22, 50)
	a.Next, b.Next = b, c
	fmt.Println(flatToSlice(Flatten(a)))
}
