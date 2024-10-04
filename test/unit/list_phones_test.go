package unit_test

import (
	"context"
	"testing"

	"worker/grpc/proto"
	"worker/grpc/server"

	"github.com/stretchr/testify/assert"
)

func TestListPhones(t *testing.T) {
	server := &server.Server{}

	req := &proto.ListPhonesRequest{
		SipId:           "test_sip",
		TenantId:        "tenant1",
		TelephonicState: "idle",
	}

	resp, err := server.ListPhones(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int32(1), resp.Total)
	assert.Equal(t, "123", resp.Phones[0].SipId)
}
