package fib

import (
	"fmt"
	"testing"
)

// 单元测试
func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 1
	//var (
	//	a int =1
	//	b int = 1
	//)
	a := 1
	b := 1
	//fmt.Print(a)
	t.Log(a)
	for i := 0; i < 5; i++ {
		//fmt.Print("", a)
		t.Log(a)
		tmp := a
		a = b
		b = tmp + a

	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2

	//temp := a
	//a = b
	//b = temp
	// golang 的一句话赋值
	a, b = b, a
	t.Log("a=", a, "b=", b)
}
