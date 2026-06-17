# Routing

## Basic Route

```go
app.Get("/", func(c *hawk.Context) {
    c.String(200, "Welcome to Hawk")
})

Param in usage
app.Get("/users/:id", func(c *hawk.Context) {
    id := c.Param("id")

    c.JSON(200, map[string]string{
        "id": id,
    })
})

// Query in usage
// GET /users?page=1
// page := c.Query("page")

app.Get("/users?page=num", func(c *hawk.Context) {
    page := c.Query("page")

    c.JSON(200, map[string]string{
        "page": page,
    })
})



// Group routing 
api := app.Group("/api")

api.Get("/users", UserController.Index)
api.Post("/users", UserController.Store)