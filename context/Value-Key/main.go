package main

import (
	"context"
	"go-samples/context/Value-Key/fuga"
	"go-samples/context/Value-Key/hoge"
)

func main() {
	ctx := context.Background()

	ctx = hoge.SetValue(ctx)
	ctx = fuga.SetValue(ctx)

	hoge.GetValue(ctx)
	fuga.GetValue(ctx)
}
