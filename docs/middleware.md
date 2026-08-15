# Middleware

Middleware allows you to execute code during the processing of an HTTP request.

Middleware is useful for functionality that should run before or around route handlers, such as:

* Authentication
* Authorization
* Logging
* CORS
* Request validation
* Error handling
* Request processing

A Hawk middleware receives a `*hawk.Context`.

```go
func Middleware(c *hawk.Context) {
    // middleware logic
    c.Next()
}
```

---

## Middleware Lifecycle

A request can pass through multiple middleware functions before reaching the route handler.

A simplified lifecycle looks like:

```text
Request
   ↓
Global Middleware
   ↓
Route Middleware
   ↓
Route Handler
   ↓
Response
```

Calling `c.Next()` tells Hawk to continue executing the next handler in the current handler chain.

Calling `c.Abort()` stops the remaining handlers from executing.

---

## Registering Middleware

Global middleware can be registered using `Use()`.

```go
router.Use(func(c *hawk.Context) {
    // middleware logic

    c.Next()
})
```

Once registered, the middleware participates in requests handled by that router.

For example:

```go
router := hawk.Default()

router.Use(func(c *hawk.Context) {
    fmt.Println("Request received")

    c.Next()
})

router.Get("/users", func(c *hawk.Context) {
    c.JSON(200, map[string]any{
        "message": "Users",
    })
})
```

A request to `/users` passes through the middleware before reaching the route handler.

---

## `Next()`

`Next()` continues execution through the handler chain.

For example:

```go
router.Use(func(c *hawk.Context) {

    fmt.Println("Before handler")

    c.Next()

    fmt.Println("After handler")
})
```

The middleware can therefore perform work both before and after the remaining handlers execute.

The general flow is:

```text
Middleware starts
       ↓
Before handler
       ↓
c.Next()
       ↓
Next handler
       ↓
After handler
       ↓
Middleware finishes
```

This makes middleware useful for tasks such as measuring request duration.

For example:

```go
func Logger(c *hawk.Context) {

    start := time.Now()

    c.Next()

    duration := time.Since(start)

    fmt.Println("Request took:", duration)
}
```

---

## `Abort()`

`Abort()` stops the current handler chain.

This is useful when a request should not continue.

For example, an authentication middleware can stop unauthenticated requests:

```go
func Auth(c *hawk.Context) {

    authenticated := false

    if !authenticated {
        c.Abort()
        return
    }

    c.Next()
}
```

The important part is to return after calling `Abort()`:

```go
c.Abort()
return
```

This prevents the middleware itself from continuing its own logic.

---

## Authentication Example

A common middleware use case is authentication.

```go
func Auth(c *hawk.Context) {

    token := c.Request.Header.Get("Authorization")

    if token == "" {
        c.Abort()
        return
    }

    c.Next()
}
```

The middleware checks the request before allowing the request to continue.

The flow becomes:

```text
Request
   ↓
Auth Middleware
   ↓
Token exists?
   ├── No  → Abort
   │
   └── Yes
        ↓
      Next()
        ↓
     Handler
```

---

## Middleware Execution Order

Middleware executes according to the order in which it is registered.

For example:

```go
router.Use(MiddlewareA)
router.Use(MiddlewareB)
router.Use(MiddlewareC)
```

The request enters the middleware chain in this order:

```text
MiddlewareA
    ↓
MiddlewareB
    ↓
MiddlewareC
    ↓
Route Handler
```

When middleware calls `Next()`, execution continues to the next handler.

---

## Multiple Middleware

Multiple middleware functions can be registered:

```go
router.Use(func(c *hawk.Context) {

    fmt.Println("Logger")

    c.Next()
})

router.Use(func(c *hawk.Context) {

    fmt.Println("Authentication")

    c.Next()
})
```

The resulting chain is:

```text
Logger
   ↓
Authentication
   ↓
Route Handler
```

Each middleware should call `Next()` when it wants the request to continue.

---

## Middleware That Stops Execution

Consider:

```go
router.Use(func(c *hawk.Context) {

    if !authenticated {
        c.Abort()
        return
    }

    c.Next()
})
```

If the request is not authenticated:

```text
Request
   ↓
Middleware
   ↓
Abort()
   ↓
Route Handler is NOT executed
```

This behavior is important for authentication and authorization.

---

## Route Middleware

Middleware can also be applied to specific routes when supported by the route registration API.

For example:

```go
router.Get(
    "/admin",
    AuthMiddleware,
    AdminHandler,
)
```

This allows authentication to apply specifically to the `/admin` route.

Other routes can remain unaffected:

```go
router.Get("/users", UsersHandler)

router.Get(
    "/admin",
    AuthMiddleware,
    AdminHandler,
)
```

The resulting behavior is:

```text
/users
   ↓
UsersHandler


/admin
   ↓
AuthMiddleware
   ↓
AdminHandler
```

Use the route registration API provided by your installed Hawk version when registering route-specific middleware.

---

## CORS Middleware

Middleware is also useful for CORS.

For example:

```go
func CORS(c *hawk.Context) {

    c.Response.Header().Set(
        "Access-Control-Allow-Origin",
        "*",
    )

    c.Next()
}
```

Register it globally:

```go
router.Use(CORS)
```

Every request handled by the router can then receive the configured CORS headers.

For production applications, configure CORS according to the application's security requirements rather than automatically allowing every origin.

---

## Logging Middleware

Logging middleware can inspect every request.

```go
func Logger(c *hawk.Context) {

    start := time.Now()

    c.Next()

    duration := time.Since(start)

    fmt.Printf(
        "%s %s %v\n",
        c.Request.Method,
        c.Request.URL.Path,
        duration,
    )
}
```

Register it:

```go
router.Use(Logger)
```

This provides a central place for request logging.

---

## Middleware and Response Handling

Middleware can also inspect or modify the response before or after calling `Next()`.

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

This can be useful for applying headers consistently across an application.

---

## Middleware Chain

Conceptually, Hawk executes middleware and route handlers as a chain:

```text
┌─────────────────┐
│     Request     │
└────────┬────────┘
         ↓
┌─────────────────┐
│   Middleware A  │
└────────┬────────┘
         ↓ Next()
┌─────────────────┐
│   Middleware B  │
└────────┬────────┘
         ↓ Next()
┌─────────────────┐
│   Route Handler │
└────────┬────────┘
         ↓
┌─────────────────┐
│    Response     │
└─────────────────┘
```

If a middleware calls `Abort()`:

```text
┌─────────────────┐
│     Request     │
└────────┬────────┘
         ↓
┌─────────────────┐
│   Middleware A  │
└────────┬────────┘
         ↓
      Abort()
         ↓
       STOP
```

The remaining handlers aren't executed.

---

## Middleware Best Practices

### Keep middleware focused

A middleware should generally have one responsibility.

Good:

```go
AuthMiddleware
LoggerMiddleware
CORSMiddleware
```

Avoid creating one huge middleware that handles authentication, logging, validation, database operations, and unrelated application logic.

### Call `Next()` when appropriate

If the middleware allows the request to continue:

```go
c.Next()
```

### Return after `Abort()`

Use:

```go
if unauthorized {
    c.Abort()
    return
}
```

rather than continuing execution after aborting.

### Avoid unnecessary global middleware

Global middleware runs across requests handled by the router. If functionality is only required for a small group of routes, consider applying it at the appropriate route or group level.

---

## Testing Middleware

Middleware behavior should be tested using Go's `httptest` package.

For example, an aborting middleware can be tested by checking that the request is stopped:

```go
func TestAbortMiddleware(t *testing.T) {

    router := hawk.Default()

    authenticated := false

    router.Use(func(c *hawk.Context) {

        if !authenticated {
            c.Abort()
            return
        }

        c.Next()
    })

    router.Get("/users", func(c *hawk.Context) {
        c.String(200, "users")
    })

    req := httptest.NewRequest(
        http.MethodGet,
        "/users",
        nil,
    )

    rec := httptest.NewRecorder()

    router.ServeHTTP(rec, req)
}
```

Tests should verify the behavior Hawk promises for aborted requests and handler execution.

---

## Summary

Hawk middleware provides a way to intercept and control request processing.

The key methods are:

```text
Use()
Next()
Abort()
```

The basic rules are:

```text
Use()
  ↓
Middleware
  ↓
Next()
  ↓
Next Handler
```

or:

```text
Middleware
  ↓
Abort()
  ↓
Stop
```

Middleware is one of the main mechanisms for implementing cross-cutting functionality in Hawk applications.
