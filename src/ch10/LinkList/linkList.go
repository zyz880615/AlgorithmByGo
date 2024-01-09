package LinkList

import "fmt"

/* 链表
 * @Description
 * @Author: zyz
 * @Date: 2024/1/9 下午 09:03
 */

type LinkList struct {
	head *LinkList
	next *LinkList
	prev *LinkList
	key  int
}

func InitLinkList() (list *LinkList) {
	return &LinkList{
		head: nil,
		next: nil,
		prev: nil,
		key:  0,
	}
}

// 创建一个节点
func NewListNode(x int) (listNode *LinkList) {

	listNode = &LinkList{
		head: nil,
		next: nil,
		prev: nil,
		key:  x,
	}

	return listNode
}

func (list *LinkList) Display() {
	head := list.head
	for head != nil {
		fmt.Println("key is ", head.key)
		head = head.next
	}
}

/**
 * 搜索
 */
func (list *LinkList) Search(k int) *LinkList {
	//伪代码
	// 1 x = L.head
	// 2 while x != NIL and x.key != k
	// 3	x = x.next
	// 4 return x
	x := list.head
	for x != nil && x.key != k {
		x = x.next
	}
	return x
}

func (list *LinkList) Insert(x *LinkList) {
	//伪代码
	//1 x.next = L.head
	//2 if L.head != NIL
	//3	   L.head.prev = x
	//4 L.head = x
	//5 x.prev = NIL
	x.next = list.head
	if list.head != nil {
		list.head.prev = x
	}
	list.head = x
	x.prev = nil
}

func (list *LinkList) Delete(x *LinkList) {
	//伪代码
	//1 if x.prev != NIL
	//2		x.prev.next = x.next
	//3 else L.head = x.next
	//4 if x.next != NIL
	//5		x.next.prev = x.prev
	if x.prev != nil {
		x.prev.next = x.next
	} else {
		list.head = x.next
	}
	if x.next != nil {
		x.next.prev = x.prev
	}
}
