package transport

import (
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/mikeblum/teapotbot.dev/conf"

	"github.com/sirupsen/logrus"
)

const (
	MaxIdleConns        = 100
	MaxConnsPerHost     = 100
	MaxIdleConnsPerHost = 100
	TimeoutSeconds      = 30
)

type Transport struct {
	http.Transport
	ctx  *http.Request
	host string
	log  *logrus.Entry
}

func NewTransport(confName string) *Transport {
	log := conf.NewLog(confName)
	log.Logger.SetLevel(logrus.DebugLevel)
	t := &Transport{
		log: log,
	}
	t.MaxIdleConns = MaxIdleConns
	t.MaxConnsPerHost = MaxConnsPerHost
	t.MaxIdleConnsPerHost = MaxIdleConnsPerHost
	return t
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.ctx = req
	return http.DefaultTransport.RoundTrip(req)
}

func (t *Transport) Trace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		DNSStart:          t.DNSStart,
		DNSDone:           t.DNSDone,
		ConnectStart:      t.ConnectStart,
		ConnectDone:       t.ConnectDone,
		TLSHandshakeStart: t.TLSHandshakeStart,
		TLSHandshakeDone:  t.TLSHandshakeDone,
		GetConn:           t.GetConn,
		GotConn:           t.GotConn,
	}
}

func (t *Transport) Client() *http.Client {
	return &http.Client{
		Timeout:   TimeoutSeconds * time.Second,
		Transport: t,
	}
}

// DNSStart is called when a DNS lookup begins.
func (t *Transport) DNSStart(info httptrace.DNSStartInfo) {
	t.host = info.Host
	t.log.WithFields(logrus.Fields{
		"host": t.host,
	}).Debug("DNS lookup starting")
}

// DNSDone is called when a DNS lookup ends.
func (t *Transport) DNSDone(info httptrace.DNSDoneInfo) {
	t.log.WithFields(logrus.Fields{
		"host":  t.host,
		"addrs": info.Addrs,
	}).WithError(info.Err).Debug("DNS lookup done")
}

// ConnectStart is called when a new connection's Dial begins.
// If net.Dialer.DualStack (IPv6 "Happy Eyeballs") support is
// enabled, this may be called multiple times.
func (t *Transport) ConnectStart(network, addr string) {
	t.log.WithFields(logrus.Fields{
		"network": network,
		"addr":    addr,
	}).Debug("connection starting")
}

// ConnectDone is called when a new connection's Dial
// completes. The provided err indicates whether the
// connection completed successfully.
// If net.Dialer.DualStack ("Happy Eyeballs") support is
// enabled, this may be called multiple times.
func (t *Transport) ConnectDone(network, addr string, err error) {
	t.log.WithFields(logrus.Fields{
		"network": network,
		"addr":    addr,
	}).WithError(err).Debug("connection done")
}

// TLSHandshakeStart is called when the TLS handshake is started. When
// connecting to an HTTPS site via an HTTP proxy, the handshake happens
// after the CONNECT request is processed by the proxy.
func (t *Transport) TLSHandshakeStart() {
	t.log.Debug("tls handshake starting")
}

// TLSHandshakeDone is called after the TLS handshake with either the
// successful handshake's connection state, or a non-nil error on handshake
// failure.
func (t *Transport) TLSHandshakeDone(state tls.ConnectionState, err error) {
	t.log.WithFields(logrus.Fields{
		"host":         t.host,
		"complete":     state.HandshakeComplete,
		"protocol":     state.NegotiatedProtocol,
		"tls_version":  state.Version,
		"cipher_suite": state.CipherSuite,
	}).WithError(err).Debug("tls handshake done")
}

// GetConn is called before a connection is created or
// retrieved from an idle pool. The hostPort is the
// "host:port" of the target or proxy. GetConn is called even
// if there's already an idle cached connection available.
func (t *Transport) GetConn(hostPort string) {
	t.log.WithFields(logrus.Fields{
		"hostPort": hostPort,
	}).Debug("get conn")
}

// GotConn is called after a successful connection is
// obtained. There is no hook for failure to obtain a
// connection; instead, use the error from
// Transport.RoundTrip.
func (t *Transport) GotConn(info httptrace.GotConnInfo) {
	t.log.WithFields(logrus.Fields{
		"local_addr":  info.Conn.LocalAddr(),
		"remote_addr": info.Conn.RemoteAddr(),
		"idle":        info.WasIdle,
		"reused":      info.Reused,
		"idle_time":   info.IdleTime,
	}).Debug("got conn")
}

func (t *Transport) Do() *Transport {
	req, _ := http.NewRequest("GET", "https://google.com", nil)
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), t.Trace()))

	client := t.Client()
	if res, err := client.Do(req); err != nil {
		t.log.WithError(err).Error("failed to crawl request")
	} else {
		defer res.Body.Close()
	}
	return t
}
