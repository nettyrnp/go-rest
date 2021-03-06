package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/nettyrnp/go-rest/app"
	"github.com/nettyrnp/go-rest/models"
)

type (
	// userService specifies the interface for the user service needed by userResource.
	userService interface {
		Get(rs app.RequestScope, id int) (*models.User, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.User, error)
		Create(rs app.RequestScope, model *models.User) (*models.User, error)
		Delete(rs app.RequestScope, id int) (*models.User, error)
		Count(rs app.RequestScope) (int, error)
	}

	// userResource defines the handlers for the CRUD APIs.
	userResource struct {
		service userService
	}
)

// ServeUser sets up the routing of user endpoints and the corresponding handlers.
func ServeUserResource(rg *routing.RouteGroup, service userService) {
	r := &userResource{service}
	rg.Get("/users/<id>", r.get)
	rg.Get("/users", r.query)
	rg.Post("/users", r.create)
	rg.Delete("/users/<id>", r.delete)
}

func (r *userResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *userResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}

func (r *userResource) create(c *routing.Context) error {
	var model models.User
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *userResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}
