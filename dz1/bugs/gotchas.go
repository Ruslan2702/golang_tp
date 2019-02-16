package main

import (
	"sort"
	"strconv"
)

// сюда вам надо писать функции, которых не хватает, чтобы проходили тесты в gotchas_test.go

func ReturnInt() int{
	return 1
}

func ReturnFloat() float32{
	return 1.1
}

func ReturnIntArray() [3]int{
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int{
	return []int{1, 2, 3}
}

func IntSliceToString(arr []int) string{
	var result string
	for _, val := range arr {
		result += strconv.Itoa(val)
	}

	return result
}

func MergeSlices(floarArr []float32, intArr []int32)  []int{
	result := make([]int, len(floarArr) + len(intArr), len(floarArr) + len(intArr))
	for idx, val := range floarArr {
		result[idx] = int(val)
	}

	for idx, val := range intArr {
		result[len(floarArr) + idx] = int(val)
	}

	return result
}

func GetMapValuesSortedByKey(seasons map[int]string) []string{
	keys := make([]int, 0, len(seasons))
	result := make([]string, 0, len(seasons))
	for val := range seasons {
		keys = append(keys, val)
	}

	sort.Ints(keys)
	for _, key := range keys {
		result = append(result, seasons[key])
	}

	return result
}