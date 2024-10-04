package server

import (
	"context"
	"log"

	"worker/grpc/proto"

	"google.golang.org/grpc/metadata"
)

// Executes the business logic for phone deregistration.
func (s *Server) UnregisterPhone(ctx context.Context, req *proto.UnregisterPhoneRequest) (*proto.ActionResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("Method: RegisterPhone, Metadata: %v, Request: %v", md, req)

	log.Printf("\tSIP ID: %s\n", req.SipId)

	return &proto.ActionResponse{
		Message: "Phone unregistered successfully for SIP ID " + req.SipId,
	}, nil
}
