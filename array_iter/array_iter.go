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
}



