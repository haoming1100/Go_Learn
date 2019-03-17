package loop

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	// 其他语言 while (n < 5)
	for n < 5 {
		t.Log(n)
		n++
	}
}
