package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	return nil
}

func main() {
	// TimeSeverをインスタンス化(newだとポインタが返る)
	timeServer := new(TimeServer)

	// TimeServerを登録
	rpc.Register(timeServer)

	// DefaultServerに登録
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":4444")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	fmt.Println("RPC Server Start...")
	http.Serve(l, nil)
}
