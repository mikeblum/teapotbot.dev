package ping

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	buf "teapotbot.dev/api/internal/ping"
)

const (
	headerContentType = "Content-Type"
	contentTypeText   = "application/text"
	tcpNetwork        = "tcp"
	ipv4Addr          = "127.0.0.1:0"
	ipv6Addr          = "[::]:0"
	routePing         = "/ping"
	responsePong      = "PONG"
	jitterMaxMs       = 500
	protocolV4        = "ipv4"
	protocolV6        = "ipv6"
)

type TCPTestSuite struct {
	suite.Suite
	srv  *http.Server
	ipv4 net.Listener
	ipv6 net.Listener
}

func serve(suite *TCPTestSuite, protocol string, listener net.Listener) {
	if err := suite.srv.Serve(listener); err != nil {
		log.Printf("[%s] serve: %v", protocol, err)
	}
	log.Printf("[%s] listening @ %s", protocol, listener.Addr().String())
}

func createTCPListener(suite *TCPTestSuite, network, address string) (net.Listener, error) {
	srv, err := net.Listen(network, address)
	suite.Assert().Nil(err, fmt.Sprintf("[ network: %s; address: %s ] test server", network, address))
	return srv, err
}

func createTestURL(suite *TCPTestSuite, listener net.Listener) *url.URL {
	return &url.URL{
		Scheme: buf.Scheme_HTTP.String(),
		Host:   listener.Addr().String(),
		Path:   routePing,
	}
}

func ping(rw http.ResponseWriter, r *http.Request) {
	maxMs := big.NewInt(jitterMaxMs)
	if jitter, err := rand.Int(rand.Reader, maxMs); err != nil {
		panic("failed to read crypto/rand")
	} else {
		jitter := time.Duration(jitter.Int64()) * time.Millisecond
		log.Printf("jittering ping by %dms", jitter.Milliseconds())
		time.Sleep(jitter)
	}
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set(headerContentType, contentTypeText)
	if _, err := rw.Write([]byte(responsePong)); err != nil {
		panic(fmt.Sprintf("failed to resolve /ping: %v", err))
	}

}

func (suite *TCPTestSuite) SetupSuite() {
	log.Print("spinning up tcp test servers")
	srv := &http.Server{
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
	suite.srv = srv
	ipv4, _ := createTCPListener(suite, tcpNetwork, ipv4Addr)
	suite.ipv4 = ipv4

	ipv6, _ := createTCPListener(suite, tcpNetwork, ipv6Addr)
	suite.ipv6 = ipv6

	http.HandleFunc(routePing, ping)

	go serve(suite, protocolV4, suite.ipv4)
	go serve(suite, protocolV6, suite.ipv6)
}

func (suite *TCPTestSuite) TearDownSuite() {
	log.Print("tearing down tcp test servers")
	var err error
	err = suite.ipv4.Close()
	suite.Assert().Nil(err, "ipv4 test server closed")
	err = suite.ipv6.Close()
	suite.Assert().Nil(err, "ipv6 test server closed")
}

func (suite *TCPTestSuite) TestPingIpv4() {
	resp, err := http.Get(createTestURL(suite, suite.ipv4).String())
	suite.Assert().Nil(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	suite.Assert().Nil(err)
	suite.Assert().Equal(responsePong, string(body))
}

func (suite *TCPTestSuite) TestPingIpv6() {
	resp, err := http.Get(createTestURL(suite, suite.ipv6).String())
	suite.Assert().Nil(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	suite.Assert().Nil(err)
	suite.Assert().Equal(responsePong, string(body))
}

func TestTCPTestSuite(t *testing.T) {
	suite.Run(t, new(TCPTestSuite))
}
