package shapes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/vitaodemolay/poc-generic-payload/internal/application/shapes"
	"github.com/vitaodemolay/poc-generic-payload/internal/application/shapes/contracts"
	"github.com/vitaodemolay/poc-generic-payload/internal/infrastructure/web/entrypoint"
	internalerrors "github.com/vitaodemolay/poc-generic-payload/pkg/internal-errors"
)

type Controller struct {
	service shapes.ShapeService
}

func NewController(service shapes.ShapeService) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) Path() string {
	return "/v1/shapes"
}

func (c *Controller) GetRoutes() []entrypoint.Route {
	return []entrypoint.Route{
		{
			Method:  http.MethodPost,
			Pattern: "/",
			Handler: c.CreateShape,
		},
		{
			Method:  http.MethodGet,
			Pattern: "/{id}",
			Handler: c.GetShapeById,
		},
		{
			Method:  http.MethodPut,
			Pattern: "/{id}/position",
			Handler: c.ChangeShapePosition,
		},
	}
}

func (c *Controller) CreateShape(w http.ResponseWriter, r *http.Request) (any, int, error) {
	var request contracts.CreateShape
	render.DecodeJSON(r.Body, &request)
	id, err := c.service.CreateShape(r.Context(), request)
	if err != nil {
		return nil, 0, err
	}

	return map[string]string{"shape_id": id}, http.StatusCreated, nil
}

func (c *Controller) GetShapeById(w http.ResponseWriter, r *http.Request) (any, int, error) {
	shapeID := chi.URLParam(r, "id")
	if shapeID == "" {
		return nil, 0, internalerrors.ErrBadRequest
	}

	shape, err := c.service.GetShape(r.Context(), shapeID)
	if err != nil {
		return nil, 0, err
	}

	return shape, http.StatusOK, nil
}

func (c *Controller) ChangeShapePosition(w http.ResponseWriter, r *http.Request) (any, int, error) {
	shapeID := chi.URLParam(r, "id")
	if shapeID == "" {
		return nil, 0, internalerrors.ErrBadRequest
	}
	var request contracts.ChangeShapePosition
	render.DecodeJSON(r.Body, &request)
	request.ID = shapeID

	err := c.service.ChangeShapePosition(r.Context(), request)
	if err != nil {
		return nil, 0, err
	}

	return map[string]string{"message": "Shape position updated successfully"}, http.StatusOK, nil
}
