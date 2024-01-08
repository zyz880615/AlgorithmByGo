package heap

/* 堆
 * @Description 包含堆的很多方法
 * @Author: zyz
 * @Date: 2024/1/8 下午 07:49
 */

var largest = -1

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

// MaxHeapify 维护最大堆
func MaxHeapify(array []int, len, i int) {
	//MAX-HEAPIFY(A, i)
	//1 l = LEFT(i)
	//2 r = RIGHT(i)
	//3 if l <= A.heap-size and A[l] > A[i]
	//4 	 largest = l
	//5 else largest = i
	//6 if r <= A.heap-size and A[r] > A[largest]
	//7	     largest = r
	//8 if largest != i
	//9 	 exchange A[i] with A[largest]
	//10	 MAX-HEAPIFY(A, largest)
	l := left(i)
	r := right(i)
	if l < len && array[l] > array[i] {
		largest = l
	} else {
		largest = i
	}

	if r < len && array[r] > array[largest] {
		largest = r
	}

	if largest != i {
		array[i], array[largest] = array[largest], array[i]
		MaxHeapify(array, len, largest)
	}
}

// BuildMaxHeap 建立最大堆
func BuildMaxHeap(array []int, len int) {
	for i := len / 2; i >= 0; i-- {
		MaxHeapify(array, len, i)
	}
}

// HeapSort 堆排序 //
func HeapSort(array []int) {
	n := len(array)
	BuildMaxHeap(array, n)
	for i := n - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		MaxHeapify(array, i, 0)
	}
}
