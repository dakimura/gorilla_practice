package main

import (
	"fmt"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
)

type HelloArgs struct {
	Who string
}

type HelloResponse struct {
	Message string
}

type HelloService struct{}

func (h *HelloService) SayHi(r *http.Request, args *HelloArgs, response *HelloResponse) error {
	response.Message = "Hi, " + args.Who + "!"
	log.Println(r, args, response)
	return nil
}

func main() {
	fmt.Println("hello")

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application_json")
	s.RegisterService(new(HelloService), "")
	http.Handle("/rpc", s)

	http.ListenAndServe(":8888", nil)
}
