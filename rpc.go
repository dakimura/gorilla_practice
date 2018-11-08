package main

import (
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
	log.Println("start server...")

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(HelloService), "")
	http.Handle("/rpc", s)

	http.ListenAndServe(":8888", nil)
}
