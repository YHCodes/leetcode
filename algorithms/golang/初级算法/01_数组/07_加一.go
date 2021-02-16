/**/

package main

import "fmt"
//import "strconv"

func plusOne(digits []int) []int {
	n := len(digits)
	carry := 1
	for i := n - 1; i >= 0; i-- {
		val := digits[i]
		digits[i] = (val + carry) % 10
		carry = (val + carry) / 10
	}

	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	digits := []int{1,2,9}
	fmt.Println(plusOne(digits))

	digits2 := []int{1,2,3}
	fmt.Println(plusOne(digits2))

	digits3 := []int{9,9}
	fmt.Println(plusOne(digits3))
}