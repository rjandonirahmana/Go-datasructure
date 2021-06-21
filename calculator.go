package main

import (
	"fmt"
	"strconv"
)

func main() {
	cmd := readLine("wellcome to calculattor, click a:add, s:substract, m:multiply, d:divide")
	if cmd == "a" {
		num1, num2 := getUserNumber()
		result := num1 + num2
		result1 := fmt.Sprintf("%d", result)
		fmt.Printf("result is : %s\n", result1)

	} else if cmd == "s" {
		num1, num2 := getUserNumber()
		result := num1 - num2
		result1 := fmt.Sprintf("%d", result)
		fmt.Printf("result is : %s\n", result1)

	} else if cmd == "m" {
		num1, num2 := getUserNumber()
		result := num1 * num2
		result1 := fmt.Sprintf("%d\n", result)
		fmt.Printf("result is : %s", result1)

	} else if cmd == "d" {
		num1, num2 := getUserNumber()
		result := num1 / num2
		result1 := fmt.Sprintf("%d\n", result)
		fmt.Printf("result is : %s", result1)

	} else {
		fmt.Println("error")
	}

}

func readLine(message string) string {
	fmt.Println(message)
	var input string
	fmt.Scanln(&input)
	return input
}
func getUserNumber() (int, int) {
	num1 := readLine("input your first number : ")
	num1int, err := strconv.Atoi(num1)
	if err != nil {
		return 0, 0
	}

	num2 := readLine("input your second number : ")
	num2int, err := strconv.Atoi(num2)
	if err != nil {
		return 0, 0
	}
	return num1int, num2int

}
