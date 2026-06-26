# Routing

Routing allows you to map HTTP requests to handlers.

---

## Basic Routes

```go
app.Get("/", func(c *hawk.Context) {
    c.String(200, "Welcome to Hawk")
})

app.Post("/users", UserController.Store)

app.Put("/users/:id", UserController.Update)

app.Delete("/users/:id", UserController.Delete)
```

---

## Route Parameters

Route parameters allow you to capture values from the URL.

```go
app.Get("/users/:id", func(c *hawk.Context) {
    id := c.Param("id")

    c.JSON(200, map[string]string{
        "id": id,
    })
})
```

Request:

```text
GET /users/10
```

Response:

```json
{
    "id": "10"
}
```

---

## Query Parameters

Query parameters are obtained from the request URL.

```go
app.Get("/users", func(c *hawk.Context) {
    page := c.Query("page")

    c.JSON(200, map[string]string{
        "page": page,
    })
})
```

Request:

```text
GET /users?page=1
```

Response:

```json
{
    "page": "1"
}
```

---

## Route Groups

Route groups allow you to organize routes with a common prefix.

```go
api := app.Group("/api")

api.Get("/users", UserController.Index)
api.Post("/users", UserController.Store)
```

Generated routes:

```text
GET  /api/users
POST /api/users
```

---

## Middleware Groups

Middleware can be attached to route groups.

```go
api := app.Group("/api")

api.Use(Auth())

api.Get("/profile", UserController.Profile)
```

The `Auth` middleware runs before the route handler.

---

## Available Methods

```go
app.Get()
app.Post()
app.Put()
app.Delete()
```
