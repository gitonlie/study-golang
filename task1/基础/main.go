package main

import "fmt"

// 两数之和
func twoSum(nums []int, target int) []int {
	tmpMap := make(map[int][]int, len(nums))
	for i, e := range nums {
		v, exist := tmpMap[e]
		if exist {
			tmpMap[e] = append([]int{i}, v...)
		} else {
			tmpMap[e] = []int{i}
		}
	}
	fmt.Println(tmpMap)

	var res []int
	for i := range nums {
		t := target - nums[i]
		v, exist := tmpMap[t]
		if exist {
			if t == nums[i] && len(v) == 1 {
				//表明是唯一值
				continue
			}

			if t != nums[i] {
				res = append(res, i)
			}
			res = append(res, v...)
			break
		}
	}

	return res
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
