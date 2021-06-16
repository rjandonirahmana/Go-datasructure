package main

import "fmt"

type node struct {
	data     int
	nextData *node //nilai sambung data
}

type linkedList struct {
	head   *node
	length int //inidikasi sepanjang apa linked list nya
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.nextData = second
	l.length++
}

//untuk melihat data dibuat method value receiver
func (l linkedList) printlist() {
	diPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", diPrint.data)
		diPrint = diPrint.nextData
		l.length--

	}

	fmt.Printf("\n")
}

func (l *linkedList) deleteValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.nextData
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.nextData.data != value {
		if previousToDelete.nextData.nextData == nil {
			return
		}
		previousToDelete = previousToDelete.nextData
	}
	previousToDelete.nextData = previousToDelete.nextData.nextData
	l.length--

}

func main() {
	link := linkedList{}
	node1 := &node{data: 40}
	node2 := &node{data: 35}
	node3 := &node{data: 22}
	node4 := &node{data: 90}
	node5 := &node{data: 88}
	node6 := &node{data: 100}
	node7 := &node{data: 300}

	link.prepend(node1)

	link.prepend(node2)
	link.prepend(node3)
	link.prepend(node4)
	link.prepend(node5)
	link.prepend(node6)
	link.prepend(node7)

	link.printlist()
	link.deleteValue(100)

	link.printlist()

}
