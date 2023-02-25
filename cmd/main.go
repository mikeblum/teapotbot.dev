package main

import (
	"github.com/mikeblum/teapotbot.dev/conf"
	"github.com/mikeblum/teapotbot.dev/sql"
)

func main() {
	cfg, _ := conf.NewConf(conf.Provider(""))
	log := conf.NewLog("")
	if err := sql.Setup(cfg); err != nil {
		log.WithError(err).Fatal("failed to bootstrap sql backend")
	}
}
