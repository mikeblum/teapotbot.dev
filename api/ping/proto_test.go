package ping

import (
	"fmt"
	"testing"

	buf "github.com/mikeblum/teapotbot.dev/api/internal/ping"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
)

type mockedPingRequest struct {
	mock.Mock
}

func (m *mockedPingRequest) Ping() (*buf.PingResponse, error) {
	args := m.Called()
	return nil, args.Error(1)
}

func defaultPing() *buf.PingRequest {
	return &buf.PingRequest{
		Method:  buf.Method_GET,
		Url:     &buf.Url{Host: "localhost:80"},
		Timeout: &durationpb.Duration{Seconds: 10},
	}
}

func TestNewPingRequestMock(t *testing.T) {
	mocked := new(mockedPingRequest)
	mocked.On("Ping", mock.Anything).Return(nil, fmt.Errorf("ping error"))
	response, err := mocked.Ping()
	mocked.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestNewPingRequestDefaults(t *testing.T) {
	req := defaultPing()
	assert.Equal(t, buf.Scheme_TCP, req.Url.Scheme, "defaults to tcp")
	assert.Equal(t, buf.Method_GET, req.Method, "defaults to GET")
}

func TestNewPingRequestProto(t *testing.T) {
	t.Run("serialize PingRequest", func(t *testing.T) {
		req := defaultPing()
		dat, err := proto.Marshal(req)
		assert.Nil(t, err)
		assert.True(t, len(dat) > 0)
	})
	t.Run("deserialize PingRequest", func(t *testing.T) {
		req := defaultPing()
		dat, err := proto.Marshal(req)
		assert.Nil(t, err)
		expected := &buf.PingRequest{}
		err = proto.Unmarshal(dat, expected)
		assert.Nil(t, err)
	})
}
