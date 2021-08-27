/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/28 00:27 8月
 **/
package main


// PostReport 伪码
func PostReport(id string) error {
	result, err := lowleverl.DoWork()
	if err != nil {
		// 在这里，我们检查一下接收到的异常信息，以确保它的结构是良好的。
		// 如果不是，我们就简单地将异常对到栈上，以显示出这个bug。
		if _, ok := err.(lowlevel.Error); ok {
			// 我们使用一个假设的函数将传入的异常和模块相关的信息封装起来，并赋予它
			// 一个新的类型。请注意，封装异常可能会隐藏一些底层细节，这些细节对于用户
			// 来说可能并不重要。
			err = WrapErr(err, "cannot post report with id %q", id)
		}
		return err
	}
}

func main() {

}
