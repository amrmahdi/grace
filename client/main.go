package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mastahyeti/grace/grace"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(grace.Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	defer conn.Close()

	go func() {
		for {
			fmt.Printf("conn-state: %v\n", conn.GetState())
			conn.WaitForStateChange(context.Background(), conn.GetState())
		}
	}()

	c := grace.NewDemoClient(conn)

	resp, err := c.Sleep(context.Background(), &grace.SleepRequest{Duration: uint64(time.Minute)})
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println("Result: ", resp.GetOk())
}
