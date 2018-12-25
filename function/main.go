package main

import (
	"context"

	proto "go-micro-demo/proto"
	"github.com/micro/go-micro"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

// just run one time
func main() {
	// create a new function
	fnc := micro.NewFunction(
		micro.Name("greeter"),
	)

	// init the command line
	fnc.Init()

	// register a handler
	fnc.Handle(new(Greeter))

	// run the function
	fnc.Run()
}