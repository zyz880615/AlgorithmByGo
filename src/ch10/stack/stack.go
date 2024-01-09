package stack

import "errors"

/* 数据结构 - 栈
 * @Description 后进先出LIFO
 * 包含入栈、出栈以及判断栈是否为空等操作
 * @Author: zyz
 * @Date: 2024/1/9 下午 07:17
 */

type stack struct {
	Top      int
	elements []int
}

/**
 * 判断栈是否为空
 */
func stackEmpty(array stack) bool {
	//伪代码如下
	// STACK-EMPTY(S)
	// 1 if S.top == 0
	// 2    return TRUE
	// 3 else return FALSE
	if array.Top == 0 {
		return true
	}
	return false
}

/**
 * 入栈操作
 */
func push(array stack, x int) {
	// 伪代码如下
	// PUSH(S, x)
	// 1 S.top = S.top + 1
	// 2 S[S.top] = x
	array.Top++
	array.elements[array.Top] = x
}

/**
 * 出栈操作
 */
func pop(array stack) (int, error) {
	// 伪代码如下
	// POP(S)
	// 1 if STACK-EMPTY(S)
	// 2	error "underflow"
	// 3 else S.top = S.top - 1
	// 4	return S[S.top + 1]
	if stackEmpty(array) {
		return 0, errors.New("underflow")
	} else {
		array.Top--
	}
	return array.elements[array.Top+1], nil
}
