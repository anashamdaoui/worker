package server

import (
	"log"
	"net"
	"strconv"
	"worker/grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedSoftPhoneServiceServer
}

func StartGRPCServer(port int) {

	portStr := strconv.Itoa(port)
	lis, err := net.Listen("tcp", ":"+portStr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterSoftPhoneServiceServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
