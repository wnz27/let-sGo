/**
 * @project let-sGo
 * @Author 27
 * @Description 数组
 * @Date 2021/3/15 23:46 三月
 **/
package main

import "fmt"

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
}

func printArray2(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray3(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray4(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// 数组遍历
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 数组是值类型
	printArray(arr1)
	printArray(arr3)
	//     printArray(arr2)   cannot use arr2 (type [3]int) as type [5]int in argument to printArray

	fmt.Println(arr1, arr3)

	fmt.Println("arr1 with func2")
	printArray2(arr1)


	fmt.Println("arr3 with func2")
	printArray2(arr3)

	fmt.Println(arr1, arr3)

	fmt.Println("arr1 with func3")
	printArray3(&arr1)


	fmt.Println("arr3 with func3")
	printArray3(&arr3)

	fmt.Println(arr1, arr3)

	fmt.Println(&arr1[0], &arr1[1], &arr1[2], &arr1[3])


	fmt.Println("arr1 with func4")
	printArray4(arr1[:])


	fmt.Println("arr3 with func4")
	printArray4(arr3[:])

	fmt.Println(arr1, arr3)
}



