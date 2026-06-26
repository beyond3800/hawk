# Project Structure

A typical Hawk application looks like this:

```text
app/
    Http/
        Controllers/
        Middleware/
        Services/

config/

database/
    migrations/

jobs/

routes/

.env
go.mod
main.go
```

---

## app/

Contains your application code.

### Controllers

```text
app/Http/Controllers/
```

Handles HTTP requests.

### Middleware

```text
app/Http/Middleware/
```

Processes requests before they reach routes.

### Services

```text
app/Http/Services/
```

Contains business logic.

---

## config/

Stores application configuration.

---

## database/

Contains migrations.

---

## jobs/

Contains background jobs.

---

## routes/

Defines application routes.

```text
routes/web.go
```

---

## .env

Stores environment variables.

---

## main.go

Application entry point.
