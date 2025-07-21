package main

import "fmt"

func singleNumber(nums []int) int {
	tmp := map[int]int{}
	for i, e := range nums {
		fmt.Printf("index:%d,nums:%d\n", i, e)
		v, exists := tmp[e]
		if exists {
			tmp[e] = v + 1
		} else {
			tmp[e] = 1
		}
	}

	var once int
	for k, v := range tmp {
		fmt.Printf("k:%d,v:%d\n", k, v)
		if v == 1 {
			once = k
			break
		}
	}

	return once
}

func main() {
	fmt.Println("测试一次")
	nums := []int{4, 1, 2, 1, 2}

	singleNumber(nums)
}
