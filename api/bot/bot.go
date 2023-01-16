package bot

import (
	"net/http"

	"teapotbot.dev/api/transport"
)

type Bot struct {
	*http.Client
}

func New() *Bot {
	return &Bot{
		transport.New().Client(),
	}
}
