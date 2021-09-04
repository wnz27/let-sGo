/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/9/4 11:06 9月
 **/
package s_2

import (
	"fzkprac/design-pattern/content/singleton/s_1"
	"sync"
)

var (
	lazySingleton *s_1.Singlton
	once = &sync.Once{}
)

func GetLazyInstance() *s_1.Singlton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &s_1.Singlton{}
		})
	}
	return lazySingleton
}

