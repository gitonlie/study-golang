package main

import (
	"fmt"
	"sort"
)

// 26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 0; j < len(nums); j++ {
		if j > 0 {
			if nums[i] != nums[j] {
				i++
				nums[i] = nums[j]
			}
		}
	}

	fmt.Println(nums[:i+1])
	return i + 1
}

// 56. 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	//利用GO内置排序 按照起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)

	for i := 0; i < len(intervals); i++ {
		if i == 0 {
			merged = append(merged, intervals[i])
		} else {
			//当前
			cur := intervals[i]
			//切片中最后一个
			last := merged[len(merged)-1]

			if cur[0] <= last[1] {
				if cur[1] > last[1] {
					last[1] = cur[1]
				}
			} else {
				merged = append(merged, cur)
			}
		}
	}
	return merged
}

func main() {
	array := []int{0, 0, 1, 1, 2, 2, 2, 5, 5, 8}
	fmt.Println(removeDuplicates(array))

	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
