/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/19 23:54 8月
 **/
package con_demo

import (
	"fmt"
	"fzkprac/related_book_learn/Concurrency_in_go/chapter4/fan_out_fan_in/base_func"
	"math/rand"
	"time"
)

func Con_slow_prime_demo() {
	rand := func() interface{} { return rand.Intn(50000000)}

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := base_func.ToInt(done, base_func.RepeatFn(done, rand))
	fmt.Println("Primes:")
	for prime := range base_func.Take(done, base_func.PrimeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}
