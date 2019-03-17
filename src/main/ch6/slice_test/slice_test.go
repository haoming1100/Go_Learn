package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	// 切片 slice 和数组的初始化很相似,不同 就是 [] 这是可变长的
	var s0 []int
	// cap 内部容量
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	// make 函数参数: 第一个参数是类型, 第二个参数是长度, 第三个参数是 cap 容量
	//cap 代表容量, len 代表个数
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		// 只所以 s = append(s, i) ,而不是 append(s, i),是因为 切片内部空间在 扩容的时候发生了变化
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
	/*
		    slice_test.go:26: 1 1
	    	slice_test.go:26: 2 2
	    	slice_test.go:26: 3 4
	    	slice_test.go:26: 4 4
	    	slice_test.go:26: 5 8
	    	slice_test.go:26: 6 8
	    	slice_test.go:26: 7 8
	    	slice_test.go:26: 8 8
	    	slice_test.go:26: 9 16
	    	slice_test.go:26: 10 16
		根据 输出结果 可以看到, cap 容量 开始是1,然后再填充的时候不够, 就 当前容量 乘以 2
	*/
}

// 切片共享内存
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	//CAP 是9 是因为 从第3位开始算起,然后加上之后所有的空间
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	// 改一个 共享的全部改变
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}
