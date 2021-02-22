package main

import "fmt"

// 标准声明(全局)
var name string
var age int
var isOk bool

// 批量声明
var (
	a string
	b int
	c bool
	d float32
)

func main() {
	// 赋值(初始化)
	name = "Kphilleani"
	age, isOk = 24, false

	// 类型推导
	var addr = "shanghai"
	fmt.Println(addr)

	// 短变量
	n := 10
	m := 200
	fmt.Println(n, m)

	// 匿名变量 _
	x, _ := foo()
	_, y := foo()
	fmt.Println(x, y)

	const pi = 3.1415
	const e = 2.7182
	// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
	const (
		n1 = 100
		n2
		n3
	)

	// iota
	const (
		x1 = iota
		x2
		_
		x3
	)
	fmt.Println(x1, x2, x3)

	const (
		y1 = iota
		y2 = 100
		y3 = iota
		y4
	)
	fmt.Println(y1, y2, y3, y4)
}

func foo() (int, string){
	return 10, "Kp"
}
