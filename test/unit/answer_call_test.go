package unit_test

import (
	"context"
	"testing"

	"worker/grpc/proto"
	"worker/grpc/server"

	"github.com/stretchr/testify/assert"
)

func TestAnswerCall(t *testing.T) {
	server := &server.Server{}

	req := &proto.CallActionRequest{
		CallId: "call_123",
	}

	resp, err := server.AnswerCall(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Call answered for call ID call_123", resp.Message)
}

func TestAnswerCall_MissingCallID(t *testing.T) {
	server := &server.Server{}

	req := &proto.CallActionRequest{
		CallId: "",
	}

	resp, err := server.AnswerCall(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
}
