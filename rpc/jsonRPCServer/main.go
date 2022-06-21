package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	gjson "github.com/gorilla/rpc/json"
)

type Args struct {
	ID string
}

type Product struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type JSONServer struct{}

func (t *JSONServer) GiveProductDetail(r *http.Request, args *Args, reply *Product) error {
	var products []Product
	absPath, _ := filepath.Abs("jsonRPCServer/products.json")
	raw, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(raw, &products)
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range products {
		if product.ID == args.ID {
			*reply = product
			return nil
		}
	}
	return errors.New("your ID doesnt exist")
}

func main() {
	s := rpc.NewServer()

	// データのエンコードとデコードを双方向で行う
	s.RegisterCodec(gjson.NewCodec(), "application/json")
	s.RegisterService(new(JSONServer), "")

	r := mux.NewRouter()
	r.Handle("/rpc", s)
	fmt.Println("JSON RPC Server Start...")
	http.ListenAndServe(":1234", r)
}
