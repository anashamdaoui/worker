package server

import (
	"context"
	"log"

	"worker/grpc/proto"

	"google.golang.org/grpc/metadata"
)

// Executes the business logic for phone registration.
func (srv *server) RegisterPhone(ctx context.Context, req *proto.RegisterPhoneRequest) (*proto.RegisterPhoneResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("Method: RegisterPhone, Metadata: %v, Request: %v", md, req)

	log.Printf("\tSIP ID: %s on Platform URI: %s\n", req.SipId, req.PlatformUri)

	return &proto.RegisterPhoneResponse{
		Message: "Phone registered successfully",
	}, nil
}
