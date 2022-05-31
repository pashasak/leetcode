package leetcode

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	f := func(v string) int {
		i, _ := strconv.Atoi(v)
		return i
	}

	v1 := Map(strings.Split(version1, "."), f)
	v2 := Map(strings.Split(version2, "."), f)

	ln := Min(len(v1), len(v2))

	for i := 0; i < ln; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}
	return 0
}

func Map[T1, T2 any](input []T1, f func(T1) T2) (output []T2) {
	output = make([]T2, 0, len(input))
	for _, v := range input {
		output = append(output, f(v))
	}
	return output
}

type Ordered interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		uintptr | float32 | float64 | string
}

func Max[T1 Ordered](input ...T1) (max T1) {
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	return max
}

func Min[T1 Ordered](input ...T1) T1 {
	min := input[0]
	for _, v := range input {
		if v < min {
			min = v
		}
	}
	return min
}
