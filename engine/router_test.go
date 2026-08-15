package hawk

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestGETRoute(t *testing.T) {

	router := Default()

	called := false

	router.Get("/users", func(c *Context) {
		called = true
	})

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if !called {
		t.Fatal("expected handler to be called")
	}
}

func TestPOSTRoute(t *testing.T) {

	router := Default()

	called := false

	router.Post("/users", func(c *Context) {
		called = true
	})
	req := httptest.NewRequest(http.MethodPost, "/users", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if !called {
		t.Fatal("expected handler to be called")
	}
}

func TestPUTRoute(t *testing.T) {

	router := Default()

	called := false

	router.Put("/users", func(c *Context) {
		called = true
	})

	req := httptest.NewRequest(http.MethodPut, "/users", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if !called {
		t.Fatal("expected handler to be called")
	}

}

func TestDELETERoute(t *testing.T) {


	router := Default()
	called := false

	router.Delete("/users", func(c *Context) {
		called = true
	})

	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if !called {
		t.Fatal("expected handler to be called")
	}

}

func TestRouteNotFound(t *testing.T) {
	router := Default()

	called := false

	router.Get("/users", func(c *Context) {
		called = true
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/posts",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status 404, got %d",
			rec.Code,
		)
	}

	if called {
		t.Fatal("expected handler not to be called")
	}
}

func TestWrongHTTPMethod(t *testing.T) {
	router := Default()

	called := false

	router.Post("/users", func(c *Context) {
		called = true
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status 404, got %d",
			rec.Code,
		)
	}

	if called {
		t.Fatal("expected POST handler not to be called")
	}
}

func TestStaticRouteTakesPriorityOverParameterRoute(t *testing.T) {
	router := Default()

	var matched string

	router.Get("/users/:id", func(c *Context) {
		matched = "parameter"
	})

	router.Get("/users/profile", func(c *Context) {
		matched = "static"
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

	if matched != "static" {
		t.Fatalf(
			"expected static route to match, got %s",
			matched,
		)
	}
}

func TestStaticPathTraversal(t *testing.T) {
	router := Default()

	dir := t.TempDir()

	secretDir := t.TempDir()

	err := os.WriteFile(
		filepath.Join(secretDir, "secret.txt"),
		[]byte("secret"),
		0644,
	)
	if err != nil {
		t.Fatal(err)
	}

	router.Static("/assets", dir)

	req := httptest.NewRequest(
		http.MethodGet,
		"/assets/../../secret.txt",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code == http.StatusOK {
		t.Fatal("static handler allowed path traversal")
	}
}

func TestStaticOnlyGET(t *testing.T) {
	router := Default()

	dir := t.TempDir()

	err := os.WriteFile(
		filepath.Join(dir, "hello.txt"),
		[]byte("Hello"),
		0644,
	)
	if err != nil {
		t.Fatal(err)
	}

	router.Static("/assets", dir)

	req := httptest.NewRequest(
		http.MethodPost,
		"/assets/hello.txt",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status 404, got %d",
			rec.Code,
		)
	}
}

func TestStaticNestedFile(t *testing.T) {
	router := Default()

	dir := t.TempDir()

	err := os.MkdirAll(
		filepath.Join(dir, "images"),
		0755,
	)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(
		filepath.Join(dir, "images", "logo.txt"),
		[]byte("logo"),
		0644,
	)
	if err != nil {
		t.Fatal(err)
	}

	router.Static("/assets", dir)

	req := httptest.NewRequest(
		http.MethodGet,
		"/assets/images/logo.txt",
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
}

func TestStaticFileNotFound(t *testing.T) {
	router := Default()

	dir := t.TempDir()

	router.Static("/assets", dir)

	req := httptest.NewRequest(
		http.MethodGet,
		"/assets/missing.txt",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status 404, got %d",
			rec.Code,
		)
	}
}

func TestStaticFile(t *testing.T) {
	router := Default()

	dir := t.TempDir()

	err := os.WriteFile(
		filepath.Join(dir, "hello.txt"),
		[]byte("Hello Hawk"),
		0644,
	)
	if err != nil {
		t.Fatal(err)
	}

	router.Static("/assets", dir)

	req := httptest.NewRequest(
		http.MethodGet,
		"/assets/hello.txt",
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

	if rec.Body.String() != "Hello Hawk" {
		t.Fatalf(
			"expected Hello Hawk, got %s",
			rec.Body.String(),
		)
	}
}

func TestCategoriesRoute(t *testing.T) {
    router := Default()

    called := false

    router.Get("/categories", func(c *Context) {
        called = true
    })

    req := httptest.NewRequest(
        http.MethodGet,
        "/categories",
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

    if !called {
        t.Fatal("expected categories handler to be called")
    }
}

func TestMatchCategories(t *testing.T) {
    matched, params := match(
        "/categories",
        "/categories",
    )

    if !matched {
        t.Fatal("expected route to match")
    }

    if len(params) != 0 {
        t.Fatalf(
            "expected no params, got %v",
            params,
        )
    }
}

func TestStaticRouteDoesNotOverrideAPIStaticRoute(t *testing.T) {
	router := Default()

	called := false

	router.Static("/", "./public")

	router.Get("/categories", func(c *Context) {
		called = true
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/categories",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if !called {
		t.Fatal("expected categories route to be called")
	}
}