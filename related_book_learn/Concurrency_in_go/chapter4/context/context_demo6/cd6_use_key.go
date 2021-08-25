/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/26 01:38 8月
 **/
package main

import (
	"context"
	"fmt"
)

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func UserID(c context.Context) string {
	return c.Value(ctxUserID).(string)
}

func AuthToken(c context.Context) string {
	return c.Value(ctxAuthToken).(string)
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"handle response for userID: %v (authToken: %v)",
		UserID(ctx),
		AuthToken(ctx),
	)
}

func main() {
	ProcessRequest("jane", "abc12")
}

