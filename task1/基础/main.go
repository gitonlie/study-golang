package main

import "fmt"

// 两数之和
func twoSum(nums []int, target int) []int {
	vMap := map[int]int{}

	var res []int
	for i := range nums {
		t := target - nums[i]
		v, exist := vMap[t]
		if exist {
			res = []int{v, i}
		} else {
			vMap[nums[i]] = i
		}
	}

	return res
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
