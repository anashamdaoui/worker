package integration_test

import (
	"context"
	"log"
	"net"
	"testing"
	"worker/grpc/proto"
	"worker/grpc/server"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Helper function to start the gRPC test server
func startTestGRPCServer() (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", ":0") // Random available port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterSoftPhoneServiceServer(grpcServer, &server.Server{})

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	return grpcServer, lis
}

func TestRegisterPhoneIntegration(t *testing.T) {
	grpcServer, lis := startTestGRPCServer()
	defer grpcServer.Stop()

	// Connect to the gRPC server
	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := proto.NewSoftPhoneServiceClient(conn)

	// Make the RegisterPhone request
	req := &proto.RegisterPhoneRequest{
		SipId:       "test_sip",
		Password:    "password",
		PlatformUri: "platform_uri",
	}
	resp, err := client.RegisterPhone(context.Background(), req)

	// Check response and error
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Phone registered successfully", resp.Message)
	assert.Equal(t, "12345", resp.PhoneId)
}

func TestListPhonesIntegration(t *testing.T) {
	grpcServer, lis := startTestGRPCServer()
	defer grpcServer.Stop()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := proto.NewSoftPhoneServiceClient(conn)

	// Make the ListPhones request
	req := &proto.ListPhonesRequest{
		SipId:    "test_sip",
		TenantId: "tenant1",
	}
	resp, err := client.ListPhones(context.Background(), req)

	// Check response
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int32(1), resp.Total)
	assert.Equal(t, "123", resp.Phones[0].SipId)
}

func TestCallIntegration(t *testing.T) {
	grpcServer, lis := startTestGRPCServer()
	defer grpcServer.Stop()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := proto.NewSoftPhoneServiceClient(conn)

	// Make the Call request
	req := &proto.CallRequest{
		SipId:        "test_sip",
		RemoteNumber: "123456789",
	}
	resp, err := client.Call(context.Background(), req)

	// Check response
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Call initiated to 123456789", resp.Message)
}

func TestAnswerCallIntegration(t *testing.T) {
	grpcServer, lis := startTestGRPCServer()
	defer grpcServer.Stop()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := proto.NewSoftPhoneServiceClient(conn)

	// Make the AnswerCall request
	req := &proto.CallActionRequest{
		CallId: "call_123",
	}
	resp, err := client.AnswerCall(context.Background(), req)

	// Check response
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Call answered for call ID call_123", resp.Message)
}

func TestUnregisterPhoneIntegration(t *testing.T) {
	grpcServer, lis := startTestGRPCServer()
	defer grpcServer.Stop()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := proto.NewSoftPhoneServiceClient(conn)

	// Make the UnregisterPhone request
	req := &proto.UnregisterPhoneRequest{
		SipId: "test_sip",
	}
	resp, err := client.UnregisterPhone(context.Background(), req)

	// Check response
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Phone unregistered successfully for SIP ID test_sip", resp.Message)
}
