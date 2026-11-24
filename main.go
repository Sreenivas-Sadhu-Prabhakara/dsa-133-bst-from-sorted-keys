package main

import (
	"fmt"
	"sort"
)

// TreeNode is a binary search tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// BuildBalanced returns a height-balanced BST holding the given keys:
// sort them, then recursively root each subtree at its range's midpoint.
func BuildBalanced(keys []int) *TreeNode {
	sorted := append([]int{}, keys...)
	sort.Ints(sorted)
	var build func(lo, hi int) *TreeNode // half-open [lo, hi)
	build = func(lo, hi int) *TreeNode {
		if lo >= hi {
			return nil
		}
		mid := (lo + hi) / 2
		return &TreeNode{
			Val:   sorted[mid],
			Left:  build(lo, mid),
			Right: build(mid+1, hi),
		}
	}
	return build(0, len(sorted))
}

func inorder(n *TreeNode) []int {
	if n == nil {
		return nil
	}
	return append(append(inorder(n.Left), n.Val), inorder(n.Right)...)
}

func main() {
	root := BuildBalanced([]int{7, 1, 3, 5, 9, 11, 13})
	fmt.Println(root.Val, inorder(root))
}
