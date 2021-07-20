/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/20 4:13 下午
* Description:
 */
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// 对第三方调用要传入 context
// 在 Go 语言中，Context 的默认支持已经是约定俗称的规范了。因此在我们对第三方有调用诉求的时候，要传入 context：
func main() {
	req, err := http.NewRequest("GET", "https://eddycjy.com/", nil)
	if err != nil {
		fmt.Printf("http.NewRequest err: %+v", err)
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 50*time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("http.DefaultClient.Do err: %+v", err)
		return
	}
	defer resp.Body.Close()
}
/*
这样子由于第三方开源库已经实现了根据 context 的超时控制，那么当你所传入的时间到达时，将会中断调用。
若你发现第三方开源库没支持 context，那建议赶紧跑，换一个。免得在微服务体系下出现级联故障，还没有简单的手段控制，那就很麻烦了。
 */
