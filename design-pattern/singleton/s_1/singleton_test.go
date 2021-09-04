/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/18 23:33 5月
 **/
package s_1_test

import (
	"fzkprac/design-pattern/singleton/s_1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, s_1.GetInstance(), s_1.GetInstance())
}

func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if s_1.GetInstance() != s_1.GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

/*
BenchmarkGetInstance-8   	1000000000	         0.3988 ns/op
 */

