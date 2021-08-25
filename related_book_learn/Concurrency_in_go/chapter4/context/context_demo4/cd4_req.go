/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/26 00:38 8月
 **/
package main

import (
	"context"
	"fmt"
)

func main() {
	ProcessRequest("jane", "abc123")
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), "userID", userID)
	ctx = context.WithValue(ctx, "authToken", authToken)
	//fmt.Printf("-->: %v\n", ctx)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"handling response for userID: %v (authToken: %v)",
		ctx.Value("userID"),
		ctx.Value("authToken"),
	)
}
