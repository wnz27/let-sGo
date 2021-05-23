/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/23 20:12 5月
 **/
package main

func Sum(numbers [5]int) (sum int) {
	sum = 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return
}

