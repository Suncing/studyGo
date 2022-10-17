package main

import (
	"fmt"
)

/**
*   @author wangchenyang
*   @date 2022/8/29 16:14
*   @description:
 */
func bubbleSort() {
	temp := 0
	arr := []int{98, 9, 47, 5, 1, 8, 44, 1, 5, 6}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}
