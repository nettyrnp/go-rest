package apis

import (
	"testing"
	"net/http"
)

func TestAuth(t *testing.T) {
	router := newRouter()
	router.Post("/auth", Auth("secret"))
	runAPITests(t, router, []apiTestCase{
		{"t1 - successful login", "POST", "/auth", `{"username":"regular_user", "password":"pass"}`, http.StatusOK, ""},
		{"t1 - successful login", "POST", "/auth", `{"username":"admin", "password":"pass2"}`, http.StatusOK, ""},
		{"t3 - unsuccessful login", "POST", "/auth", `{"username":"regular_user", "password":"pass2"}`, http.StatusUnauthorized, ""},
		{"t4 - unsuccessful login", "POST", "/auth", `{"username":"admin", "password":"pass"}`, http.StatusUnauthorized, ""},
	})
}