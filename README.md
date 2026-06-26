# Hawk

A lightweight Laravel-inspired web framework written in Go.

Hawk provides a familiar MVC structure, routing, middleware, validation, migrations, and a powerful CLI for building modern web applications in Go.

---

## Installation

```bash
go install github.com/beyond3800/hawk/cmd/hawk@latest
```

Create a new project:

```bash
hawk new blog
cd blog
hawk serve
```

---

## Features

* HTTP Routing
* Route Groups
* Middleware
* Request Validation
* Query Builder
* Database Migrations
* MVC Structure
* CLI Generators
* Redis Support
* Job System
* Auto Reload Development Server

---

## Quick Start

```go
app := hawk.New() // without recovery and logger

// or

app := hawk.Default() // with recovery and logger

app.Get("/", func(c *hawk.Context) {
    c.JSON(200, map[string]string{
        "message": "Hello Hawk",
    })
})
```

---

## Route Groups

```go
api := app.Group("/api")

api.Get("/users", func(c *hawk.Context) {
    c.JSON(200, users)
})
```

---

## Validation

```go
type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
}

if err := c.BindAndValidate(&user); err != nil {
    return
}
```

---

## Commands

```bash
hawk new blog
hawk serve
hawk migrate
hawk rollback
hawk status
```

---

## Documentation

Documentation is available inside the `docs` directory.

* Installation
* Routing
* Middleware
* Validation
* Database
* Migrations

---

## Contributing

Contributions, bug reports, and feature requests are welcome.

---

## License

MIT License.
