package unit_test

import (
	"context"
	"testing"

	"worker/grpc/proto"
	"worker/grpc/server"

	"github.com/stretchr/testify/assert"
)

func TestRegisterPhone(t *testing.T) {
	server := &server.Server{}

	req := &proto.RegisterPhoneRequest{
		SipId:       "test_sip",
		Password:    "password",
		PlatformUri: "platform_uri",
	}

	resp, err := server.RegisterPhone(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Phone registered successfully", resp.Message)
	assert.Equal(t, "12345", resp.PhoneId)
}

func TestRegisterPhone_MissingFields(t *testing.T) {
	server := &server.Server{}

	// Missing sipId and platformUri
	req := &proto.RegisterPhoneRequest{
		SipId:       "",
		Password:    "password",
		PlatformUri: "",
	}

	resp, err := server.RegisterPhone(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
}
