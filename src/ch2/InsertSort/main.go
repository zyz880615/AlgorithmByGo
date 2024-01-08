package main

import "fmt"

/** 插入排序
 * @Description Go语言编写插入排序
 * @Author zyz
 * @Date 2024/1/3 16:35
 **/

func main() {
	array := []int{5, 2, 4, 6, 1, 3}
	fmt.Println(insertSort(array))
}

/*
*
插入排序
*/
func insertSort(array []int) []int {
	//伪代码
	//1 for j = 2 to A.length
	//2	key = A[j]
	//3	i = j - 1
	//4	while i > 0 and A[i] > key
	//5		A[i + 1] = A[i]
	//6		i = i - 1
	//7	A[i + 1] = key
	for j := 1; j < len(array); j++ {
		key := array[j]
		i := j - 1
		for i >= 0 && array[i] > key {
			array[i+1] = array[i]
			i = i - 1
		}
		array[i+1] = key
	}
	return array
}
