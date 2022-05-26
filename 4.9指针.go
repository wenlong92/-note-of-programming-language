package main

import (
	"fmt"
)

func main() {
	// 不同于Java和.NET，Go提供了控制数据结构的指针能力，但不能进行指针运算

	// 程序再内存中存储值，每个内存块有一个地址，通常用十六进制表示，如 0x6b0820 或 0xf84001d7f0
	// Go语言取地址符是&，放到一个变量前使用就会返回相应变量的内存地址
	var i1 = 5
	fmt.Printf("%d\n%p\n",i1,&i1) //&i1的值随着每次运行程序而变化

	// 内存地址可以存储在一个叫做指针的特殊数据类型中
	// 上面例子中是一个指向int的指针，即 i1  此处使用*int表示
	// 如果想调用指针intP，可以这样声明：
	var intP *int
	intP = &i1
	fmt.Printf("%p\n",intP)

	// 可以在指针类型前面加上 * 号来获取指针所指向的内容，* 号是一个类型更改器
	// 使用一个指针引用一个值被称为间接引用
	fmt.Printf("%d\n",*intP)
	
	// 当一个指针被定义后没有分配到任何变量时，它的值为 nil

	// 一个指针变量一般缩写为 ptr

	// 不能获取字面量或者常量的地址，如
	// const i = 5
	// var ptr, ptr2 *int
	// ptr = &i  //报错
	// ptr2 = &10  //报错

	// 指针的一个高级应用是可以传递一个变量的引用，不会传递变量的拷贝
	// 这样的传递值占用4个或8个字节，会减少内存占用、提高效率

	// 另一方面(虽然不太可能)，由于一个指针导致的间接引用(一个进程执行了另一个地址),指针的过度频繁使用也会导致性能下降

	// 指针也可以指向另一个指针，并且可以任意深度的嵌套，但会导致代码结构不清晰

	// 对于空指针的反向引用(间接引用)是不合法的
	var p *int = nil
	*p = 0 //报错
}