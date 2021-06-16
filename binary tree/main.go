 package main

import "fmt"

type node struct {
	key   int
	left  *node
	right *node
}

//method insert menempatkan nilai di sisi kiri jika lebih kecil dari key dan sebelah kanan jika lebih besar
func (n *node) insert(k int) {
	if n.key > k {
		//diletakan disebelah kiri
		if n.left == nil {
			n.left = &node{key: k}
		} else {
			n.left.insert(k)
		}
	} else if n.key < k {
		//diletakan disebelah kanan
		if n.right == nil {
			n.right = &node{key: k}
		} else {
			//recursive method/ pengulangan mehod dari awal
			n.right.insert(k)
		}
	}
}

func (n *node) search(k int) bool {
	if n == nil {
		return false
	}

	if n.key < k {
		return n.right.search(k)
	} else if n.key > k {
		return n.left.search(k)
	}

	return true
}

func main() {
	tree := &node{key: 200}
	fmt.Println(tree)

	tree.insert(100)
	tree.insert(201)
	tree.insert(300)
	tree.insert(99)
	tree.insert(38)

	fmt.Println(tree.search(300))
}
