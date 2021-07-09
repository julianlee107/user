package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/julianlee107/user/handler"
	pb "github.com/julianlee107/user/proto/user/pb/user"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
	)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
