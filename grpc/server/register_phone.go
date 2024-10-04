package server

import (
	"context"
	"errors"
	"log"

	"worker/grpc/proto"

	"google.golang.org/grpc/metadata"
)

// Executes the business logic for phone registration.
func (srv *Server) RegisterPhone(ctx context.Context, req *proto.RegisterPhoneRequest) (*proto.RegisterPhoneResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("Method: RegisterPhone, Metadata: %v, Request: %v", md, req)

	log.Printf("\tSIP ID: %s on Platform URI: %s\n", req.SipId, req.PlatformUri)

	if req.SipId == "" || req.Password == "" || req.PlatformUri == "" {
		return nil, errors.New("invalid requests: missing fields")
	}

	return &proto.RegisterPhoneResponse{
		Message: "Phone registered successfully",
		PhoneId: "12345",
	}, nil
}
