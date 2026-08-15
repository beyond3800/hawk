package hawk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	router := Default()

	middlewareCalled := false
	handlerCalled := false

	router.Use(func(c *Context) {
		middlewareCalled = true
	})

	router.Get("/users", func(c *Context) {
		handlerCalled = true
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			rec.Code,
		)
	}

	if !middlewareCalled {
		t.Fatal("expected middleware to be called")
	}

	if !handlerCalled {
		t.Fatal("expected handler to be called")
	}
}

func TestMiddlewareRunsBeforeHandler(t *testing.T) {
	router := Default()

	var order []string

	router.Use(func(c *Context) {
		order = append(order, "middleware")
	})

	router.Get("/users", func(c *Context) {
		order = append(order, "handler")
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if len(order) != 2 {
		t.Fatalf("expected 2 calls, got %d", len(order))
	}

	if order[0] != "middleware" {
		t.Fatalf("expected middleware first, got %s", order[0])
	}

	if order[1] != "handler" {
		t.Fatalf("expected handler second, got %s", order[1])
	}
}

func TestGroupMiddleware(t *testing.T) {
	router := Default()

	middlewareCalled := false
	handlerCalled := false

	group := router.Group("/users")
	group.Use(func(c *Context) {
		middlewareCalled = true
	})

	group.Get("/profile", func(c *Context) {
		handlerCalled = true
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/users/profile",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			rec.Code,
		)
	}

	if !middlewareCalled {
		t.Fatal("expected middleware to be called")
	}

	if !handlerCalled {
		t.Fatal("expected handler to be called")
	}
}

func TestAbortMiddleware(t *testing.T) {
		router := Default()

	authentificated := false

	router.Use(func(c *Context) {
		if !authentificated{
			c.Abort()
		}
	})

	router.Get("/users", func(c *Context) {
		t.Fatal("expected handler not to be called")
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

}

func TestAbortGroupMiddleware(t *testing.T) {
	router := Default()

	authentificated := false

	group := router.Group("/users")
	group.Use(func(c *Context) {
		if !authentificated{
			c.Abort()
		}
	})
	group.Use(func(c *Context) {
		if !authentificated{
			c.Abort()
		}
	})

	group.Get("/profile", func(c *Context) {
		t.Fatal("expected handler not to be called")
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/users/profile",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

}
