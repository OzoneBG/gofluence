package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ozonebg/gofluence/pkg/api"
	"google.golang.org/grpc"
)

//RunServer does what it says.
func RunServer(ctx context.Context, userapi api.UserServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	api.RegisterUserServiceServer(server, userapi)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
