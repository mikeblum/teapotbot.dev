package bot

import (
	"net/http"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"
	buf "teapotbot.dev/api/internal/ping"
)

const (
	MaxIdleConns        = 100
	MaxConnsPerHost     = 100
	MaxIdleConnsPerHost = 100
	TimeoutSeconds      = 30
)

type Bot struct {
	*buf.PingRequest
	http.Client
}

func New() *Bot {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = MaxIdleConns
	t.MaxConnsPerHost = MaxConnsPerHost
	t.MaxIdleConnsPerHost = MaxIdleConnsPerHost

	return &Bot{
		nil,
		http.Client{
			Timeout:   TimeoutSeconds * time.Second,
			Transport: t,
		},
	}
}

func (b *Bot) Proto() protoreflect.Message {
	return b.ProtoReflect()
}
