package server

import (
	"context"
	"log"
	"worker/grpc/proto"

	"google.golang.org/grpc/metadata"
)

func (s *Server) ListPhones(ctx context.Context, req *proto.ListPhonesRequest) (*proto.PhoneListResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("Method: ListPhones, Metadata: %v, Request: %v", md, req)

	phones := []*proto.Phone{
		{SipId: "123", PlatformUri: "sample.platform.com", TenantName: "Tenant 1", TelephonicState: "idle"},
	}
	return &proto.PhoneListResponse{
		Phones:   phones,
		Total:    int32(len(phones)),
		Page:     1,
		PageSize: 10,
	}, nil
}
