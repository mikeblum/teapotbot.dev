package Crawl

import (
	"fmt"
	"testing"

	buf "github.com/mikeblum/teapotbot.dev/api/internal/crawl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
)

type mockedCrawlRequest struct {
	mock.Mock
}

func (m *mockedCrawlRequest) Crawl() (*buf.CrawlResponse, error) {
	args := m.Called()
	return nil, args.Error(1)
}

func defaultCrawl() *buf.CrawlRequest {
	return &buf.CrawlRequest{
		Method:  buf.Method_GET,
		Url:     &buf.Url{Host: "localhost:80"},
		Timeout: &durationpb.Duration{Seconds: 10},
	}
}

func TestNewCrawlRequestMock(t *testing.T) {
	mocked := new(mockedCrawlRequest)
	mocked.On("Crawl", mock.Anything).Return(nil, fmt.Errorf("Crawl error"))
	response, err := mocked.Crawl()
	mocked.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestNewCrawlRequestDefaults(t *testing.T) {
	req := defaultCrawl()
	assert.Equal(t, buf.Scheme_TCP, req.Url.Scheme, "defaults to tcp")
	assert.Equal(t, buf.Method_GET, req.Method, "defaults to GET")
}

func TestNewCrawlRequestProto(t *testing.T) {
	t.Run("serialize CrawlRequest", func(t *testing.T) {
		req := defaultCrawl()
		dat, err := proto.Marshal(req)
		assert.Nil(t, err)
		assert.True(t, len(dat) > 0)
	})
	t.Run("deserialize CrawlRequest", func(t *testing.T) {
		req := defaultCrawl()
		dat, err := proto.Marshal(req)
		assert.Nil(t, err)
		expected := &buf.CrawlRequest{}
		err = proto.Unmarshal(dat, expected)
		assert.Nil(t, err)
	})
}
