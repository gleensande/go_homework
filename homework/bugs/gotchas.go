package main

// сюда вам надо писать функции, которых не хватает, чтобы проходили тесты в gotchas_test.go

import (
	"strings"
	"strconv"
	"sort"
)

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	a1 := [...]int{1, 3, 4}
	return a1
}

func ReturnIntSlice() []int {
	a1 := []int{1, 2, 3}
	return a1
}

func IntSliceToString(intSlice []int) string {
	stringSlice := []string{}
	for i := 0; i < len(intSlice); i++ {
		stringSlice = append(stringSlice, strconv.Itoa(intSlice[i]))
	}
	return strings.Join(stringSlice,"")
}

func MergeSlices(floatSlice []float32, intSlice []int32) []int {
	mergedSlice := []int{}
	for i := 0; i < len(floatSlice); i++ {
		mergedSlice = append(mergedSlice, int(floatSlice[i]))
	}
	for i := 0; i < len(intSlice); i++ {
		mergedSlice = append(mergedSlice, int(intSlice[i]));
	}
	return mergedSlice;
}

func GetMapValuesSortedByKey(input map[int]string) []string {
	sortedValues := []string{}

	keys := []int{}
	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		sortedValues = append(sortedValues, input[k]);
	}

	return sortedValues
}

