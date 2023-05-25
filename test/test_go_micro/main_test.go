package main

import (
	"context"
	"fmt"
	"github.com/go-micro/plugins/v4/transport/grpc"
	"go-micro.dev/v4"
	"log"
	"practise/test/test_go_micro/pb"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	runtime.GOMAXPROCS(4)

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Transport(grpc.NewTransport()),
	)

	// create the greeter client using the service name and client
	greeter := pb.NewGreeterService("greeter", service.Client())

	// request the Hello method on the Greeter handler
	rsp, err := greeter.Hello(context.Background(), &pb.HelloRequest{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)

	wg := sync.WaitGroup{}
	quit := make(chan struct{})
	count := int64(1)

	for i := 0; i < 5; i++ {
		// create the greeter client using the service name and client
		greeter := pb.NewGreeterService("greeter", service.Client())
		for j := 0; j < 30; j++ {
			wg.Add(1)
			go func(client pb.GreeterService) {
				defer wg.Done()
				for {
					select {
					case <-quit:
						return
					default:
						// request the Hello method on the Greeter handler
						_, err := client.Hello(context.Background(), &pb.HelloRequest{
							Name: "John",
						})
						if err != nil {
							log.Fatalf("could not greet: %v", err)
						}
						atomic.AddInt64(&count, 1)
					}
				}
			}(greeter)
		}
	}

	second := time.Duration(30)
	time.Sleep(time.Second * second)
	close(quit)
	wg.Wait()
	fmt.Println(count, float64(count)/float64(second))
}
