
---

# `README.md`


```md
# Hawk

A lightweight Laravel-inspired web framework written in Go.

## Features

- Fast HTTP Router
- Middleware
- Route Groups
- Validation
- Query Builder
- Migrations
- CLI Generator
- MVC Structure

## Quick Start

```go
app := hawk.New() -> without recovery or logger
app := hawk.Default() -> with recovery or logger

app.Get("/", func(c *hawk.Context) {
    c.JSON(200, map[string]string{
        "message": "Hello Hawk",
    })
})

app.Run(":8000")