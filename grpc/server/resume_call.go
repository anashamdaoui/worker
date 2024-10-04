package server

import (
	"context"
	"log"

	"worker/grpc/proto"

	"google.golang.org/grpc/metadata"
)

// Executes the business logic for phone registration.
func (srv *Server) ResumeCall(ctx context.Context, req *proto.CallActionRequest) (*proto.ActionResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("Method: ResumeCall, Metadata: %v, Request: %v", md, req)

	log.Printf("\tCall ID: %s\n", req.CallId)

	return &proto.ActionResponse{
		Message: "Call resumed for call ID " + req.CallId,
	}, nil
}
