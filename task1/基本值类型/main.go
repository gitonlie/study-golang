package main

import (
	"fmt"
)

// 66. 加一
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	t := false
	for i := len(digits) - 1; i >= 0; i-- {
		if t {
			if digits[i] == 9 {
				digits[i] = 0
				t = true
			} else {
				digits[i] = digits[i] + 1
				t = false
			}
		}

		if i == len(digits)-1 {
			if digits[i] == 9 {
				digits[i] = 0
				t = true
			} else {
				digits[i] = digits[i] + 1
			}
		}

	}

	if t {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	digits := []int{9}
	fmt.Println(plusOne(digits))

}
