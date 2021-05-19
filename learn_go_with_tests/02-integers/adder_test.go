/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/19 11:19 下午
* Description:
 */
package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

