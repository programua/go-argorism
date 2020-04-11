package main

import "fmt"

type param []int

func main() {
	p := param{2,4,6,3,9}
	//p := param {5, 10, 30, 27, 4, 0, 2, 13, 17, 8, 20, 23, 26}
	result := p.QuickSort(0, len(p)-1, p)

	fmt.Println(result)
}

func (p param) QuickSort(low, high int, param param) []int {
	if low >= high {
		return param
	}

	start := low
	max := high

	// 基準にする値を先頭の値とする
	pick := param[low]

	for low < high {
		// 基準値より大きい値を前から検索
		for low <= high && param[low] <= pick {
			low++
		}
		// 基準値より小さい値を後ろから検索
		for low <= high && param[high] > pick {
			high--
		}

		// 値の入れ替え
		if low < high {
			p[low], p[high] = p[high], p[low]
		}
	}

	// 基準値を中央に移動
	p[start], p[high] = p[high], p[start]

	// 再帰的呼び出し
	p.QuickSort(start, high-1, p)
	p.QuickSort(high + 1, max, p)
	return param
}
