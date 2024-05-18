package code

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{s: "cbaebabacd", p: "abc"}, wantAns: []int{0, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findAnagrams() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_bestSeqAtIndex(t *testing.T) {
	type args struct {
		height []int
		weight []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{height: []int{65, 70, 56, 75, 60, 68}, weight: []int{100, 150, 90, 190, 95, 110}}, wantAns: 6},
		{name: "test2", args: args{height: []int{8378, 8535, 8998, 3766, 648, 6184, 5506, 5648, 3907, 6773}, weight: []int{9644, 849, 3232, 3259, 5229, 314, 5593, 9600, 6695, 4340}}, wantAns: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := bestSeqAtIndex(tt.args.height, tt.args.weight); gotAns != tt.wantAns {
				t.Errorf("bestSeqAtIndex() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{s: "abc", t: "acvbsc"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
