/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/5 10:49 上午
* Description:
 */
package s_1

// 饿汉式单例

type Singlton struct{}

var singleton *Singlton

func init() {
	singleton = &Singlton{}
}

// GetInstance 获取实例
func GetInstance() *Singlton {
	return singleton
}




