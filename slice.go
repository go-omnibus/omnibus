package omnibus

import (
	"strings"
)

func Implode(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	}

	return strings.Split(text, delimiter)
}

func Int64sToStrings(ints []int64) []string {
	var ss []string
	for _, i := range ints {
		ss = append(ss, DefiniteString(i))
	}
	return ss
}

func IntsToStrings(ints []int) []string {
	var ss []string
	for _, i := range ints {
		ss = append(ss, DefiniteString(i))
	}
	return ss
}

func StringsToInts(ss []string) []int {
	var ints []int
	for _, s := range ss {
		ints = append(ints, DefiniteInt(s))
	}
	return ints
}

func StringsToInt64s(ss []string) []int64 {
	var ints []int64
	for _, s := range ss {
		ints = append(ints, DefiniteInt64(s))
	}
	return ints
}

func StringArrayUnique(arr []string) []string {
	size := len(arr)
	ret := make([]string, 0, size)
	memo := make(map[string]bool)
	for i := 0; i < size; i++ {
		if _, ok := memo[arr[i]]; ok != true {
			memo[arr[i]] = true
			ret = append(ret, arr[i])
		}
	}
	return ret
}

func Int64ArrayUnique(arr []int64) []int64 {
	size := len(arr)
	ret := make([]int64, 0, size)
	memo := make(map[int64]bool)
	for i := 0; i < size; i++ {
		if _, ok := memo[arr[i]]; ok != true {
			memo[arr[i]] = true
			ret = append(ret, arr[i])
		}
	}
	return ret
}

func IntArrayUnique(arr []int) []int {
	size := len(arr)
	ret := make([]int, 0, size)
	memo := make(map[int]bool)
	for i := 0; i < size; i++ {
		if _, ok := memo[arr[i]]; ok != true {
			memo[arr[i]] = true
			ret = append(ret, arr[i])
		}
	}
	return ret
}

func InArray(needle interface{}, hystack interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range hystack.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range hystack.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range hystack.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
	}
	return false
}

func StringArrayDiff(array1 []string, arrayOthers ...[]string) []string {
	c := make(map[string]bool)
	for i := 0; i < len(array1); i++ {
		if _, hasKey := c[array1[i]]; hasKey {
			c[array1[i]] = true
		} else {
			c[array1[i]] = false
		}
	}
	for i := 0; i < len(arrayOthers); i++ {
		for j := 0; j < len(arrayOthers[i]); j++ {
			if _, hasKey := c[arrayOthers[i][j]]; hasKey {
				c[arrayOthers[i][j]] = true
			} else {
				c[arrayOthers[i][j]] = false
			}
		}
	}
	result := make([]string, 0)
	for k, v := range c {
		if !v {
			result = append(result, k)
		}
	}
	return result
}
