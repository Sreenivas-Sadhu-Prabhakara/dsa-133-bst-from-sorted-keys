# Build Balanced BST from Sorted Keys

> Striver SDE Sheet 133 · Day 20 BST — root every range at its midpoint; the recursion halves the keys, so the height is forced logarithmic. Go, stdlib only, tested.

## Problem

Given keys (sorted or not), build a *height-balanced* binary search tree containing them. Inserting sorted keys naively produces a linked-list-shaped BST — the degenerate case this construction exists to avoid.

## Approach

On a sorted array, the BST property and balance can both be had by construction: take the midpoint as the root — equal halves go left and right, so subtree sizes halve at every level and height is ⌈log₂(n+1)⌉, while the inorder sequence is the sorted array by definition. Half-open ranges `[lo, hi)` keep the indexing clean. The tests check the two claims separately: inorder equals the sorted keys (it *is* a BST with the right contents) and every node passes the balance check — asserting the properties, not one blessed shape, since several equally valid trees exist for even-length ranges.

**Complexity:** O(n log n) for the sort (O(n) if pre-sorted), O(log n) recursion.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
