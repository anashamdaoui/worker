package server

import (
	"context"
	"errors"
	"log"

	"worker/grpc/proto"

	"google.golang.org/grpc/metadata"
)

// Executes the business logic for phone registration.
func (srv *Server) AnswerCall(ctx context.Context, req *proto.CallActionRequest) (*proto.ActionResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("Method: AnswerCall, Metadata: %v, Request: %v", md, req)

	log.Printf("\tCall ID: %s\n", req.CallId)

	if req.CallId == "" {
		return nil, errors.New("invalid Call ID")
	}

	return &proto.ActionResponse{
		Message: "Call answered for call ID " + req.CallId,
	}, nil
}
