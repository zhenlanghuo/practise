package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"practise/test/test_grpc/pb"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	defaultName = "world"
)

func TestClient(t *testing.T) {
	runtime.GOMAXPROCS(6)

	addr := "localhost:50052"

	// 连接grpc服务器
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()

	// 初始化Greeter服务客户端
	c := pb.NewGreeterClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 延迟关闭请求会话
	defer cancel()

	// 调用SayHello接口，发送一条消息
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// 打印服务的返回的消息
	fmt.Printf("Greeting: %s\n", r.Message)
	//fmt.Println(1, rsp)

	wg := sync.WaitGroup{}
	quit := make(chan struct{})
	count := int64(1)

	for i := 0; i < 6; i++ {
		// 连接grpc服务器
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		// 延迟关闭连接
		defer conn.Close()

		// 初始化Greeter服务客户端
		c := pb.NewGreeterClient(conn)

		for j := 0; j < 15; j++ {
			wg.Add(1)
			go func(client pb.GreeterClient) {
				defer wg.Done()
				for {
					select {
					case <-quit:
						return
					default:
						// 初始化上下文，设置请求超时时间为1秒
						ctx, cancel := context.WithTimeout(context.Background(), time.Second)
						// 延迟关闭请求会话
						defer cancel()

						// 调用SayHello接口，发送一条消息
						_, err := client.SayHello(ctx, &pb.HelloRequest{Name: "world"})
						if err != nil {
							log.Fatalf("could not greet: %v", err)
						}
						atomic.AddInt64(&count, 1)
					}
				}
			}(c)
		}
	}

	second := time.Duration(30)
	time.Sleep(time.Second * second)
	close(quit)
	wg.Wait()
	fmt.Println(count, float64(count)/float64(second))
}
