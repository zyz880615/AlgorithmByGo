package main

import "github.com/zyz880615/AlgorithmByGo/src/ch10/LinkList"

/*
 * @Description
 * @Author: zyz
 * @Date: 2024/1/9 下午 09:26
 */

func main() {
	linkList := LinkList.InitLinkList()
	linkList.Insert(LinkList.NewListNode(1))
	linkList.Insert(LinkList.NewListNode(4))
	linkList.Insert(LinkList.NewListNode(16))
	linkList.Insert(LinkList.NewListNode(9))
	linkList.Insert(LinkList.NewListNode(25))
	linkList.Display()
	linkList.Delete(linkList.Search(4))
	linkList.Display()
}
