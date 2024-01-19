package main

import "fmt"

func main() {
	array := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	r := cutRod(array, 10)
	fmt.Printf("最大收益为:%d\n", r)
}

func cutRod(array []int, n int) int {
	if n == 0 {
		return 0
	}
	q := -1
	for i := 1; i <= n; i++ {
		q = max(q, array[i]+cutRod(array, n-i))
	}
	return q
}
