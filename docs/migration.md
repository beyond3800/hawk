# Migrations

Hawk provides a migration system for managing database schema changes in a structured and repeatable way.

Migrations allow you to:

* Create database tables
* Add or remove columns
* Modify database structure
* Roll back schema changes
* Track executed migrations
* Check migration status

Migrations are especially useful when working with a team or deploying an application to different environments.

---

# Migration Workflow

A typical Hawk migration workflow looks like this:

```text
Create Migration
      ↓
Write Schema Changes
      ↓
Run Migration
      ↓
Migration Recorded
      ↓
Check Status
      ↓
Rollback When Necessary
```

---

# Migration Directory

Hawk stores migration files inside the application's migrations directory.

A typical project may look like:

```text
project/
├── database/migrations/
│   ├── 20260808120000_create_users_table.go
│   ├── 20260808121000_create_posts_table.go
│   └── ...
├── routes/
├── models/
└── main.go
```

The exact migration filename depends on the migration generator used by the Hawk CLI.

---

# Creating a Migration

Use the Hawk CLI to generate a migration:

```bash
hawk make migration create_users_table
```

This creates a migration file that can be edited to define the required schema changes.

For example:

```text
migrations/
└── create_users_table.go
```

Migration generation keeps migration files consistent and avoids manually creating the migration structure.

---

# Migration Structure

A migration contains two main operations:

```text
Up
↓
Apply the schema change

Down
↓
Reverse the schema change
```

Conceptually:

```go
type Migration interface {
    Up() error
    Down() error
}
```

The exact migration implementation should follow the migration interface provided by the current Hawk version.

---

# Up Migration

The `Up` operation contains the changes that should be applied to the database.

For example:

```text
Create users table
      ↓
Add columns
      ↓
Create indexes
```

A migration might create a table containing:

```text
users
├── id
├── name
├── email
└── created_at
```

The `Up` operation is executed when the migration is applied.

---

# Down Migration

The `Down` operation reverses the changes made by `Up`.

For a migration that creates a `users` table:

```text
Up
↓
CREATE users table

Down
↓
DROP users table
```

The goal is for a migration to be reversible.

---

# Running Migrations

Use:

```bash
hawk migrate
```

to run pending migrations.

Hawk checks which migrations have already been executed and runs only migrations that have not yet been applied.

Conceptually:

```text
Migration files
      ↓
Check migration history
      ↓
Find pending migrations
      ↓
Run pending migrations
      ↓
Record successful migrations
```

---

# Migration Tracking

Hawk keeps track of migrations that have already been executed.

This prevents the same migration from being executed repeatedly.

For example:

```text
Migration A   ✓
Migration B   ✓
Migration C   ✗
```

Running:

```bash
hawk migrate
```

will apply migration C rather than running A and B again.

---

# Migration History

The migration tracking system stores information about executed migrations.

The migration system can use this information to determine:

```text
Which migrations have run?
Which migrations are pending?
Which migration should be rolled back?
```

This is important because migration files themselves do not indicate whether their changes have already been applied to a particular database.

---

# Migration Status

You can check migration status using:

```bash
hawk status
```

The status command provides information about migration execution.

Conceptually, the output can distinguish between:

```text
Ran
Pending
```

For example:

```text
Migration                         Status

create_users_table                Ran
create_posts_table                Ran
create_comments_table             Pending
```

This is useful when debugging database setup.

---

# Pending Migrations

A migration is considered pending when:

1. It exists in the migrations directory.
2. It has not been recorded as executed.

For example:

```text
Migration files:

A
B
C
D

Executed:

A
B

Pending:

C
D
```

Running:

```bash
hawk migrate
```

will apply C and D.

---

# Migration Order

Migrations should be executed in a predictable order.

For example:

```text
001_create_users
002_create_categories
003_create_posts
```

If `posts` depends on `users`, the users migration should run first.

This is particularly important when using foreign keys.

A typical dependency might look like:

```text
users
  ↓
categories
  ↓
posts
```

Therefore:

```text
create_users
      ↓
create_categories
      ↓
create_posts
```

---

# Foreign Keys

When creating relationships between tables, migrations can define foreign keys.

For example:

```text
users
  │
  │ user_id
  ↓
posts
```

The `posts.user_id` column references the users table.

The users table should therefore exist before the posts migration is executed.

---

# Rollback

Rollback reverses an applied migration.

Use:

```bash
hawk rollback
```

A rollback should execute the migration's reverse operation.

Conceptually:

```text
Migration
    ↓
Up()
    ↓
Database changed
    ↓
rollback
    ↓
Down()
    ↓
Database restored
```

---

# Example Rollback

Suppose the migration creates:

```text
users
```

Running:

```bash
hawk migrate
```

creates the table.

Running:

```bash
hawk rollback
```

executes the migration's `Down` operation and removes the schema change.

---

# Why Down Matters

A migration without a reliable reverse operation can make development and deployment more difficult.

For example:

```text
Up:
Create users table

Down:
Drop users table
```

This allows a developer to move backward during development.

A good migration should therefore define both directions clearly.

---

# Migration Failures

A migration should be considered successful only when its database operation completes successfully.

If a migration fails:

```text
Migration
    ↓
Execute Up()
    ↓
Database Error
    ↓
Migration fails
```

The error should be returned rather than silently ignored.

For example:

```go
if err != nil {
    return err
}
```

This allows the CLI to report the failure.

---

# Migration Registration

Hawk maintains a collection of registered migrations.

Conceptually:

```text
Migration files
      ↓
Migration registration
      ↓
Migration manager
      ↓
Database
```

The migration system needs to know which migrations are available before it can determine which ones are pending.

---

# Executed Migrations

The migration system can query the database to determine which migrations have already been executed.

Conceptually:

```go
executed := ExecutedMigrations()
```

The migration manager can then compare:

```text
Registered migrations
        VS
Executed migrations
```

to determine which migrations need to run.

---

# Migration State

Migration state can be thought of as:

```text
Registered
    │
    ├── Executed
    │
    └── Pending
```

For example:

```text
Registered:
    create_users
    create_posts
    create_comments

Executed:
    create_users
    create_posts

Pending:
    create_comments
```

---

# Creating a Users Table

A migration can be used to create the initial users table.

Conceptually, the schema might contain:

```text
users
├── id
├── name
├── email
└── password
```

The migration's `Up` operation creates the table.

The `Down` operation removes it.

---

# Creating Related Tables

Suppose Hawk is being used to build a blog API.

The schema might contain:

```text
users
   │
   └──── posts
           │
           └──── categories
```

The migrations can be separated:

```text
create_users_table
create_categories_table
create_posts_table
```

This makes each schema change easier to understand and manage.

---

# One Change Per Migration

Prefer migrations that describe a focused schema change.

For example:

```bash
hawk make migration create_users_table
```

followed by:

```bash
hawk make migration create_posts_table
```

is generally easier to maintain than one large migration responsible for creating the entire application's schema.

Focused migrations make it easier to:

* Understand schema history
* Debug failures
* Roll back changes
* Review changes in Git

---

# Migration Naming

Use descriptive migration names.

Good:

```text
create_users_table
create_posts_table
add_slug_to_posts
add_category_id_to_posts
create_comments_table
```

Avoid vague names such as:

```text
update_database
changes
fix_table
new_stuff
```

A migration name should describe the schema change it introduces.

---

# Migrations and Version Control

Migration files should normally be committed to Git.

For example:

```bash
git add migrations/
git commit -m "Add posts migration"
```

This allows the database schema to evolve alongside the application code.

When another developer checks out the project, they can run:

```bash
hawk migrate
```

to bring their database up to date.

---

# Migrations in Deployment

A deployment can follow this pattern:

```text
Deploy Application
       ↓
Run hawk migrate
       ↓
Apply Pending Migrations
       ↓
Start Application
```

This ensures that required database changes are applied before the application starts using them.

Migration execution should be handled carefully in production, especially when migrations modify existing data or large tables.

---

# Migrations and Testing

Tests can use a separate test database so that schema changes do not affect development or production data.

A typical test setup is:

```go
func setupTestDB(t *testing.T) {
    // Create test database
    // Configure HawkDB
    // Create required tables
}
```

Tests can then create the schema they need.

For example:

```go
func createWalletTable(t *testing.T) {

    _, err := HawkDB().Conn.Exec(`
        CREATE TABLE wallets(
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id TEXT,
            balance TEXT
        );
    `)

    if err != nil {
        t.Fatal(err)
    }
}
```

This keeps database tests isolated from real application data.

---

# Migration Best Practices

### Keep migrations small

Each migration should represent a clear schema change.

### Make migrations reversible

Whenever possible, make sure `Down` correctly reverses `Up`.

### Use descriptive names

Migration names should explain what changed.

### Commit migrations to Git

The migration history is part of the application's source code.

### Test migrations

Make sure migrations can be applied successfully in a clean database.

### Be careful with production data

Schema changes affecting existing data should be planned carefully.

### Preserve migration history

Avoid casually modifying migrations that have already been applied to shared or production databases.

Instead, create a new migration for subsequent changes.

---

# CLI Commands

The main migration-related commands are:

```bash
hawk make migration <name>
```

Creates a new migration.

```bash
hawk migrate
```

Runs pending migrations.

```bash
hawk rollback
```

Rolls back the applicable migration according to Hawk's migration rollback behavior.

```bash
hawk status
```

Displays migration status.

---

# Migration Lifecycle

The complete lifecycle is:

```text
hawk make:migration create_users_table
             │
             ▼
       Migration File
             │
             ▼
        Define Up()
        Define Down()
             │
             ▼
       hawk migrate
             │
             ▼
       Database Updated
             │
             ▼
    Migration Recorded
             │
             ▼
       hawk status
             │
             ├── Ran
             └── Pending
             │
             ▼
       hawk rollback
             │
             ▼
       Down() Executed
```

---

# Summary

Hawk migrations provide a structured way to manage database schema changes.

The core commands are:

```bash
hawk make migration create_users_table
hawk migrate
hawk status
hawk rollback
```

The fundamental migration model is:

```text
Up()
 ↓
Apply change

Down()
 ↓
Reverse change
```

Hawk tracks executed migrations so that pending migrations can be distinguished from migrations that have already been applied.

A well-managed migration system gives Hawk applications a predictable database lifecycle:

```text
Migration File
      ↓
Schema Change
      ↓
Migration Tracking
      ↓
Status
      ↓
Rollback
```

Migrations should be treated as part of the application's source code and maintained alongside the rest of the project.
