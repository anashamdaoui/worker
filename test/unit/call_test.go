package unit_test

import (
	"context"
	"testing"

	"worker/grpc/proto"
	"worker/grpc/server"

	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	server := &server.Server{}

	req := &proto.CallRequest{
		SipId:        "test_sip",
		RemoteNumber: "123456789",
	}

	resp, err := server.Call(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Call initiated to 123456789", resp.Message)
}
