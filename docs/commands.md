# CLI Commands

The Hawk CLI helps you generate files, run migrations, and manage your application.

## Create a Controller

Generate a new controller:

```bash
hawk make controller UserController
```

Creates:

```text
app/Http/Controllers/UserController.go
```

---

## Create a Model

```bash
hawk make model User
```

Creates:

```text
app/Models/User.go
```

---

## Create a Service

```bash
hawk make service UserService
```

Creates:

```text
app/Http/Services/UserService.go
```

---

## Create Middleware

```bash
hawk make middleware Auth
```

Creates:

```text
app/Http/Middleware/Auth.go
```

---

## Create a Migration

```bash
hawk make migration create_users_table
```

Creates:

```text
database/migrations/
```

---

## Run Migrations

```bash
hawk migrate
```

Executes all pending migrations.

---

## Roll Back Migrations

```bash
hawk rollback
```

Reverts the last migration batch.

---

## Migration Status

```bash
hawk status
```

Displays the status of all migrations.

---

## Start Development Server

```bash
hawk serve
```

Starts the Hawk development server.

Specify a custom port:

```bash
hawk serve 9000
```

---

## Create a New Project

```bash
hawk new blog
```

Creates a new Hawk application.
