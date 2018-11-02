
package main 

import (
	"fmt"
)

type treeNode struct {
	value int
	left, right *treeNode
}

// 工厂函数
func createNode(value int) *treeNode{
	// 注意返回了局部变量的地址 
	return &treeNode{value: value}
}

// 定义方法  (node treeNode)接受者
func (node treeNode) print(){
	fmt.Print(node.value," ")
}

func (node treeNode) setValue(value int){
	node.value = value
}

func (node *treeNode) traverse(){
	if node == nil {
		return
	}
	// 树的遍历
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main(){

	var root treeNode
	// 创建结构体
	root = treeNode{value: 3}
	// 因为left right是指针，所以取地址
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	// new()创建结构体
	root.right.left = new(treeNode)

	fmt.Println(root)

	root.left.right = createNode(2)

	fmt.Println(root.left.right)


	root.left.setValue(10)
	// 没有改变value的值，说明是传的是值 而不是地址
	root.left.print()
	// 只有指针才可以改变结构内容
	// nil指针也可以调用方法

	fmt.Println()
	root.traverse()
	// 接受者是go语言特有
	// 值/指针接受者均可以接收值/指针（值接受者也可以调用指针方法）
}























