package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	// 在 go 中, 2 个 数组长度相同可进行比较,但是相同位置的元素也必须相同才返回 true
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	t.Log(Readable, Writable, Executable)
	a := 7 // 0111
	// go 中的 按位清零 运算符 &^
	// 对于 左右两边的操作数来说 右边 操作数 是0,左边操作数进行保留,如果右边操作数是1,则 清零为 0
	// 1 &^ 0 -- 1
	// 1 &^ 1 -- 0
	// 0 &^ 1 -- 0
	// 0 &^ 0 -- 0
	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
