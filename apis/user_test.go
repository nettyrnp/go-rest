package apis

import (
	"net/http"
	"testing"

	"github.com/nettyrnp/go-rest/daos"
	"github.com/nettyrnp/go-rest/services"
	"github.com/nettyrnp/go-rest/migrate"
)

func TestUser(t *testing.T) {
	migrate.ResetDB()
	router := newRouter()
	ServeUserResource(&router.RouteGroup, services.NewUserService(daos.NewUserDAO()))

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	nameRequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"name","error":"cannot be blank"}]}`

	runAPITests(t, router, []apiTestCase{
		{"t1 - get a user", "GET", "/users/2", "", http.StatusOK, `{"id":2,"name":"Some Name 2","role":"user"}`},
		{"t2 - get a nonexisting user", "GET", "/users/99999", "", http.StatusNotFound, notFoundError},
		{"t3 - create a user", "POST", "/users", `{"name":"John Dow","role":"admin"}`, http.StatusOK, `{"id": 34, "name":"John Dow","role":"admin"}`},
		{"t4 - create a user with validation error", "POST", "/users", `{"name":"","role":"user"}`, http.StatusBadRequest, nameRequiredError},
		//{"t7 - delete a user", "DELETE", "/users/2", ``, http.StatusOK, `{"id": 2, "name":"John Dow"}`},
		//{"t8 - delete a nonexisting user", "DELETE", "/users/99999", "", http.StatusNotFound, notFoundError},
	})
}
