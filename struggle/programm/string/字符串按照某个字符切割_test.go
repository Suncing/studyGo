package string

import (
	"reflect"
	"testing"
)

/**
*   @author wangchenyang
*   @date 2022/8/29 17:25
*   @description:
 */

func TestSplit(t *testing.T) {
	type args struct {
		s   string
		str string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		// TODO: Add test cases.
		{name: "测试abc", args: args{s: "abc", str: "b"}, wantResult: []string{"a", "c"}},
		{name: "测试我是好人", args: args{s: "我是好人", str: "好"}, wantResult: []string{"我是", "人"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Split(tt.args.s, tt.args.str); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Split() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
