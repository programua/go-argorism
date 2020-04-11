package main

import (
	"fmt"
)

func main() {
	param := []int{5, 10, 4, 0, 2, 13, 17, 8, 20}
	BubbleSort(param)

	fmt.Println(param)
}

/**
 ele[i] > ele[i+1] が常にfalseになるまでひたすら繰り返すループ
 */
func BubbleSort(ele []int) []int {
	flag := true

	for flag {
		flag = false

		for i := 0; i < len(ele)-1; i++ {
			if ele[i] > ele[i+1] {
				flag = true
				ele[i], ele[i+1] = ele[i+1], ele[i]
			}
		}
	}

	return ele
}
