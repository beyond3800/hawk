# Context

The `Context` is the central object passed to Hawk route handlers and middleware.

It provides access to:

* Request data
* URL parameters
* Query parameters
* JSON request bodies
* Form data
* Uploaded files
* Cookies
* Response helpers
* Middleware control
* Validation errors

A typical handler receives the context like this:

```go
app.Get("/users", func(c *hawk.Context) {
    // use c here
})
```

---

## JSON Request Body

### `BindJSON`

```go
func (c *Context) BindJSON(obj any) error
```

Decodes the JSON request body into the supplied Go value.

Example:

```go
type UserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func Store(c *hawk.Context) {
    var input UserRequest

    if err := c.BindJSON(&input); err != nil {
        c.AbortWithError(400, hawk.H{
            "error": err.Error(),
        })
        return
    }

    // input.Name
    // input.Email
}
```

The request should have:

```http
Content-Type: application/json
```

---

## JSON Binding and Validation

### `BindAndValidate`

```go
func (c *Context) BindAndValidate(obj any) error
```

Combines JSON binding and Hawk validation.

It:

1. Reads the JSON request body.
2. Validates the resulting object.
3. Automatically returns a validation response when validation fails.

Example:

```go
type UserRequest struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
}

func Store(c *hawk.Context) {
    var input UserRequest

    if err := c.BindAndValidate(&input); err != nil {
        return
    }

    // Continue with valid input.
}
```

---

# Query Parameters

### `Query`

```go
func (c *Context) Query(key string) string
```

Returns a query parameter from the URL.

Request:

```text
GET /users?page=2
```

Usage:

```go
page := c.Query("page")
```

Result:

```text
"2"
```

---

### `QueryInt`

```go
func (c *Context) QueryInt(key string, defaultValue int) int
```

Returns a query parameter as an integer.

If the parameter is missing or cannot be converted to an integer, the supplied default value is returned.

Example:

```go
page := c.QueryInt("page", 1)
```

For:

```text
GET /users?page=3
```

`page` is:

```text
3
```

For:

```text
GET /users
```

`page` is:

```text
1
```

For:

```text
GET /users?page=abc
```

`page` is also:

```text
1
```

---

# Route Parameters

### `Param`

```go
func (c *Context) Param(key string) string
```

Returns a parameter captured from the route.

Example route:

```go
app.Get("/users/:id", func(c *hawk.Context) {
    id := c.Param("id")

    c.JSON(200, hawk.H{
        "id": id,
    })
})
```

Request:

```text
GET /users/15
```

Result:

```text
id = "15"
```

Wildcard parameters are also available through `Param`.

Example:

```go
app.Get("/files/*filepath", func(c *hawk.Context) {
    filepath := c.Param("filepath")
})
```

---

# Responses

## JSON

### `JSON`

```go
func (c *Context) JSON(status int, data any) error
```

Returns a JSON response.

Example:

```go
c.JSON(200, hawk.H{
    "message": "Hello Hawk",
})
```

Response:

```json
{
    "message": "Hello Hawk"
}
```

Hawk automatically sets:

```http
Content-Type: application/json
```

`JSON()` also supports Hawk's `Resource` and `CollectionResource` types.

---

## String

### `String`

```go
func (c *Context) String(status int, data string) error
```

Returns a plain-text response.

Example:

```go
c.String(200, "Hello Hawk")
```

Hawk sets:

```http
Content-Type: text/plain
```

---

## HTML

### `HTML`

```go
func (c *Context) HTML(status int, html string) error
```

Returns an HTML response.

Example:

```go
c.HTML(200, "<h1>Hello Hawk</h1>")
```

Hawk sets:

```http
Content-Type: text/html
```

---

## Status

### `Status`

```go
func (c *Context) Status(code int)
```

Sets the HTTP response status code.

Example:

```go
c.Status(201)
```

Usually, handlers can use `JSON`, `String`, or `HTML` directly because those methods also set the status.

---

# Form Data

## Form

### `Form`

```go
func (c *Context) Form(key string) string
```

Returns a form field.

Hawk parses multipart form data before retrieving the value.

This is especially useful when a request contains both normal fields and uploaded files.

Example:

```go
title := c.Form("title")
content := c.Form("content")
```

A multipart request can contain:

```text
title
content
cover
```

and Hawk can retrieve the text fields with `Form()` while retrieving the file with `File()` or `OpenFile()`.

---

## Default Form Value

### `DefaultForm`

```go
func (c *Context) DefaultForm(key, defaultValue string) string
```

Returns a form value or a default value when the field is empty.

Example:

```go
title := c.DefaultForm("title", "Untitled")
```

If `title` is not supplied:

```text
Untitled
```

is returned.

---

## All Form Values

### `Forms`

```go
func (c *Context) Forms() map[string][]string
```

Returns all parsed form values.

Example:

```go
forms := c.Forms()

for key, values := range forms {
    // use key and values
}
```

---

# File Uploads

## File

### `File`

```go
func (c *Context) File(name string) (*multipart.FileHeader, error)
```

Returns the metadata for an uploaded file.

Example:

```go
header, err := c.File("cover")

if err != nil {
    c.AbortWithError(400, hawk.H{
        "error": err.Error(),
    })
    return
}

fmt.Println(header.Filename)
fmt.Println(header.Size)
```

---

## Open File

### `OpenFile`

```go
func (c *Context) OpenFile(name string) (
    multipart.File,
    *multipart.FileHeader,
    error,
)
```

Opens an uploaded file and returns both the file and its header.

Example:

```go
file, header, err := c.OpenFile("cover")

if err != nil {
    c.AbortWithError(400, hawk.H{
        "error": err.Error(),
    })
    return
}

defer file.Close()

fmt.Println(header.Filename)
```

This is useful when storing uploaded files:

```go
file, header, err := c.OpenFile("cover")

if err != nil {
    return
}

defer file.Close()

err = storage.Default().Put(
    "uploads/"+header.Filename,
    file,
)
```

---

# Cookies

## Get Cookie

### `Cookie`

```go
func (c *Context) Cookie(key string) (string, error)
```

Returns the value of a cookie.

Example:

```go
token, err := c.Cookie("token")

if err != nil {
    // cookie doesn't exist
}
```

---

## Set Cookie

### `SetCookie`

```go
func (c *Context) SetCookie(
    name string,
    value string,
    maxAge int,
    path string,
    domain string,
    secure bool,
    httpOnly bool,
)
```

Sets an HTTP cookie.

Example:

```go
c.SetCookie(
    "token",
    token,
    3600,
    "/",
    "",
    true,
    true,
)
```

---

## All Cookies

### `Cookies`

```go
func (c *Context) Cookies() []*http.Cookie
```

Returns all cookies sent with the request.

Example:

```go
cookies := c.Cookies()

for _, cookie := range cookies {
    fmt.Println(cookie.Name, cookie.Value)
}
```

---

## Delete Cookie

### `DeleteCookie`

```go
func (c *Context) DeleteCookie(name string)
```

Removes a cookie by setting its maximum age to `-1`.

Example:

```go
c.DeleteCookie("token")
```

---

# Middleware

## Next

### `Next`

```go
func (c *Context) Next()
```

Continues execution of the middleware and handler chain.

Hawk middleware can call:

```go
c.Next()
```

to continue processing the request.

---

## Abort

### `Abort`

```go
func (c *Context) Abort()
```

Stops the remaining handlers from executing.

Example:

```go
if !authenticated {
    c.Abort()
    return
}
```

---

## Abort With Error

### `AbortWithError`

```go
func (c *Context) AbortWithError(status int, err any)
```

Sends a JSON error response and stops the remaining handlers.

Example:

```go
c.AbortWithError(401, hawk.H{
    "error": "Unauthorized",
})
```

This is useful for authentication, authorization, validation, and request errors.

---

## Replace Handlers

### `ReplaceHandlers`

```go
func (c *Context) ReplaceHandlers(handlers ...HandlerFunc)
```

Replaces the current handler chain with the supplied handlers.

This is primarily used internally by Hawk's router when it has found the route that should handle the request.

Application developers normally do not need to call this directly.

---

# Validation

## ValidationError

### `ValidationError`

```go
func (c *Context) ValidationError(err any)
```

Returns Hawk's standard validation error response.

Example:

```go
c.ValidationError(errors)
```

The response uses HTTP status:

```text
422 Unprocessable Entity
```

and has the structure:

```json
{
    "message": "validation failed",
    "errors": {}
}
```

`BindAndValidate()` uses this method automatically when validation fails.

---

# Context Usage Example

A complete handler can combine several Context features:

```go
func Store(c *hawk.Context) {
    title := c.Form("title")

    file, header, err := c.OpenFile("cover")
    if err != nil {
        c.AbortWithError(400, hawk.H{
            "error": err.Error(),
        })
        return
    }

    defer file.Close()

    // Save file...

    c.JSON(201, hawk.H{
        "message": "Post created successfully",
        "title":   title,
        "file":    header.Filename,
    })
}
```

The `Context` is designed to keep common HTTP operations inside Hawk's API so application code does not need to repeatedly work directly with `http.Request` and `http.ResponseWriter`.
