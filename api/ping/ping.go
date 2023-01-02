package ping

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	buf "teapotbot.dev/api/internal/ping"
)

type Bot struct {
	buf.PingRequest
}

func (b *Bot) Proto() protoreflect.Message {
	return b.ProtoReflect()
}

func (b *Bot) Ping() (*buf.PingResponse, error) {
	return &buf.PingResponse{
		Timeout: b.Timeout,
	}, nil
}
