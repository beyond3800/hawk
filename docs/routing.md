# Routing

Hawk provides a simple HTTP router for registering application routes.

Routes connect an HTTP method and URL pattern to a handler function.

```go
router.Get("/users", func(c *hawk.Context) {
    c.JSON(200, map[string]any{
        "message": "Users",
    })
})
```

---

## HTTP Methods

Hawk supports the common HTTP methods used when building APIs.

### GET

Use `Get()` for retrieving resources.

```go
router.Get("/users", func(c *hawk.Context) {
    c.JSON(200, users)
})
```

### POST

Use `Post()` for creating resources.

```go
router.Post("/users", func(c *hawk.Context) {
    c.JSON(201, map[string]any{
        "message": "User created",
    })
})
```

### PUT

Use `Put()` for updating resources.

```go
router.Put("/users/:id", func(c *hawk.Context) {
    id := c.Param("id")

    c.JSON(200, map[string]any{
        "id": id,
    })
})
```

### DELETE

Use `Delete()` for deleting resources.

```go
router.Delete("/users/:id", func(c *hawk.Context) {
    id := c.Param("id")

    c.JSON(200, map[string]any{
        "id": id,
    })
})
```

---

## Static Routes

A static route contains a fixed URL path.

```go
router.Get("/users", func(c *hawk.Context) {
    c.String(200, "Users")
})
```

A request to:

```text
GET /users
```

matches the route.

However:

```text
GET /posts
```

does not match `/users`.

---

## Route Parameters

Route parameters allow a section of the URL to be dynamic.

Define a parameter using `:`:

```go
router.Get("/users/:id", func(c *hawk.Context) {
    id := c.Param("id")

    c.JSON(200, map[string]any{
        "id": id,
    })
})
```

The following request:

```text
GET /users/123
```

produces:

```text
id = 123
```

The parameter can be retrieved with:

```go
c.Param("id")
```

Route parameters are returned as strings.

---

## Multiple Parameters

A route can contain multiple parameters.

```go
router.Get(
    "/users/:userID/posts/:postID",
    func(c *hawk.Context) {

        userID := c.Param("userID")
        postID := c.Param("postID")

        c.JSON(200, map[string]any{
            "user_id": userID,
            "post_id": postID,
        })
    },
)
```

A request to:

```text
GET /users/123/posts/456
```

produces:

```text
userID = 123
postID = 456
```

---

## Query Parameters

Query parameters are different from route parameters.

Given:

```text
GET /users?page=2&search=adam
```

the path is:

```text
/users
```

while the query string is:

```text
?page=2&search=adam
```

Hawk provides the `Query()` method:

```go
page := c.Query("page")
search := c.Query("search")
```

Both values are returned as strings.

For example:

```go
router.Get("/users", func(c *hawk.Context) {

    page := c.Query("page")

    c.JSON(200, map[string]any{
        "page": page,
    })
})
```

---

## Route Groups

Route groups allow multiple routes to share a common path prefix.

For example:

```go
api := router.Group("/api")
```

Routes can then be registered through the group:

```go
api.Get("/users", getUsers)
api.Get("/posts", getPosts)
```

These routes become:

```text
/api/users
/api/posts
```

Groups are useful for organizing APIs.

For example:

```go
api := router.Group("/api")

v1 := api.Group("/v1")

v1.Get("/users", getUsers)
v1.Get("/posts", getPosts)
```

The resulting routes are:

```text
/api/v1/users
/api/v1/posts
```

---

## Middleware on Routes

Hawk allows middleware to participate in request handling.

A middleware function receives the current context:

```go
func AuthMiddleware(c *hawk.Context) {
    // authentication logic

    c.Next()
}
```

Middleware can be registered with the router:

```go
router.Use(AuthMiddleware)
```

A middleware can stop the request by calling:

```go
c.Abort()
```

For example:

```go
func AuthMiddleware(c *hawk.Context) {

    authenticated := checkAuthentication(c)

    if !authenticated {
        c.Abort()
        return
    }

    c.Next()
}
```

When `Abort()` is called, subsequent handlers are not executed.

---

## Route Middleware

Middleware can also be associated with a specific route when registering its handlers.

For example, conceptually:

```go
router.Get(
    "/admin",
    AuthMiddleware,
    AdminController,
)
```

This allows authentication or other middleware to apply only to specific routes.

Use the API provided by your Hawk version when registering route-specific handlers.

---

## Route Matching

Hawk matches a request against registered routes using:

1. HTTP method
2. URL path
3. Static path segments
4. Route parameters

For example:

```text
/users/:id
```

can match:

```text
/users/123
/users/456
/users/abc
```

but it does not match:

```text
/posts/123
```

---

## Static Routes and Parameter Routes

When a static route and a parameter route could both match a request, the static route should take priority.

For example:

```go
router.Get("/users/me", func(c *hawk.Context) {
    c.String(200, "static")
})

router.Get("/users/:id", func(c *hawk.Context) {
    c.String(200, "parameter")
})
```

For:

```text
GET /users/me
```

Hawk should execute:

```text
static
```

rather than treating `me` as an `id`.

This behavior makes routes such as:

```text
/users/me
/users/profile
/users/settings
```

safe to define alongside:

```text
/users/:id
```

---

## HTTP Method Matching

Routes are matched against the HTTP method as well as the path.

For example:

```go
router.Get("/users", getUsers)
router.Post("/users", createUser)
```

A:

```text
GET /users
```

request matches the GET route.

A:

```text
POST /users
```

request matches the POST route.

A:

```text
DELETE /users
```

request does not match either route unless a DELETE route has been registered.

---

## 404 — Route Not Found

When no registered route matches the request, Hawk returns an HTTP 404 response.

For example, if only this route exists:

```go
router.Get("/users", getUsers)
```

then:

```text
GET /unknown
```

results in:

```text
404 Not Found
```

A 404 can also occur when the path exists but no route has been registered for the requested HTTP method.

---

## Route Handler

A Hawk route handler receives a `*hawk.Context`.

```go
func getUsers(c *hawk.Context) {
    c.JSON(200, map[string]any{
        "message": "Users",
    })
}
```

The handler can use the context to access request information and produce a response.

---

## Complete Example

A small API can be organized like this:

```go
router := hawk.Default()

router.Get("/users", func(c *hawk.Context) {
    c.JSON(200, map[string]any{
        "message": "List users",
    })
})

router.Get("/users/:id", func(c *hawk.Context) {
    id := c.Param("id")

    c.JSON(200, map[string]any{
        "id": id,
    })
})

router.Post("/users", func(c *hawk.Context) {
    c.JSON(201, map[string]any{
        "message": "User created",
    })
})

router.Put("/users/:id", func(c *hawk.Context) {
    id := c.Param("id")

    c.JSON(200, map[string]any{
        "id": id,
    })
})

router.Delete("/users/:id", func(c *hawk.Context) {
    id := c.Param("id")

    c.JSON(200, map[string]any{
        "id": id,
    })
})
```

This provides:

```text
GET     /users
GET     /users/:id
POST    /users
PUT     /users/:id
DELETE  /users/:id
```

---

## Routing Best Practices

### Keep routes organized

Group related routes together:

```go
api := router.Group("/api")

api.Get("/users", getUsers)
api.Post("/users", createUser)
```

### Use meaningful parameter names

Prefer:

```text
/users/:userID/posts/:postID
```

over:

```text
/users/:x/posts/:y
```

Meaningful names make handlers easier to understand.

### Put specific routes before generic parameter routes

For example:

```text
/users/me
/users/:id
```

The static route `/users/me` should take precedence when both can match.

### Keep business logic outside route definitions

For larger applications, use controllers or other application-layer functions rather than putting large amounts of logic directly inside route registration.

---

## Summary

Hawk routing provides:

* HTTP method routing
* Static routes
* Route parameters
* Multiple route parameters
* Query parameters
* Route groups
* Middleware
* Static-route priority
* 404 handling

The router is designed to provide a simple foundation for building APIs and web applications in Go.
