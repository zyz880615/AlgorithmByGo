package main

import "testing"

/**
 * @Description
 * @Author zyz
 * @Date 2024/1/22 09:17
 **/

func TestCutRod(t *testing.T) {
	array := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	cutRod(array, 10)
}

func BenchmarkCunRod(b *testing.B) {
	array := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	for i := 0; i < b.N; i++ {
		cutRod(array, 10)
	}
}

func BenchmarkBottomUpCutRod(b *testing.B) {
	array := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	for i := 0; i < b.N; i++ {
		bottomUpCutRod(array, 10)
	}
}
