package hawk


import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestParam(t *testing.T) {
	router := Default()

	called := false

	var userID string
	router.Get("/users/:id", func(c *Context) {
		called = true
		userID = c.Param("id")
	})

	req := httptest.NewRequest(http.MethodGet, "/users/123", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if !called {
		t.Fatal("expected handler to be called")
	}
	if userID != "123" {
		t.Fatalf("expected param id to be 123, got %s", userID)
	}
}

func TestParamNotFound(t *testing.T) {
	router := Default()

	called := false

	router.Get("/users/:id", func(c *Context) {
		called = true
	})

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rec.Code)
	}

	if called {
		t.Fatal("expected handler not to be called")
	}
}

func TestMultipleParams(t *testing.T) {
	router := Default()

	called := false

	var userID, postID string
	router.Get("/users/:userID/posts/:postID", func(c *Context) {
		called = true
		userID = c.Param("userID")
		postID = c.Param("postID")
	})

	req := httptest.NewRequest(http.MethodGet, "/users/123/posts/456", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if !called {
		t.Fatal("expected handler to be called")
	}
	if userID != "123" {
		t.Fatalf("expected param userID to be 123, got %s", userID)
	}
	if postID != "456" {
		t.Fatalf("expected param postID to be 456, got %s", postID)
	}
}

func TestQueryParameter(t *testing.T) {
	router := Default()

	var page string

	router.Get("/users", func(c *Context) {
		page = c.Query("page")
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/users?page=2",
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

	if page != "2" {
		t.Fatalf(
			"expected page=2, got %s",
			page,
		)
	}
}

func TestMultipleQueryParameters(t *testing.T) {
	router := Default()

	var page string
	var search string

	router.Get("/users", func(c *Context) {
		page = c.Query("page")
		search = c.Query("search")
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/users?page=2&search=adam",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if page != "2" {
		t.Fatalf("expected page=2, got %s", page)
	}

	if search != "adam" {
		t.Fatalf("expected search=adam, got %s", search)
	}
}