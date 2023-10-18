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
