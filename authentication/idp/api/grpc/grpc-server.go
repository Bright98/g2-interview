package grpc

import (
	"flag"
	"fmt"
	"g2/auth/idp/domain"
	pb "g2/proto/idp"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcServer struct {
	pb.IdpServiceServer
	domain domain.ServiceInterface
}

func NewGrpcServer(service domain.ServiceInterface) *GrpcServer {
	return &GrpcServer{domain: service}
}
func GrpcServerConnection(server *GrpcServer, port string) {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	pb.RegisterIdpServiceServer(s, server)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
