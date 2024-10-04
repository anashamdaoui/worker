package unit_test

import (
	"context"
	"testing"

	"worker/grpc/proto"
	"worker/grpc/server"

	"github.com/stretchr/testify/assert"
)

func TestUnregisterPhone(t *testing.T) {
	server := &server.Server{}

	req := &proto.UnregisterPhoneRequest{
		SipId: "sip_400.enterprise1",
	}

	resp, err := server.UnregisterPhone(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Phone unregistered successfully for SIP ID sip_400.enterprise1", resp.Message)
}
