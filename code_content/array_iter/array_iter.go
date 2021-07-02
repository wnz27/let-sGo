/**
 * @project let-sGo
 * @Author 27
 * @Description 切片
 * @Date 2021/3/16 01:06 3月
 **/
package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])

	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	s2 := arr[:]
	fmt.Println("s2 =", s2)
	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("==========================")
	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	// slice 扩展
	fmt.Println("Extending Slice")
	arr1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = arr1[2:6]
	s2 = s1[3:5]
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)
	/*
		slice:
		ptr 指向开头元素 len 长度，方括号取值小于这个长度，大于等于报错越界
		cap 从ptr到结束整个长度
	可以向后扩展，不能向前扩展
	s[i] 不可以超越len(s), 向后扩展不可以超越底层数组cap(s)
	*/
	fmt.Println("arr1 =", arr1)

	fmt.Printf("s1=%v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2) = %d, cap(s2) = %d\n", s2, len(s2), cap(s2))


	fmt.Println("========================================================================")

	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	ss := make([]int, 0)
	ss = append(ss, 1, 2, 3, 4)
	fmt.Println(ss)

	fmt.Println("========================================================================")

	ss1 := []int{1, 2, 3}
	ss2 := []int{4, 5}
	//ss1 = append(ss1, ss2)
	ss1 = append(ss1, ss2...)

	fmt.Println("========================================================================")

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for index, val := range slice {
		m[index] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
		fmt.Println(k, "->", &v)
	}

	// 上面的看打印就知道 每个变量用的是 同一个地址，每次拷贝到那里

	fmt.Println("XXXXXXXXXXXX")

	// 如果每个都保存则需要:
	for index, v := range slice {
		value := v
		m[index] = &value
	}
	for k, v := range m {
		fmt.Println(k, "====>", *v)
		fmt.Println(k, "====>", &v)
	}

	/*
	a := [2]int{5, 6}
	  b := [3]int{5, 6}
	  if a == b {
	    fmt.Println("equal")
	  } else {
	    fmt.Println("not equal")
	  }
	Go 中的数组是值类型，可比较，另外一方面，数组的⻓度也是数组类型的组成部分，所以 a 和 b 是不同的类型，是不能比较的，所以编译错误。
	 */


	var s111 []int
	//var s222 = []int{}
	if s111 == nil {
		fmt.Println("yes nil")
	}else{
		fmt.Println("no nil")
	}
	/*
	只有s111 可以 和 nil做判断
	nil 切片和 nil 相等，一般用来表示一个不存在的切片; 空切片和 nil 不相等，表示一个空的集合。
	 */

}



