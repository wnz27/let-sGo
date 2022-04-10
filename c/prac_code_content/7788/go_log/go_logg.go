/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/4/22 20:51 4月
 **/
package main

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func main() {
	err := errors.Errorf("fzk27")
	err2 := errors.Errorf("66666666")
	errs := []error{err, err2}
	logger1 := zap.NewExample()
	defer func() {
		fmt.Println("--first--")
	}()

	defer logger1.Sync()
	logger1.Info(
		"1",
		zap.String("1", "2"),
		zap.Int("aaa", 2),
		zap.Errors("Hi好", errs),
		)
	//fmt.Println(logger1)
	// ----------------------   ---------------------- //
	fmt.Println(" ----------------------   ", " ----------------------   ")
	pLogger, _ := zap.NewProduction()
	pLogger.Error("high_level",
		zap.Errors("err", errs),
		)


	// fatal defer不会被调用
	//pLogger.Fatal("alalalal",
	//	zap.Errors("fatal", errs),
	//	)

	// defer 会调用的
	pLogger.Panic("alalalal",
		zap.Errors("fatal", errs),
	)

	fmt.Println(" =====================   ", " =====================   ")

}
