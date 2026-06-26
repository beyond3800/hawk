# Middleware

Middleware allows you to intercept requests before they reach your route handlers.

---

## Global Middleware

Global middleware runs on every request.

```go
app.Use(Logger())
app.Use(Recovery())
```

---

## Group Middleware

Middleware can also be attached to a route group.

```go
api := app.Group("/api")

api.Use(Auth())

api.Get("/profile", ProfileController.Show)
```

The `Auth` middleware will run before the `/profile` route.

---

## Creating Middleware

A middleware returns a `hawk.HandlerFunc`.

```go
func Auth() hawk.HandlerFunc {
    return func(c *hawk.Context) {

        token := c.Request.Header.Get("Authorization")

        if token == "" {
            c.JSON(401, map[string]string{
                "error": "Unauthorized",
            })

            c.Abort()
            return
        }

        c.Next()
    }
}
```

---

## Next()

`Next()` executes the next middleware or the route handler.

```go
c.Next()
```

If `Next()` is not called, request execution stops.

---

## Abort()

`Abort()` immediately stops the middleware chain.

```go
c.Abort()
```

Example:

```go
if token == "" {
    c.JSON(401, map[string]string{
        "error": "Unauthorized",
    })

    c.Abort()
    return
}
```

The route handler will never execute.

---

## Execution Order

```go
app.Use(Logger())
app.Use(Auth())

app.Get("/profile", handler)
```

Execution order:

1. Logger middleware
2. Auth middleware
3. Route handler

---

## Built-in Middleware

```go
app.Use(hawk.Logger())
app.Use(hawk.Recovery())
```

* `Logger()` logs incoming requests.
* `Recovery()` recovers from panics and prevents server crashes.
