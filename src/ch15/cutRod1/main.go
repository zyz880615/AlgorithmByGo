package main

import "fmt"

func main() {
	array := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	r := cutRod(array, 10)
	fmt.Printf("最大收益为:%d\n", r)
	r1 := bottomUpCutRod(array, 10)
	fmt.Printf("最大收益为:%d\n", r1)
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

func bottomUpCutRod(p []int, n int) int {
	//伪代码
	//1 let r[0..n] be a new array
	//2 r[0] = 0
	//3 for j= 1 to n
	//4 		q= -1
	//5		for i = 1 to j
	//6			q = max(q,p[i] + r[j-i])
	//7		r[j] = q
	//8 return r[n]
	r := [11]int{0}
	for j := 1; j <= n; j++ {
		q := -1
		for i := 1; i <= j; i++ {
			q = max(q, p[i]+r[j-i])
		}
		r[j] = q
	}
	return r[n]
}
