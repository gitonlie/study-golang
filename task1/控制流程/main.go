package main

import (
	"fmt"
	"strconv"
)

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func singleNumber(nums []int) int {
	tmp := map[int]int{}
	for i := range nums {
		key := nums[i]
		v, exists := tmp[key]
		if exists {
			tmp[key] = v + 1
		} else {
			tmp[key] = 1
		}
	}

	var once int
	for k, v := range tmp {
		if v == 1 {
			once = k
			break
		}
	}

	return once
}

// 9.给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	a := strconv.Itoa(x)
	runes := []rune(a)
	length := len(runes)
	for i := 0; i < length; i++ {
		t1 := runes[i]
		t2 := runes[length-1-i]
		if t1 != t2 {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
	fmt.Println(isPalindrome(112211))
}
