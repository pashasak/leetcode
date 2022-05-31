package leetcode

import (
	"fmt"
	"testing"
)

func Test_392(test *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s, %s, %t", tt.s, tt.t, tt.want)
		test.Run(testname, func(t *testing.T) {
			if x := isSubsequence(tt.s, tt.t); x != tt.want {
				t.Fatalf("Got %t. Want %t", x, tt.want)
			}
		})
	}
	s := "abc"
	t := "ahbgdc"
	if x := isSubsequence(s, t); x != true {
		test.Fatalf("Got %t. Want %t. %q is a subsequence of  %q.", x, !x, s, t)
	}
}

func Test_165(t *testing.T) {
	tests := []struct {
		v1, v2 string
		want   int
	}{
		{"1.02", "1.001", 1},
		{"1.01", "1.001", 0},
		{"1.0", "1.0.0", 0},
		{"0.1", "1.1", -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s, %s, %d", tt.v1, tt.v2, tt.want)
		t.Run(testname, func(t *testing.T) {
			if x := compareVersion(tt.v1, tt.v2); x != tt.want {
				t.Fatalf("Got %d. Want %d", x, tt.want)
			}

		})
	}
}
