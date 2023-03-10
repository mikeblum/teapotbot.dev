package transport

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	buf "github.com/mikeblum/teapotbot.dev/api/internal/crawl"
	"github.com/mikeblum/teapotbot.dev/conf"
	"github.com/mikeblum/teapotbot.dev/conftest"
)

const (
	headerContentType = "Content-Type"
	contentTypeText   = "application/text"
	tcpNetwork        = "tcp"
	ipv4Addr          = "127.0.0.1:0"
	ipv6Addr          = "[::]:0"
	routeCrawl        = "/crawl"
	responsePong      = "PONG"
	jitterMaxMs       = 500
	protocolV4        = "ipv4"
	protocolV6        = "ipv6"
)

type TCPTestSuite struct {
	client *http.Client
	log    *logrus.Entry
	srv    *http.Server
	ipv4   net.Listener
	ipv6   net.Listener
}

func (s *TCPTestSuite) serve(protocol string, listener net.Listener) {
	if err := s.srv.Serve(listener); err != nil {
		s.log.WithError(err).Errorf("[%s] serve", protocol)
	}
	s.log.Infof("[%s] listening @ %s", protocol, listener.Addr().String())
}

func (s *TCPTestSuite) Crawl(rw http.ResponseWriter, r *http.Request) {
	maxMs := big.NewInt(jitterMaxMs)
	if jitter, err := rand.Int(rand.Reader, maxMs); err != nil {
		panic("failed to read crypto/rand")
	} else {
		jitter := time.Duration(jitter.Int64()) * time.Millisecond
		s.log.Infof("jittering crawl by %dms", jitter.Milliseconds())
		time.Sleep(jitter)
	}
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set(headerContentType, contentTypeText)
	if _, err := rw.Write([]byte(responsePong)); err != nil {
		panic(fmt.Sprintf("failed to resolve /crawl: %v", err))
	}
}

// Utils

func createTCPListener(t *testing.T, network, address string) (net.Listener, error) {
	srv, err := net.Listen(network, address)
	assert.Nil(t, err, fmt.Sprintf("[ network: %s; address: %s ] test server", network, address))
	return srv, err
}

func createTestURL(listener net.Listener) *url.URL {
	return &url.URL{
		Scheme: buf.Scheme_HTTP.String(),
		Host:   listener.Addr().String(),
		Path:   routeCrawl,
	}
}

// Setup

func setupSuite(t *testing.T) (*TCPTestSuite, func(t *testing.T, suite *TCPTestSuite)) {
	_, err := conftest.SetupConf()
	assert.Nil(t, err)
	suite := &TCPTestSuite{
		client: NewTransport(conftest.TestConfFile).Client(),
		log:    conf.NewLog(conftest.TestConfFile),
	}
	srv := &http.Server{
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
	suite.srv = srv
	ipv4, _ := createTCPListener(t, tcpNetwork, ipv4Addr)
	suite.ipv4 = ipv4

	ipv6, _ := createTCPListener(t, tcpNetwork, ipv6Addr)
	suite.ipv6 = ipv6

	http.HandleFunc(routeCrawl, suite.Crawl)

	go suite.serve(protocolV4, suite.ipv4)
	go suite.serve(protocolV6, suite.ipv6)

	return suite, teardownSuite
}

func teardownSuite(t *testing.T, suite *TCPTestSuite) {
	conftest.CleanupConf(t)
	var err error
	err = suite.ipv4.Close()
	assert.Nil(t, err, "ipv4 test server closed")
	err = suite.ipv6.Close()
	assert.Nil(t, err, "ipv6 test server closed")
}

// Tests

func TestTransport(t *testing.T) {
	// <setup code>
	suite, teardown := setupSuite(t)
	defer teardown(t, suite)
	t.Run("transport=do", TransportDoTest)
	t.Run("IPV=4", suite.CrawlIpv4Test)
	t.Run("IPV=6", suite.CrawlIpv6Test)
}

func TransportDoTest(t *testing.T) {
	transport := NewTransport(conftest.TestConfFile)
	ctx := transport.Do()
	assert.NotNil(t, ctx)
}

func (s *TCPTestSuite) CrawlIpv4Test(t *testing.T) {
	resp, err := s.client.Get(createTestURL(s.ipv4).String())
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, responsePong, string(body))
}

func (s *TCPTestSuite) CrawlIpv6Test(t *testing.T) {
	resp, err := s.client.Get(createTestURL(s.ipv6).String())
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, responsePong, string(body))
}
