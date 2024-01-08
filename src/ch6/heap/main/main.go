package main

import (
	"fmt"
	"github.com/zyz880615/AlgorithmByGo/src/ch6/heap"
)

/*
 * @Description
 * @Author: zyz
 * @Date: 2024/1/8 下午 09:16
 */
func main() {
	array := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heap.HeapSort(array)
	fmt.Println(array)
}
