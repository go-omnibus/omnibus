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
	temp := map[string]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			ret = append(ret, arr[i])
		}
	}
	return ret
}

func Int64ArrayUnique(arr []int64) []int64 {
	size := len(arr)
	ret := make([]int64, 0, size)
	temp := map[int64]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			ret = append(ret, arr[i])
		}
	}
	return ret
}

func IntArrayUnique(arr []int) []int {
	size := len(arr)
	ret := make([]int, 0, size)
	temp := map[int]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			ret = append(ret, arr[i])
		}
	}
	return ret
}
