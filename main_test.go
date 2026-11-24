package main

import (
	"reflect"
	"sort"
	"testing"
)

func height(n *TreeNode) int {
	if n == nil {
		return 0
	}
	l, r := height(n.Left), height(n.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}

func balanced(n *TreeNode) bool {
	if n == nil {
		return true
	}
	l, r := height(n.Left), height(n.Right)
	d := l - r
	return d >= -1 && d <= 1 && balanced(n.Left) && balanced(n.Right)
}

func TestBuildBalanced(t *testing.T) {
	keys := []int{7, 1, 3, 5, 9, 11, 13}
	root := BuildBalanced(keys)
	// Inorder must be the sorted keys: proves it's a BST with all keys.
	want := append([]int{}, keys...)
	sort.Ints(want)
	if got := inorder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("inorder = %v, want %v", got, want)
	}
	if !balanced(root) {
		t.Error("tree is not height-balanced")
	}

	if BuildBalanced(nil) != nil {
		t.Error("empty keys should build nil")
	}
	if got := BuildBalanced([]int{5}); got.Val != 5 || got.Left != nil || got.Right != nil {
		t.Error("single key build wrong")
	}
	// 15 keys -> height exactly 4 (perfectly fillable).
	seq := make([]int, 15)
	for i := range seq {
		seq[i] = i
	}
	if h := height(BuildBalanced(seq)); h != 4 {
		t.Errorf("height = %d, want 4", h)
	}
}
