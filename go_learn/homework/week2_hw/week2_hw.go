/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/4/19 20:05 4月
 **/
package week2_hw

/*
TODO 作业
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
 */

/*
分析:
错误处理有几个原则：
1、错误要被日志记录
2、应用程序处理错误，保证100% 完整性
3、之后不再报告当前错误。
 */

func T() {
	println("hello world!!")
}
