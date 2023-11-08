package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:1234",
		Password: "",
		DB:       0, // use default DB
		//OnConnect: func(conn *redis.Conn) error {
		//	conn.Process(redis.NewStatusCmd("xxx"))
		//	return nil
		//},
	})
	//
	cmd := redisCli.Set("1", 2, time.Minute*60)
	if cmd.Err() != nil {
		fmt.Printf("Set failed, err: %v\n", cmd.Err())
		return
	}
	fmt.Println(cmd.Val())

	redisCli.Pipeline()

	//redisCli.Ping()
	cmd := redisCli.Info("replication")
	fmt.Println(cmd.Val())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	//redis.NewClusterClient()
}
