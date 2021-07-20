/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/20 4:20 下午
* Description:
 */
package  main

import (
	"context"
	"go.opencensus.io/trace"
)


func List(ctx context.Context, db *sqlx.DB) ([]User, error) {
	ctx, span := trace.StartSpan(ctx, "internal.user.List")
	defer span.End()

	users := []User{}
	const q = `SELECT * FROM users`

	if err := db.SelectContext(ctx, &users, q); err != nil {
		return nil, errors.Wrap(err, "selecting users")
	}

	return users, nil
}
/*
像在上述例子中，我们会把所传入方法的 context 一层层的传进去下一级方法。这里就是将外部的 context 传入 List 方法，
再传入 SQL 执行的方法，解决了 SQL 执行语句的时间问题。
 */
