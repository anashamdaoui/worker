package integration

import (
	"context"
	"net"
	"testing"
	"worker/proto"
	"worker/worker/handlers"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	proto.RegisterSoftPhoneServiceServer(s, handlers.NewSoftPhoneService())

	go func() {
		if err := s.Serve(lis); err != nil {
			panic("Server exited with error: " + err.Error())
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestRegisterPhoneIntegration(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := proto.NewSoftPhoneServiceClient(conn)

	req := &proto.RegisterPhoneRequest{
		SipId:       "1234",
		Password:    "secret",
		PlatformUri: "http://example.com",
	}

	resp, err := client.RegisterPhone(ctx, req)
	if err != nil {
		t.Fatalf("RegisterPhone failed: %v", err)
	}
	if resp.PhoneId != "12345" || resp.Message != "Phone registered successfully" {
		t.Errorf("Unexpected response: %v", resp)
	}
}
