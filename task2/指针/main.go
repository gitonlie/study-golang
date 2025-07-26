package main

import "fmt"

// 指针1
func pointerPerfect1(p *int) {
	*p = *p + 10
}

// 指针2
func pointerPerfect2(a []int) {
	for i := range a {
		v := a[i] * 2
		pointer := &a[i]
		*pointer = v
	}
}

func main() {
	fmt.Println("*****************指针***************")
	//指针1
	a := 10
	pointerPerfect1(&a)
	fmt.Println("a修改后的值:", a)

	//指针2
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pointerPerfect2(array)
	fmt.Println("array:", array)

}
