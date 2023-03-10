package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	go doSomething(ctx)
	select {
	case v := <-ctx.Done():
		fmt.Println("timeline exceeded of 2 sec", v)
	}
	time.Sleep(time.Second * 3)
}
func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			// err := ctx.Err()     //timeout context deadline exceeded timeline exceeded of 2 sec {}
			// fmt.Println(err)
			return
		default:
			fmt.Println("Doing something bakwaas")
		}
	}
}

//	func main() {
//		ctx := context.Background()
//		//seed some data in ctx
//		ctx = seedContext(ctx)
//		readCtx(ctx)
//	}
func readCtx(ctx context.Context) {
	value := ctx.Value("one")
	fmt.Println(value)
}
func seedContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "one", "111")
	return ctx
}

//store and pass the information accross the differnt layer of application
// context varibale and pass it throuh the main->router->handler->db function

//the ability to cancellation the job inbetween of the execution
//consume restful api.. and if I am not done withe it in within 2 mins I need to cancel that job
//3 queries - 1 by 1 - 2 queries got successfully executed and i got failed
