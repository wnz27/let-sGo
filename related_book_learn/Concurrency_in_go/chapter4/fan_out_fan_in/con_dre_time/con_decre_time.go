/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/21 00:55 8月
 **/
package con_dre_time

import (
	"fmt"
	"math/rand"
	"fzkprac/related_book_learn/Concurrency_in_go/chapter4/fan_out_fan_in/base_func"
	"runtime"
	"time"
)

func Con_d_time_dome() {
	done := make(chan interface{})
	defer close(done)
	start := time.Now()

	rand := func() interface{} {
		return rand.Intn(50000000)
	}

	randIntStream := base_func.ToInt(done, base_func.RepeatFn(done, rand))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)

	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = base_func.PrimeFinder(done, randIntStream)
	}

	for prime := range base_func.Take(done, base_func.FanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))

}
