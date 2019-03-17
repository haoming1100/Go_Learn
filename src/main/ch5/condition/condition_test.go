package condition

import "testing"

func TestIfMultiSec(t *testing.T) {
	t.Log("check")
	/* go 支持 多返回值*/
	if a := 1 == 1; a {
		t.Log("1 == 1")
	}
}

func TestSwitchMultiSec(t *testing.T) {
	for i := 0; i < 5; i++ {
		// switch 支持多参数匹配
		// go 语言 不用加 break
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("It's not 0 ~ 3")
		}
	}
}

// switch 充当 if
func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		// switch 支持多参数匹配
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("Don't know")
		}
	}
}
