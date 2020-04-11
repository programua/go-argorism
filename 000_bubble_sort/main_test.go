package main

import (
	"fmt"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {
	// 8要素の配列（数字）を渡す
	testArr := []int{5, 1, 0, 34, 23, 4, 9, 16}
	BubbleSort(testArr)

	// lengthが0のときはFatalエラー
	if len(testArr) < 1 {
		t.Fatalf("array length is 0")
	}

	// check max value
	if max(testArr) != testArr[len(testArr)-1] {
		t.Errorf("index %d is not max value", len(testArr)-1)
	}

	// check min value
	if min(testArr) != testArr[0] {
		t.Errorf("index 0 is not min value")
	}

	fmt.Print(testArr)
}

func max(s []int) int {
	max := s[0]
	for _, value := range s {
		if value > max {
			max = value
		}
	}

	return max
}

func min(s []int) int {
	min := s[0]
	for _, value := range s {
		if value < min {
			min = value
		}
	}

	return min
}