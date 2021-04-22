/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/4/19 20:05 4月
 **/
package week2_hw

import (
	"github.com/pkg/errors"
	"log"
)

type FindResult struct {
	rows int
}

// 这类我们认为不是我们期望的并且我们不处理的结果。
func FindSomethingRaise() (FindResult, error) {
	// find operation
	// 最底下的调用给个信息抛出去， 我们不处理所以也就不打日志交给调用方处理的来做

	// 这个查询操作如果有错 且是第三方的应该用wrapf 否则应该是自己被很多调用所以应该去用 Errorf
	res, err := mockFind()
	return res, errors.Wrapf(err, "find something no rows!")
	//return res, errors.Errorf("find something no rows!")
}



// 这类是我们确定返回这个错ok没问题的 我们可以处理
func FindSomethingCanHandle() (FindResult, error) {
	// 这里的norows 错误我们认为是符合预期的
	// 查询操作如果有now rows错误则忽略掉认为没有问题
	res, _ := mockFind()
	return res, nil
}

// 这类是我们需要终止程序，并且打日志的
func ForceFindWithException(param string) (FindResult, error) {
	// 如果我们认为某个find有问题调用方不该继续下去就把程序干掉
	res, e := mockFind()
	if e != nil {
		log.Panicf("table: [%s], param: [%s] not found", "table", param)
	}
	return res, e
}

func mockFind() (FindResult, error) {
	return FindResult{rows: 0}, errors.Errorf("root err")
}

