package fuga

import (
	"context"
	"fmt"
	"log"
)

// 異なるパッケージ間でcontextを共有したときのkey衝突を避けるために
// keyにセットする値にstring型のようなビルトインな型を使うべきではありません。
// その代わり、ユーザーはkeyには独自型を定義して使うべきです。
// https://pkg.go.dev/context@go1.17#WithValue
type ctxKey int

const (
	a ctxKey = iota + 1
)

func SetValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, a, "c")
}

func GetValue(ctx context.Context) {
	val, ok := ctx.Value(a).(string)
	if !ok {
		log.Fatal("failed to assertion hoge.GetValue()")
	}
	fmt.Println(val)
}
