package unit_test

import (
	"context"
	"testing"

	"worker/grpc/proto"
	"worker/grpc/server"

	"github.com/stretchr/testify/assert"
)

func TestHoldCall(t *testing.T) {
	server := &server.Server{}

	req := &proto.CallActionRequest{
		CallId: "call_123",
	}

	resp, err := server.HoldCall(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Call put on hold for call ID call_123", resp.Message)
}
