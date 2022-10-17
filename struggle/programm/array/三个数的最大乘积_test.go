package array

import "testing"

/**
*   @author wangchenyang
*   @date 2022/8/30 21:11
*   @description:
 */

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "测试A", args: args{nums: []int{1, 2, 3}}, want: 6},
		{name: "测试A", args: args{nums: []int{-9, -8, 3, -20}}, want: 540},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
