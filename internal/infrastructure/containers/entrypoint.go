package containers

import (
	"github.com/vitaodemolay/poc-generic-payload/internal/infrastructure/web/controllers/shapes"
	"github.com/vitaodemolay/poc-generic-payload/internal/infrastructure/web/entrypoint"
)

type Entrypoint struct {
	ShapesController *shapes.Controller
}

func NewEntrypointContainer(appContainer *Application) *Entrypoint {
	return &Entrypoint{
		ShapesController: shapes.NewController(appContainer.ShapeService),
	}
}

func (e *Entrypoint) GetControllers() []entrypoint.Router {
	return []entrypoint.Router{
		e.ShapesController,
	}
}
