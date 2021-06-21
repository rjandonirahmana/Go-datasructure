package main

import "fmt"

const arraysize int = 9

type hashtable struct {
	array [arraysize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key    int
	linked *bucketNode
}

func (h *hashtable) insert(key int) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *hashtable) deleteValue(key int) {

}
//func (h *hashtable) search(key int) bool {
//	return true
// }

func (b *bucket) insert	 {

}

func hash(value string) int {
	sum := 0
	for _, v := range value {
		sum += int(v)
	}
	return sum % arraysize
}


func Init() *hashtable {
	list := &hashtable{}
	for i := range list.array {
		list.array[i] = &bucket{}
	}
	return list
}

func main() {
	hashtable := Init()
	fmt.Print(hashtable, "\n")

	hashint := hash("anjysoliyhsbshfwopwfulordfg")
	fmt.Println(hashint)

}
