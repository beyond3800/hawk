# Response

Hawk provides a simple response API through the `Context`.

The main response methods are:

```go
c.JSON()
c.String()
c.HTML()
```

The underlying Go response writer is also available through:

```go
c.Response
```

This allows Hawk applications to use both Hawk's convenience methods and Go's standard `net/http` response functionality.

---

# Response Overview

A typical request flows through Hawk like this:

```text
HTTP Request
     ↓
   Router
     ↓
 Middleware
     ↓
 Route Handler
     ↓
   Response
```

The route handler uses the `Context` to send the response.

For example:

```go
router.Get("/users", func(c *hawk.Context) {

    c.JSON(200, map[string]any{
        "message": "Users",
    })
})
```

---

# JSON Responses

Use `JSON()` when returning JSON data.

```go
c.JSON(200, map[string]any{
    "message": "Hello from Hawk",
})
```

The first argument is the HTTP status code.

The second argument is the data Hawk should encode as JSON.

For example:

```go
router.Get("/user", func(c *hawk.Context) {

    c.JSON(200, map[string]any{
        "id":    1,
        "name":  "Adam",
        "email": "adam@test.com",
    })
})
```

The response is JSON:

```json
{
    "id": 1,
    "name": "Adam",
    "email": "adam@test.com"
}
```

---

# Returning a Struct

JSON responses can also be created from Go structs.

```go
type User struct {
    ID    int64  `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

Then:

```go
router.Get("/user", func(c *hawk.Context) {

    user := User{
        ID:    1,
        Name:  "Adam",
        Email: "adam@test.com",
    }

    c.JSON(200, user)
})
```

The JSON tags determine the JSON field names.

---

# Returning a Slice

You can return multiple records as a JSON array.

```go
users := []User{
    {
        ID:   1,
        Name: "Adam",
    },
    {
        ID:   2,
        Name: "John",
    },
}
```

Return the slice:

```go
c.JSON(200, users)
```

The response will contain a JSON array.

---

# HTTP Status Codes

The status code is passed as the first argument to the response methods.

For example:

```go
c.JSON(200, data)
```

means:

```text
200 OK
```

Common status codes include:

|  Code | Meaning               |
| ----: | --------------------- |
| `200` | OK                    |
| `201` | Created               |
| `204` | No Content            |
| `400` | Bad Request           |
| `401` | Unauthorized          |
| `403` | Forbidden             |
| `404` | Not Found             |
| `405` | Method Not Allowed    |
| `409` | Conflict              |
| `422` | Unprocessable Entity  |
| `500` | Internal Server Error |

Use the status code that accurately describes the result of the request.

---

# Creating Resources

When an API successfully creates a resource, `201 Created` is commonly appropriate.

```go
router.Post("/users", func(c *hawk.Context) {

    user := User{
        ID:   1,
        Name: "Adam",
    }

    c.JSON(201, resources.Users(user).ToSlice())
})
```

---

# Error Responses

JSON is also useful for returning errors.

```go
c.JSON(400, map[string]any{
    "error": "Invalid request",
})
```

For an unauthenticated request:

```go
c.JSON(401, map[string]any{
    "error": "Unauthorized",
})
```

For a missing resource:

```go
c.JSON(404, map[string]any{
    "error": "User not found",
})
```

A consistent error structure makes APIs easier for clients to consume.

For example:

```json
{
    "error": "User not found"
}
```

---

# String Responses

Use `String()` when you want to return plain text.

```go
c.String(200, "Hello from Hawk!")
```

Example:

```go
router.Get("/health", func(c *hawk.Context) {
    c.String(200, "OK")
})
```

This is useful for:

* Health checks
* Simple responses
* Plain-text endpoints
* Debugging
* Small responses that don't require JSON

---

# HTML Responses

Use `HTML()` when returning HTML content.

```go
c.HTML(200, "<h1>Hello Hawk</h1>")
```

Example:

```go
router.Get("/", func(c *hawk.Context) {

    c.HTML(200, `
        <html>
            <body>
                <h1>Welcome to Hawk</h1>
            </body>
        </html>
    `)
})
```

HTML responses are useful when Hawk is used to serve web pages rather than only APIs.

---

# Response Headers

The underlying response writer is available through:

```go
c.Response
```

You can set response headers using the standard Go API:

```go
c.Response.Header().Set(
    "X-Powered-By",
    "Hawk",
)
```

For example:

```go
router.Get("/users", func(c *hawk.Context) {

    c.Response.Header().Set(
        "X-API-Version",
        "v1",
    )

    c.JSON(200, users)
})
```

---

# Content Type

The response's content type should describe the data being returned.

JSON responses generally use:

```text
application/json
```

HTML responses generally use:

```text
text/html
```

Plain text responses generally use:

```text
text/plain
```

Hawk's response helpers should handle the appropriate content type for their respective response formats.

---

# Setting Custom Headers

Custom headers can be useful for metadata.

```go
c.Response.Header().Set(
    "X-Request-ID",
    requestID,
)
```

Multiple headers can be set:

```go
c.Response.Header().Set(
    "X-API-Version",
    "v1",
)

c.Response.Header().Set(
    "X-Request-ID",
    requestID,
)
```

---

# Cookies

Cookies are sent as part of the HTTP response.

Go's standard library can be used with Hawk:

```go
http.SetCookie(
    c.Response,
    &http.Cookie{
        Name:  "session",
        Value: "abc123",
        HttpOnly: true,
    },
)
```

Example:

```go
router.Post("/login", func(c *hawk.Context) {

    http.SetCookie(
        c.Response,
        &http.Cookie{
            Name:     "session",
            Value:    "abc123",
            HttpOnly: true,
        },
    )

    c.JSON(200, map[string]any{
        "message": "Logged in",
    })
})
```

The browser or HTTP client will receive the cookie through the `Set-Cookie` response header.

---

# Redirects

Because Hawk exposes the underlying `http.ResponseWriter`, Go's standard redirect functionality can be used.

```go
http.Redirect(
    c.Response,
    c.Request,
    "/login",
    http.StatusFound,
)
```

For example:

```go
router.Get("/dashboard", func(c *hawk.Context) {

    authenticated := false

    if !authenticated {
        http.Redirect(
            c.Response,
            c.Request,
            "/login",
            http.StatusFound,
        )
        return
    }

    c.HTML(200, "<h1>Dashboard</h1>")
})
```

---

# Empty Responses

Some endpoints don't need to return a response body.

For example, a successful deletion can use:

```text
204 No Content
```

With Go's standard response writer:

```go
c.Response.WriteHeader(http.StatusNoContent)
```

Example:

```go
router.Delete("/users/:id", func(c *hawk.Context) {

    // Delete user...

    c.Response.WriteHeader(
        http.StatusNoContent,
    )
})
```

A `204` response should not contain a response body.

---

# Response Lifecycle

HTTP status codes should generally be written before the response body.

Conceptually:

```text
Handler
   ↓
Set headers
   ↓
Set status code
   ↓
Write response body
```

For example:

```go
c.Response.Header().Set(
    "X-App",
    "Hawk",
)

c.JSON(200, data)
```

The response helper is responsible for writing the response correctly.

---

# Don't Write Multiple Responses

A handler should normally send one final response.

Avoid:

```go
c.JSON(400, map[string]any{
    "error": "Invalid request",
})

c.JSON(200, data)
```

Instead, return immediately after sending an error:

```go
if err != nil {
    c.JSON(400, map[string]any{
        "error": "Invalid request",
    })
    return
}

c.JSON(200, data)
```

This makes the response flow predictable.

---

# Response After Errors

A common handler pattern is:

```go
router.Get("/users/:id", func(c *hawk.Context) {

    id := c.Param("id")

    user, err := findUser(id)

    if err != nil {
        c.JSON(404, map[string]any{
            "error": "User not found",
        })
        return
    }

    c.JSON(200, user)
})
```

The flow is:

```text
Request
   ↓
Find resource
   ↓
Error?
 ┌─┴──────────────┐
 │                │
Yes              No
 │                │
404              200
 │                │
Stop             Data
```

---

# Response Middleware

Middleware can modify response headers before or around the handler.

For example:

```go
func SecurityHeaders(c *hawk.Context) {

    c.Response.Header().Set(
        "X-Content-Type-Options",
        "nosniff",
    )

    c.Next()
}
```

Register it globally:

```go
router.Use(SecurityHeaders)
```

The middleware can therefore apply common response configuration across routes.

---

# Testing Responses

Hawk responses can be tested using Go's `httptest` package.

Example:

```go
func TestJSONResponse(t *testing.T) {

    router := hawk.Default()

    router.Get("/users", func(c *hawk.Context) {

        c.JSON(200, map[string]any{
            "message": "Users",
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
}
```

You can also inspect the response body:

```go
body := rec.Body.String()
```

And inspect response headers:

```go
contentType := rec.Header().Get(
    "Content-Type",
)
```

---

# Testing String Responses

```go
func TestStringResponse(t *testing.T) {

    router := hawk.Default()

    router.Get("/hello", func(c *hawk.Context) {
        c.String(200, "Hello Hawk")
    })

    req := httptest.NewRequest(
        http.MethodGet,
        "/hello",
        nil,
    )

    rec := httptest.NewRecorder()

    router.ServeHTTP(rec, req)

    if rec.Code != http.StatusOK {
        t.Fatalf(
            "expected 200, got %d",
            rec.Code,
        )
    }

    if rec.Body.String() != "Hello Hawk" {
        t.Fatalf(
            "unexpected response: %s",
            rec.Body.String(),
        )
    }
}
```

---

# Response API

The main Hawk response methods are:

| Method     | Purpose           |
| ---------- | ----------------- |
| `JSON()`   | Return JSON data  |
| `String()` | Return plain text |
| `HTML()`   | Return HTML       |

The Context also exposes:

```go
c.Response
```

which allows access to Go's standard `http.ResponseWriter`.

This means applications can use additional HTTP functionality such as:

```go
http.SetCookie()
http.Redirect()
c.Response.Header().Set()
c.Response.WriteHeader()
```

---

# Complete Example

The following example demonstrates several response features:

```go
router.Get("/users/:id", func(c *hawk.Context) {

    id := c.Param("id")

    if id == "" {
        c.JSON(400, map[string]any{
            "error": "User ID is required",
        })
        return
    }

    c.Response.Header().Set(
        "X-API-Version",
        "v1",
    )

    c.JSON(200, map[string]any{
        "id": id,
    })
})
```

The handler:

1. Reads a route parameter.
2. Validates the value.
3. Sets a response header.
4. Returns JSON.
5. Uses an appropriate HTTP status code.

---

# Best Practices

### Use the correct status code

Don't return `200` for every operation.

For example:

```text
GET successful       → 200
Resource created     → 201
No response body     → 204
Invalid request      → 400
Unauthorized         → 401
Forbidden            → 403
Not found            → 404
Validation failure   → 422
Server error         → 500
```

### Return after an error response

Use:

```go
c.JSON(400, errorResponse)
return
```

This prevents the handler from attempting to send another response.

### Keep response formats consistent

If your API returns JSON errors, use a consistent structure across endpoints.

For example:

```json
{
    "error": "User not found"
}
```

### Use Hawk helpers for common responses

Prefer:

```go
c.JSON(200, data)
```

over manually encoding JSON for ordinary API responses.

Use the underlying `ResponseWriter` when you need functionality that isn't provided by Hawk's convenience methods.

---

# Summary

Hawk provides three main response helpers:

```go
c.JSON()
c.String()
c.HTML()
```

and exposes the underlying response writer:

```go
c.Response
```

This gives Hawk applications a simple API for common responses while retaining the flexibility of Go's `net/http` package.

The general response flow is:

```text
Handler
   ↓
Choose status code
   ↓
Set headers if necessary
   ↓
Write response
   ↓
Client
```

Understanding the response system together with `Context` and `Request` gives you the foundation needed to build complete HTTP APIs with Hawk.
