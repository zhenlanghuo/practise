package main

import (
	"fmt"
	"github.com/CodisLabs/codis/pkg/proxy"
	"github.com/CodisLabs/codis/pkg/proxy/redis"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

var bc *proxy.BackendConn
var defaultConfig = proxy.NewDefaultConfig()

func main() {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("listen failed, err: %v", err)
	}

	bc = proxy.NewBackendConn("127.0.0.1:6379", 0, defaultConfig)

	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				log.Printf("accept failed, err: %v", err)
				continue
			}
			go handleConn(c)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}

func handleConn(c net.Conn) {
	redisConn := redis.NewConn(c, 128, 128)
	for {
		multi, err := redisConn.DecodeMultiBulk()
		if err != nil {
			log.Printf("DecodeMultiBulk, err: %v", err)
			return
		}
		log.Println("request: ", formatResps(multi))

		start := time.Now()
		r := &proxy.Request{}
		r.Multi = multi
		r.Batch = &sync.WaitGroup{}
		r.Database = 0
		r.UnixNano = start.UnixNano()
		bc.PushBack(r)

		r.Batch.Wait()

		p := redisConn.FlushEncoder()
		p.MaxInterval = time.Millisecond
		p.MaxBuffered = defaultConfig.SessionMaxPipeline / 2
		err = p.Encode(r.Resp)
		if err != nil {
			log.Printf("Encode, err: %v", err)
			return
		}
		err = p.Flush(true)
		if err != nil {
			log.Printf("Flush, err: %v", err)
			return
		}
	}
}

func formatResps(resps []*redis.Resp) string {
	if len(resps) == 0 {
		return "[]"
	}
	results := make([]string, 0, len(resps))
	for _, resp := range resps {
		results = append(results, fmt.Sprintf("{Type: %v, Value: %v, Array: [%v]}", resp.Type.String(), string(resp.Value), formatResps(resp.Array)))
	}
	return fmt.Sprintf("[%s]", strings.Join(results, ","))
}
