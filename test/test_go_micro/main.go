package main

import (
	"context"
	//"github.com/micro/go-micro/transport/grpc"
	//"github.com/micro/go-plugins/transport/grpc/v2"
	"github.com/go-micro/plugins/v4/transport/grpc"
	"go-micro.dev/v4"
	"log"
	"net/http"
	_ "net/http/pprof"
	"practise/test/test_go_micro/pb"
	"runtime"
)

//	type Request struct {
//		Name string `json:"name"`
//	}
//
//	type Response struct {
//		Message string `json:"message"`
//	}
type Greeter struct{}

func (h *Greeter) Hello(ctx context.Context, req *pb.HelloRequest, rsp *pb.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	runtime.GOMAXPROCS(4)

	//service := micro.NewService()
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Transport(grpc.NewTransport()),
	)
	service.Init()

	err := pb.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
