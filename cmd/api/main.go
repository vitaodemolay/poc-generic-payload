package main

import (
	"context"
	"log"

	"github.com/vitaodemolay/poc-generic-payload/internal/infrastructure/web"
)

func main() {
	ctx := context.Background()

	if err := web.Run(ctx); err != nil {
		log.Fatal("Failed to start web server: ", err)
	}
}
