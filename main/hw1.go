package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "some", "123")
	fmt.Printf("%#v\n", ctx)
	val := ctx.Value("some")
	fmt.Println(val)
	SomeFunc(ctx)

}

func SomeFunc(ctx context.Context) {
	ctx, cancelFunc := context.WithTimeout(ctx, 4*time.Second)
	defer func() {
		fmt.Println("cancelFunc()")
		cancelFunc()
	}()
	fmt.Println(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}
