package main

import "testing"

/**
*   @author wangchenyang
*   @date 2022/8/29 16:51
*   @description:
 */

func Test_bubbleSort(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort()
		})
	}
}
