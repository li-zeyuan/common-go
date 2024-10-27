package utils

import (
	"golang.org/x/exp/constraints"
	"math/rand"
	"time"
)

func RandShuffle(slice []interface{}) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func UniqueStr(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}

	strMap := make(map[string]struct{}, len(slice))
	newSlice := make([]string, 0, len(slice))
	for _, item := range slice {
		if _, ok := strMap[item]; ok {
			continue
		}

		strMap[item] = struct{}{}
		newSlice = append(newSlice, item)
	}

	return newSlice
}

func UniqueArray(iArray []int64) []int64 {
	uniqueMap := make(map[int64]struct{})
	result := make([]int64, 0, len(iArray))

	for _, i := range iArray {
		if _, ok := uniqueMap[i]; !ok {
			result = append(result, i)
			uniqueMap[i] = struct{}{}
		}
	}

	return result
}

func IsInArray(item int, array []int) bool {
	for _, arrayItem := range array {
		if arrayItem == item {
			return true
		}
	}

	return false
}

func IsInStrArray(item string, array []string) bool {
	for _, arrayItem := range array {
		if arrayItem == item {
			return true
		}
	}

	return false
}

func ExcludeSlice(all, exclude []string) []string {
	if len(all) == 0 || len(exclude) == 0 {
		return all
	}

	excludeMap := make(map[string]struct{}, len(exclude))
	for _, e := range exclude {
		excludeMap[e] = struct{}{}
	}

	newSlice := make([]string, 0, len(all))
	for _, item := range all {
		if _, ok := excludeMap[item]; !ok {
			newSlice = append(newSlice, item)
		}

	}

	return newSlice
}

func IsInSlice[T constraints.Ordered](a T, s []T) bool {
	if len(s) == 0 {
		return false
	}

	for _, item := range s {
		if item == a {
			return true
		}
	}

	return false
}
