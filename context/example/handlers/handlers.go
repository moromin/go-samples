package handlers

import (
	"context"
	"fmt"
	"go-samples/context/example/auth"
	"net/http"
	"time"
)

type MyHandleFunc func(context.Context, MyRequest)

type MyRequest struct {
}

type MyResponse struct {
	Code int
	Err  error
	Body string
}

var GetGreeting MyHandleFunc = func(ctx context.Context, req MyRequest) {
	var res MyResponse

	// tokenからユーザー検証→ダメなら即return
	userID, err := auth.VerifyAuthToken(ctx)
	if err != nil {
		res = MyResponse{Code: http.StatusForbidden, Err: err}
		fmt.Println(res)
		return
	}

	// DBリクエストをいつタイムアウトさせるかcontext経由で設定
	dbReqCtx, cancel := context.WithTimeout(ctx, 2*time.Second)

	// DBからデータ取得
	rcvChan := db.DefaultDB.Search(dbReqCtx, userID)
	data, ok := <-rcvChan
	cancel()

	// DBリクエストがタイムアウトしていたら408で返す
	if !ok {
		res = MyResponse{Code: http.StatusRequestTimeout, Err: err}
		fmt.Println(res)
		return
	}

	// レスポンスの作成
	res = MyResponse{
		Code: http.StatusOK,
		Body: fmt.Sprintf("From path %s, Hello! your ID is %d\ndata -> %s", req.path, userID, data),
	}

	// レスポンス内容を標準出力
	fmt.Println(res)
}
