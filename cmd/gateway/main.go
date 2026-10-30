package main

import (
	"log"

	"github.com/nilesh/gateway-service/internal/gateway"
)

func main() {
	cfg := gateway.LoadConfig()
	app, err := gateway.NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
