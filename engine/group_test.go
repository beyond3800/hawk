package hawk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGroupRoute(t *testing.T) {
	t.Helper()
	router := Default()

	called := false

	group := router.Group("/users")

	group.Get("/profile", func(c *Context) {
		called = true
	})
	req := httptest.NewRequest(http.MethodGet, "/users/profile", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if !called {
		t.Fatal("expected handler to be called")
	}
}

func TestRouteGroupPrefixRequired(t *testing.T) {
	router := Default()

	called := false

	api := router.Group("/api")

	api.Get("/users", func(c *Context) {
		called = true
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if called {
		t.Fatal("expected handler not to be called")
	}
}