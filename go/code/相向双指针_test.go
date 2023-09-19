package code

import (
	"reflect"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		// TODO: Add test cases
		{"test1", args{nums: []int{-1, 2, 1, -4}, target: 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := threeSumClosest(tt.args.nums, tt.args.target); gotAns != tt.wantAns {
				t.Errorf("threeSumClosest() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{nums: []int{-2, -1, -1, 1, 1, 2, 2}, target: 0}, wantAns: [][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("fourSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
