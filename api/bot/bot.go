package bot

import (
	"net/http"

	"google.golang.org/protobuf/reflect/protoreflect"
	buf "teapotbot.dev/api/internal/ping"
	"teapotbot.dev/api/transport"
)

type Bot struct {
	*buf.PingRequest
	*http.Client
}

func New() *Bot {
	return &Bot{
		nil,
		transport.New().Client(),
	}
}

func (b *Bot) Proto() protoreflect.Message {
	return b.ProtoReflect()
}
