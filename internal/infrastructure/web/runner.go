package web

import (
	"context"
	"log"

	"github.com/vitaodemolay/poc-generic-payload/internal/infrastructure/containers"
	"github.com/vitaodemolay/poc-generic-payload/internal/infrastructure/web/service"
)

func Run(ctx context.Context) error {
	port := ":8080"

	webServer, err := service.CreateWebServer(port)
	if err != nil {
		return err
	}

	log.Println("Mounting Dependencies")
	infraContainer, err := containers.NewInfrastructure()
	if err != nil {
		return err
	}

	appContainer, err := containers.NewApplicationWithInjection(infraContainer)
	if err != nil {
		return err
	}

	entrypointContainer := containers.NewEntrypointContainer(appContainer)

	log.Println("Initializing Routes")
	webServer.InitalizeRoutes(entrypointContainer.GetControllers()...)

	log.Println("Starting Web Server on port", port)
	return webServer.Start()
}
