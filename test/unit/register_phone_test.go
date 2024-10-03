package handlers_test

import (
	"context"
	"testing"

	"worker/proto"
	"worker/worker/handlers"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRegisterPhone(t *testing.T) {
	srv := handlers.NewSoftPhoneService()
	req := &proto.RegisterPhoneRequest{
		SipId:       "1234",
		Password:    "secret",
		PlatformUri: "http://example.com",
	}

	// Call the method
	resp, err := srv.RegisterPhone(context.Background(), req)
	if err != nil {
		t.Errorf("RegisterPhone returned an unexpected error: %v", err)
	}

	// Check if the response is as expected
	if resp.Message != "Phone registered successfully" {
		t.Errorf("Expected success message, got %s instead", resp.Message)
	}

	// Optionally check for the specific phone ID if necessary
	if resp.PhoneId != "12345" { // This should match the mocked response in the handler
		t.Errorf("Expected phone ID '12345', got %s instead", resp.PhoneId)
	}

	// Test for a failure case with missing fields that are required
	reqBad := &proto.RegisterPhoneRequest{
		SipId:       "", // Missing SIP ID, assuming it's required
		Password:    "", // Missing password, assuming it's required
		PlatformUri: "", // Missing or incorrect URI, assuming it's required
	}
	_, errBad := srv.RegisterPhone(context.Background(), reqBad)
	if st, ok := status.FromError(errBad); !ok || st.Code() != codes.InvalidArgument {
		t.Errorf("Expected InvalidArgument, got %v", st.Code())
	}

}
