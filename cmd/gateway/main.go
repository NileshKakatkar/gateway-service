package main

import (
	"log"

	"github.com/nilesh/gateway-service/internal/gateway"
)

func main() {
	cfg := gateway.LoadConfig()
	app := gateway.NewApp(cfg)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
