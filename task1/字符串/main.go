package main

import (
	"fmt"
)

// 20. 有效的括号
func isValid(s string) bool {
	var queue []string
	bracket := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		str := string(runes[i])
		//切片长度=0
		if len(queue) == 0 {
			v, exists := bracket[str]
			if exists {
				queue = append(queue, v)
			} else {
				return false
			}
		} else {
			if str == queue[0] {
				queue = queue[1:]
			} else {
				v, exists := bracket[str]
				if exists {
					//头部添加
					queue = append([]string{v}, queue[:]...)
				} else {
					return false
				}
			}
		}

	}

	if len(queue) == 0 {
		return true
	}
	return false
}

// 14. 最长公共前缀
func longestPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}

	//给予默认值
	flag := true
	var prefix []rune
	for i := 0; i < len(str); i++ {
		s := str[i]

		if flag {
			prefix = []rune(s)
			flag = false
		} else {
			array := []rune(s)
			//取最小值
			size := len(prefix)
			if size > len(array) {
				size = len(array)
				prefix = prefix[:size]
			}

			for j := 0; j < size; j++ {
				t1 := array[j]
				t2 := prefix[j]

				if t1 != t2 {
					prefix = prefix[:j]
					break
				}
			}

		}

	}
	return string(prefix)
}

func main() {
	fmt.Println("有效的括号:", isValid("{}{)"))

	//str := []string{"flower"}
	var str = []string{"reflower", "flow", "flight"}
	fmt.Println(longestPrefix(str))
}
