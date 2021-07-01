/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/1 10:48 下午
* Description:
 */
package t_base

func GetValue() int {
	return 1
}
func main() {
	/*
	i := GetValue()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
	 */
	/*
	编译失败。考点:类型选择，类型选择的语法形如:i.(type)，其中 i 是接口，type 是固定关键字，
	需要注意的是，只有接口类型才可以使用类型选择。
	 */

}
