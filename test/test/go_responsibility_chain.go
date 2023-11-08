package main

import (
	"context"
	"fmt"
)

type HandlerFunc func(ctx context.Context)

type Wrapper func(next HandlerFunc) HandlerFunc

func NewLogWrapper() Wrapper {
	return func(next HandlerFunc) HandlerFunc {
		return func(ctx context.Context) {
			fmt.Println("before next")
			next(ctx)
			fmt.Println("after next")
		}
	}
}

func Biz(ctx context.Context) {
	fmt.Println("biz")
}

func main() {
	wrappers := []Wrapper{NewLogWrapper()}

	handle := Biz

	for _, wrapper := range wrappers {
		handle = wrapper(handle)
	}

	handle(context.Background())
}
