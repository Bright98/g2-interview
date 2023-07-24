package grpc

import (
	"flag"
	"fmt"
	"g2/user/domain"
	pb "g2/user/recieve/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcServer struct {
	pb.UserServiceServer
	domain domain.ServiceInterface
}

func NewGrpcServer(service domain.ServiceInterface) *GrpcServer {
	return &GrpcServer{domain: service}
}

func GrpcServerConnection(port string) error {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &GrpcServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
