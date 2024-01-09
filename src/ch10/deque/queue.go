package deque

/* 队列
 * @Description 先进先出
 * @Author: zyz
 * @Date: 2024/1/9 下午 07:52
 */

type deque struct {
	head     int
	tail     int
	elements []int
}

/**
 * 入队
 */
func enqueue(array deque, x int) {
	//伪代码
	// 1 Q[Q.tail] = x
	// 2 if Q.tail == Q.length
	// 3 		Q.tail = 1
	// 4 else Q.tail = Q.tail + 1
	length := len(array.elements)
	array.elements[array.tail] = x
	if array.tail == length {
		array.tail = 1
	} else {
		array.tail++
	}
}

/**
 * 出队
 */
func dequeue(array deque) int {
	// 伪代码
	// 1 x = Q[Q.head]
	// 2 if Q.head == Q.length
	// 3		Q.head = 1
	// 4 else Q.head = Q.head + 1
	// 5 return x
	length := len(array.elements)
	x := array.elements[array.head]
	if array.head == length {
		array.head = 1
	} else {
		array.head++
	}
	return x
}
