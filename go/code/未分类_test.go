package code

import (
	"testing"
)

func Test_maxOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{args: args{nums: []int{3, 1, 3, 4, 3}, k: 6}, wantCount: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := maxOperations(tt.args.nums, tt.args.k); gotCount != tt.wantCount {
				t.Errorf("maxOperations() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_maxVowels(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{args: args{s: "abciiidef", k: 3}, wantAns: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxVowels(tt.args.s, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maxVowels() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
