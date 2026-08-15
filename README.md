# Hawk

**Hawk is a lightweight web framework for Go, designed to make building backend applications simple, structured, and productive.**

Hawk provides the core tools you need to build HTTP applications and APIs in Go, including routing, middleware, request/response handling, validation, database queries, migrations, seeders, factories, transactions, and a CLI.

## Features

* 🚀 HTTP routing
* 🔗 Route parameters
* 📦 Route groups
* 🛡️ Middleware
* 🌐 Request and response handling
* 🔍 Query parameters
* 🍪 Cookies and headers
* 🧩 JSON, HTML, and string responses
* 🗄️ Database query builder
* 🔄 Database transactions
* 🏗️ Database migrations
* 🌱 Database seeders
* 🏭 Model factories
* ✅ Request validation
* 🖥️ Hawk CLI
* 🧪 Testing support
* ⚙️ CI-ready development workflow

---

## Installation

Install the Hawk CLI with Go:

```bash
go install github.com/beyond3800/hawk/cmd/hawk@latest
```

Verify the installation:

```bash
hawk --help
```

---

## Creating a New Application

Create a new Hawk application:

```bash
hawk new blog
```

Move into the project:

```bash
cd blog
```

Start the development server:

```bash
hawk serve
```

Your application is now ready for development.

---

## Routing

Hawk provides simple HTTP routing.

```go
router.Get("/users", func(c *hawk.Context) {
    c.JSON(200, map[string]any{
        "message": "Users",
    })
})
```

### Route Parameters

Parameters can be defined using `:`:

```go
router.Get("/users/:id", func(c *hawk.Context) {
    id := c.Param("id")

    c.JSON(200, map[string]any{
        "id": id,
    })
})
```

A request to:

```text
/users/123
```

makes `c.Param("id")` return:

```text
123
```

### Multiple Parameters

```go
router.Get(
    "/users/:userID/posts/:postID",
    func(c *hawk.Context) {
        userID := c.Param("userID")
        postID := c.Param("postID")

        // ...
    },
)
```

### HTTP Methods

Hawk supports the common HTTP methods:

```go
router.Get(...)
router.Post(...)
router.Put(...)
router.Delete(...)
```

---

## Route Groups

Routes can be organized using groups.

```go
api := router.Group("/api")

api.Get("/users", users)
api.Get("/posts", posts)
```

The resulting routes are:

```text
/api/users
/api/posts
```

---

## Middleware

Middleware can be registered to run before route handlers.

```go
router.Use(func(c *hawk.Context) {
    // middleware logic
    c.Next()
})
```

Middleware can also stop request execution using:

```go
c.Abort()
```

For example:

```go
router.Use(func(c *hawk.Context) {
    if !authenticated {
        c.Abort()
        return
    }

    c.Next()
})
```

---

## Request Data

Hawk provides access to common HTTP request data.

### Query Parameters

```text
/users?page=2
```

```go
page := c.Query("page")
```

### Headers

```go
token := c.Request.Header.Get("Authorization")
```

### Cookies

```go
cookie, err := c.Request.Cookie("session")
```

### Request Body

Hawk works with Go's standard HTTP request body:

```go
body := c.Request.Body
```

---

## Responses

### JSON

```go
c.JSON(200, map[string]any{
    "message": "Hello Hawk",
})
```

### String

```go
c.String(200, "Hello Hawk")
```

### HTML

```go
c.HTML(200, "<h1>Hello Hawk</h1>")
```

---

## Validation

Hawk provides built-in validation rules.

Current rules include:

* `required`
* `email`
* `min`
* `max`

Example:

```go
type User struct {
    Name     string `validate:"required|min:3|max:50"`
    Email    string `validate:"required|email"`
    Password string `validate:"required|min:8"`
}
```

Validation errors are available through Hawk's validation error structure.

See the validation documentation for complete usage.

---

## Database

Hawk includes a database query builder for common database operations.

### Create

```go
result, err := HawkDB().
    Table("users").
    Create(map[string]any{
        "name":  "Adam",
        "email": "adam@test.com",
    })
```

### Query

```go
var users []User

err := HawkDB().
    Table("users").
    Get(&users)
```

### First

```go
var user User

err := HawkDB().
    Table("users").
    First(&user)
```

### Pagination

```go
var users []User

result, err := HawkDB().
    Table("users").
    Paginate(1, 10, &users)
```

Pagination metadata is available through the returned result:

```go
result.Meta.Total
result.Meta.PerPage
```

---

## Transactions

Hawk provides database transactions.

```go
err := HawkDB().Transaction(func(tx *Builder) error {

    _, err := tx.Table("users").Create(map[string]any{
        "name":  "Adam",
        "email": "adam@test.com",
    })

    if err != nil {
        return err
    }

    _, err = tx.Table("wallets").Create(map[string]any{
        "user_id": 1,
        "balance": 0,
    })

    if err != nil {
        return err
    }

    return nil
})
```

Returning an error from the transaction callback causes the transaction to roll back.

---

## Migrations

Hawk provides a migration system for managing database schema changes.

Create a migration:

```bash
hawk make:migration create_users_table
```

Run migrations:

```bash
hawk migrate
```

Rollback migrations using the rollback command provided by your Hawk installation.

---

## Database Seeding

Hawk applications can define database seed data.

Run the application's seed command:

```bash
hawk db:seed
```

Seeders are useful for creating initial application data and test data.

---

## Factories

Hawk provides factories for generating test/development data.

For example, create a factory:

```bash
hawk make:factory User
```

Factories can use Hawk's built-in faker functionality to generate realistic test data.

If the built-in faker functionality doesn't meet your application's requirements, you can use another faker package.

---

## CLI

Hawk provides a CLI for common development tasks.

Examples:

```bash
hawk new blog
hawk serve
hawk migrate
hawk db:seed
hawk make:migration create_users_table
hawk make:factory User
```

Run:

```bash
hawk --help
```

to see the commands available in your installed version.

---

## Testing

Hawk projects can use Go's standard testing tools.

Run all tests:

```bash
go test ./...
```

Run static analysis:

```bash
go vet ./...
```

Build all packages:

```bash
go build ./...
```

Hawk itself uses automated tests for important framework components including routing, middleware, validation, database operations, transactions, and pagination.

---

## Project Documentation

Detailed documentation is available in the `docs/` directory:

* [Getting Started](docs/getting-started.md)
* [Routing](docs/routing.md)
* [Middleware](docs/middleware.md)
* [Context](docs/context.md)
* [Request Handling](docs/request.md)
* [Responses](docs/responses.md)
* [Validation](docs/validation.md)
* [Database](docs/database.md)
* [Query Builder](docs/query-builder.md)
* [Migrations](docs/migrations.md)
* [Seeders](docs/seeding.md)
* [Factories](docs/factories.md)
* [Transactions](docs/transactions.md)
* [CLI](docs/cli.md)
* [Testing](docs/testing.md)
* [Deployment](docs/deployment.md)

---

## Contributing

Contributions, bug reports, documentation improvements, and feature suggestions are welcome.

Before submitting changes, make sure the project passes:

```bash
go test ./...
go vet ./...
go build ./...
```

---

## License

Hawk is released under the license specified in the repository's `LICENSE` file.

---

## Status

Hawk is being developed toward its `v1.0.0` release.

The API and features may continue to evolve before the stable release.
