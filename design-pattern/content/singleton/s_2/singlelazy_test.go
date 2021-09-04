/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/9/4 11:06 9月
 **/
package s_2_test

import (
	"fzkprac/design-pattern/content/singleton/s_2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, t, s_2.GetLazyInstance(), s_2.GetLazyInstance())
}

func BenchmarkGetLazyInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if s_2.GetLazyInstance() != s_2.GetLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

/*
BenchmarkGetLazyInstance-8   	1000000000	         1.104 ns/op
 */
