package hawk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestContextParam(t *testing.T) {
	router := Default()

	var id string

	router.Get("/users/:id", func(c *Context) {
		id = c.Param("id")
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/users/123",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if id != "123" {
		t.Fatalf("expected id 123, got %s", id)
	}
}

func TestContextQuery(t *testing.T) {
	router := Default()

	var page string
	var limit string

	router.Get("/users", func(c *Context) {
		page = c.Query("page")
		limit = c.Query("limit")
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/users?page=2&limit=10",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if page != "2" {
		t.Fatalf("expected page 2, got %s", page)
	}

	if limit != "10" {
		t.Fatalf("expected limit 10, got %s", limit)
	}
}

func TestContextQueryMissing(t *testing.T) {
	router := Default()

	var value string

	router.Get("/users", func(c *Context) {
		value = c.Query("page")
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if value != "" {
		t.Fatalf("expected empty value, got %s", value)
	}
}

func TestContextJSON(t *testing.T) {
	router := Default()

	router.Get("/users", func(c *Context) {
		c.JSON(http.StatusOK, map[string]any{
			"name":  "Adam",
			"email": "adam@test.com",
		})
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

	contentType := rec.Header().Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {
		t.Fatalf(
			"expected JSON content type, got %s",
			contentType,
		)
	}

	var response map[string]any

	err := json.NewDecoder(rec.Body).Decode(&response)

	if err != nil {
		t.Fatal(err)
	}

	if response["name"] != "Adam" {
		t.Fatalf("expected Adam, got %v", response["name"])
	}
}

func TestContextString(t *testing.T) {
	router := Default()

	router.Get("/hello", func(c *Context) {
		c.String(http.StatusOK, "Hello Adam")
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/hello",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if rec.Body.String() != "Hello Adam" {
		t.Fatalf(
			"expected Hello Adam, got %s",
			rec.Body.String(),
		)
	}
}

func TestContextHTML(t *testing.T) {
	router := Default()

	router.Get("/hello", func(c *Context) {
		c.HTML(
			http.StatusOK,
			"<h1>Hello Adam</h1>",
		)
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/hello",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")

	if !strings.Contains(contentType, "text/html") {
		t.Fatalf(
			"expected HTML content type, got %s",
			contentType,
		)
	}

	if rec.Body.String() != "<h1>Hello Adam</h1>" {
		t.Fatalf(
			"unexpected HTML response: %s",
			rec.Body.String(),
		)
	}
}