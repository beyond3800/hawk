# Installation

This guide will help you install Hawk and create your first application.

---

## Requirements

Before installing Hawk, make sure you have:

* Go 1.24 or later
* MySQL (optional)
* Redis (optional)

Verify your Go installation:

```bash
go version
```

---

## Install Hawk CLI

Install the Hawk command-line tool:

```bash
go install github.com/beyond3800/hawk/cmd/hawk@latest
```

Verify the installation:

```bash
hawk --help
```

---

## Create a New Project

Generate a new Hawk application:

```bash
hawk new blog
```

This creates a new project:

```text
blog/
├── app/
├── config/
├── database/
├── routes/
├── jobs/
├── .env
├── .air.toml
├── go.mod
└── main.go
```

---

## Enter the Project Directory

```bash
cd blog
```

---

## Configure Environment Variables

Edit your `.env` file.

### Database

```env
DB_ENABLED=true

DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=blog
DB_USERNAME=root
DB_PASSWORD=
```

### Redis

```env
REDIS_ENABLED=true

REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_PASSWORD=
```

If you do not want to use Redis or the database, disable them:

```env
DB_ENABLED=false
REDIS_ENABLED=false
```

---

## Start the Development Server

```bash
hawk serve
```

The application will start on:

```text
http://127.0.0.1:8000
```

You may specify a custom port:

```bash
hawk serve 9000
```

The application will then run on:

```text
http://127.0.0.1:9000
```

---

## Create Your First Route

Open `routes/web.go`:

```go
package routes

import "github.com/beyond3800/hawk"

func SetupRoutes() *hawk.Hawk {
    app := hawk.Default()

    app.Get("/", func(c *hawk.Context) {
        c.String(200, "Welcome to Hawk")
    })

    return app
}
```

Visit:

```text
http://127.0.0.1:8000
```

You should see:

```text
Welcome to Hawk
```

---

## Next Steps

* Routing
* Middleware
* Validation
* Query Builder
* Migrations
* Redis
* Jobs
* CLI Commands

Continue reading the documentation inside the `docs` directory.
