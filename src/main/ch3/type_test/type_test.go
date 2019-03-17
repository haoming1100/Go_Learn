package type_test

import "testing"

// 自定义类型
type MyInt int64

/**
Go 不支持隐式类型转换
*/
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	// 自定义类型 使用
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	// 获取 a 的指针
	aPtr := &a
	//指针的运算,在 go 中不支持指针的运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	//输出变量的类型
	t.Logf(" %T %T", a, aPtr)
}

func TestString(t *testing.T) {
	// string 类型 是一个空字符串  而不是 nil
	var s string
	t.Log("****|" + s + "|******")
	t.Log(len(s))
}
