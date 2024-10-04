package server

import (
	"context"
	"log"

	"worker/grpc/proto"

	"google.golang.org/grpc/metadata"
)

// Executes the business logic for phone registration.
func (srv *Server) Call(ctx context.Context, req *proto.CallRequest) (*proto.ActionResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("Method: Call, Metadata: %v, Request: %v", md, req)

	log.Printf("\tSIP ID: %s Calling RemoteNumber %s\n", req.SipId, req.RemoteNumber)

	return &proto.ActionResponse{
		Message: "Call initiated to " + req.RemoteNumber,
	}, nil
}
