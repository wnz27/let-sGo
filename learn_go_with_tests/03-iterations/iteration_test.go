/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/20 6:10 下午
* Description:
 */
package iteration

import "testing"


func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expeted := "aaaaa"

	if repeated != expeted {
		t.Errorf("expected %q but got %q", expeted, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}





