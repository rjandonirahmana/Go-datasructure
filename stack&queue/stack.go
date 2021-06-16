package main

import "fmt"

type stack struct {
	value []int
}

//method digunakan untuk menambah nilai di akhir/ditaro dinilai paling atas
func (s *stack) push(nilai int) {
	s.value = append(s.value, nilai)
}

//nilai method ini adalah nilai aray yg dihilangkan/nilai array trakhir
func (s *stack) pop() int {
	length := len(s.value) - 1
	remove := s.value[length]
	s.value = s.value[:length]
	return remove
}

func main() {
	stack1 := stack{}
	stack1.push(11)
	stack1.push(220)
	stack1.push(230)

	fmt.Println(stack1)
	remove := stack1.pop()

	fmt.Print(stack1, "\n")
	mystack := stack1.pop()

	fmt.Println(mystack)

}
