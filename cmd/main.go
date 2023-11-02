package main

import (
	"fmt"
	"log"
	"net"

	"github.com/neel2468/go-grpc-order-svc/pkg/client"
	"github.com/neel2468/go-grpc-order-svc/pkg/config"
	"github.com/neel2468/go-grpc-order-svc/pkg/db"
	"github.com/neel2468/go-grpc-order-svc/pkg/pb"
	"github.com/neel2468/go-grpc-order-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("failed to listening", err)
	}

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	if err != nil {
		log.Fatalln("failed to listening", err)
	}

	fmt.Println("Order Svc on", c.Port)

	s := services.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
