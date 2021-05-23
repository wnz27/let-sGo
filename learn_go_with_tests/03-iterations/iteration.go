/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/20 6:09 下午
* Description:
 */
package iteration
const repeatedCount = 5


func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatedCount; i ++ {
		repeated += character
	}
	return repeated
}
