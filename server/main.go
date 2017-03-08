package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/mastahyeti/grace/grace"
	"golang.org/x/net/context"
)

// GraceDemoServer implements grace.DemoServer
type GraceDemoServer struct {
	gs *grpc.Server
}

// NewGraceDemoServer creates a new GraceDemoServer.
func NewGraceDemoServer() *GraceDemoServer {
	gs := grpc.NewServer()

	return &GraceDemoServer{gs}
}

// Sleep implements grace.DemoServer
func (gds *GraceDemoServer) Sleep(ctx context.Context, req *grace.SleepRequest) (resp *grace.SleepResponse, _ error) {
	fmt.Println("Starting sleep")
	resp = &grace.SleepResponse{Ok: true}
	time.Sleep(time.Duration(req.GetDuration()))
	fmt.Println("Done sleeping")
	return
}

func (gds *GraceDemoServer) handleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	_ = <-c
	fmt.Println("Shutting down gracefully...")
	gds.gs.GracefulStop()
	fmt.Println("Exiting")
	os.Exit(0)
}

// Run runs the server.
func (gds *GraceDemoServer) Run() {
	fmt.Println("Starting server...")

	lis, err := net.Listen("tcp", grace.Address)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	grace.RegisterDemoServer(gds.gs, gds)

	go gds.handleSignals()

	if err = gds.gs.Serve(lis); err != nil {
		fmt.Println("Error: ", err.Error())
	}
}

func main() {
	gds := NewGraceDemoServer()
	gds.Run()
}
