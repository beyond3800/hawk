# Request

Hawk provides access to incoming HTTP request data through the `Context`.

A request can contain several types of data:

```text
HTTP Request
│
├── URL
│   └── Query parameters
│
├── Headers
│
├── Cookies
│
├── JSON body
│
└── Form data
```

The underlying Go `*http.Request` is available through:

```go
c.Request
```

This means Hawk can work naturally with Go's standard `net/http` APIs.

---

## Accessing the Request

A Hawk handler receives the current context:

```go
router.Post("/users", func(c *hawk.Context) {

    request := c.Request

    // work with the request
})
```

The request is a standard Go HTTP request:

```go
*http.Request
```

You can therefore access properties such as:

```go
c.Request.Method
c.Request.URL
c.Request.Header
c.Request.Body
```

---

# HTTP Method

The HTTP method is available through:

```go
method := c.Request.Method
```

For example:

```go
router.Get("/users", func(c *hawk.Context) {

    if c.Request.Method != http.MethodGet {
        c.String(405, "Method not allowed")
        return
    }

    c.String(200, "GET request")
})
```

Common HTTP methods include:

```text
GET
POST
PUT
PATCH
DELETE
```

Hawk's router normally handles method matching before the route handler is executed.

---

# Request URL

The requested URL is available through:

```go
c.Request.URL
```

The path can be accessed with:

```go
path := c.Request.URL.Path
```

For:

```text
/users/123
```

the path is:

```text
/users/123
```

---

# Query Parameters

Query parameters are values supplied after `?` in the URL.

For example:

```text
/users?page=2&limit=10
```

The query parameters are:

```text
page = 2
limit = 10
```

Hawk provides the `Query()` method through the Context:

```go
page := c.Query("page")
limit := c.Query("limit")
```

Example:

```go
router.Get("/users", func(c *hawk.Context) {

    page := c.Query("page")
    limit := c.Query("limit")

    c.JSON(200, map[string]any{
        "page":  page,
        "limit": limit,
    })
})
```

---

## Missing Query Parameters

If a query parameter doesn't exist, the returned value is an empty string.

For example:

```text
/users
```

```go
page := c.Query("page")
```

will return:

```text
""
```

You can handle this yourself:

```go
page := c.Query("page")

if page == "" {
    page = "1"
}
```

---

# Headers

HTTP headers provide additional information about a request.

Headers can be accessed through:

```go
c.Request.Header
```

For example:

```go
authorization := c.Request.Header.Get("Authorization")
```

Another example:

```go
contentType := c.Request.Header.Get("Content-Type")
```

---

## Reading Custom Headers

Suppose a client sends:

```text
X-API-Key: abc123
```

You can retrieve it with:

```go
apiKey := c.Request.Header.Get("X-API-Key")
```

Example:

```go
router.Get("/protected", func(c *hawk.Context) {

    apiKey := c.Request.Header.Get("X-API-Key")

    if apiKey == "" {
        c.JSON(401, map[string]any{
            "error": "API key required",
        })
        return
    }

    c.JSON(200, map[string]any{
        "message": "Authenticated",
    })
})
```

---

# JSON Request Body

JSON is commonly used when building APIs.

For example, a client may send:

```json
{
    "name": "Adam",
    "email": "adam@test.com"
}
```

The JSON data is contained in:

```go
c.Request.Body
```

You can decode it using Go's `encoding/json` package.

```go
var input struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

err := json.NewDecoder(
    c.Request.Body,
).Decode(&input)

if err != nil {
    c.JSON(400, map[string]any{
        "error": "Invalid JSON",
    })
    return
}
```

The decoded values can then be used:

```go
fmt.Println(input.Name)
fmt.Println(input.Email)
```

---

## JSON Request Example

A complete route might look like:

```go
router.Post("/users", func(c *hawk.Context) {

    var input struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }

    err := json.NewDecoder(
        c.Request.Body,
    ).Decode(&input)

    if err != nil {
        c.JSON(400, map[string]any{
            "error": "Invalid JSON body",
        })
        return
    }

    c.JSON(201, map[string]any{
        "name":  input.Name,
        "email": input.Email,
    })
})
```

A client can send:

```http
POST /users
Content-Type: application/json
```

with:

```json
{
    "name": "Adam",
    "email": "adam@test.com"
}
```

---

# Checking Content Type

The `Content-Type` header tells the application what type of data is being sent.

```go
contentType := c.Request.Header.Get("Content-Type")
```

For JSON:

```text
application/json
```

For URL-encoded form data:

```text
application/x-www-form-urlencoded
```

For multipart forms:

```text
multipart/form-data
```

You can check the content type before processing the request body.

---

# Form Data

HTML forms and some API clients send data as form fields.

For example:

```text
name=Adam&email=adam@test.com
```

The request can be parsed using Go's standard HTTP functionality:

```go
err := c.Request.ParseForm()

if err != nil {
    c.JSON(400, map[string]any{
        "error": "Unable to parse form",
    })
    return
}
```

After parsing, form values can be accessed through:

```go
name := c.Request.FormValue("name")
email := c.Request.FormValue("email")
```

Example:

```go
router.Post("/users", func(c *hawk.Context) {

    err := c.Request.ParseForm()

    if err != nil {
        c.JSON(400, map[string]any{
            "error": "Invalid form data",
        })
        return
    }

    name := c.Request.FormValue("name")
    email := c.Request.FormValue("email")

    c.JSON(201, map[string]any{
        "name":  name,
        "email": email,
    })
})
```

---

# Multipart Form Data

Multipart forms are commonly used for file uploads.

A request may contain:

```text
multipart/form-data
```

You can parse the multipart request using Go's standard HTTP API.

For example:

```go
err := c.Request.ParseMultipartForm(
    10 << 20,
)

if err != nil {
    c.JSON(400, map[string]any{
        "error": "Unable to parse multipart form",
    })
    return
}
```

The exact handling of uploaded files depends on the application.

Go provides APIs such as:

```go
c.Request.FormFile("file")
```

for retrieving uploaded files.

---

# Cookies

Cookies are sent by the client using the `Cookie` header.

Hawk exposes the underlying request, so cookies can be accessed using Go's standard API.

```go
cookie, err := c.Request.Cookie("session")
```

Check for errors:

```go
cookie, err := c.Request.Cookie("session")

if err != nil {
    c.JSON(401, map[string]any{
        "error": "Session cookie not found",
    })
    return
}
```

The cookie value can then be accessed with:

```go
value := cookie.Value
```

---

# Setting Cookies

Cookies are part of the HTTP response rather than the request.

You can set a cookie using Go's standard library:

```go
http.SetCookie(
    c.Response,
    &http.Cookie{
        Name:  "session",
        Value: "abc123",
    },
)
```

For example:

```go
router.Post("/login", func(c *hawk.Context) {

    http.SetCookie(
        c.Response,
        &http.Cookie{
            Name:  "session",
            Value: "abc123",
            HttpOnly: true,
        },
    )

    c.JSON(200, map[string]any{
        "message": "Logged in",
    })
})
```

---

# Request Body

The raw request body is available through:

```go
c.Request.Body
```

It implements:

```go
io.ReadCloser
```

This allows you to process raw request data when necessary.

For JSON requests, prefer using:

```go
json.NewDecoder(c.Request.Body)
```

rather than manually reading and parsing JSON.

---

# Request Size

When accepting request bodies from clients, applications should consider limiting the amount of data accepted.

Go provides `http.MaxBytesReader` for limiting request body size.

For example:

```go
c.Request.Body = http.MaxBytesReader(
    c.Response,
    c.Request.Body,
    1<<20,
)
```

The example limits the request body to approximately 1 MB.

This is particularly useful for APIs that accept JSON or uploaded data.

---

# Combining Request Data

A request can contain several types of information at the same time.

For example:

```text
POST /users?page=1
Authorization: Bearer token
Cookie: session=abc123
Content-Type: application/json
```

with a JSON body:

```json
{
    "name": "Adam",
    "email": "adam@test.com"
}
```

A Hawk handler can access each part:

```go
router.Post("/users", func(c *hawk.Context) {

    // Query parameter
    page := c.Query("page")

    // Header
    authorization := c.Request.Header.Get(
        "Authorization",
    )

    // Cookie
    cookie, _ := c.Request.Cookie("session")

    // JSON body
    var input struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }

    err := json.NewDecoder(
        c.Request.Body,
    ).Decode(&input)

    if err != nil {
        c.JSON(400, map[string]any{
            "error": "Invalid JSON",
        })
        return
    }

    c.JSON(200, map[string]any{
        "page":           page,
        "authorization": authorization,
        "session":        cookie.Value,
        "name":           input.Name,
        "email":          input.Email,
    })
})
```

---

# Request Data Overview

| Data            | Hawk / Go API                     |
| --------------- | --------------------------------- |
| Route parameter | `c.Param()`                       |
| Query parameter | `c.Query()`                       |
| Header          | `c.Request.Header.Get()`          |
| Cookie          | `c.Request.Cookie()`              |
| JSON body       | `json.NewDecoder(c.Request.Body)` |
| Form data       | `c.Request.FormValue()`           |
| Multipart form  | `c.Request.ParseMultipartForm()`  |
| Raw body        | `c.Request.Body`                  |
| HTTP method     | `c.Request.Method`                |
| URL             | `c.Request.URL`                   |
| Path            | `c.Request.URL.Path`              |

---

# Request Validation

Request parsing and validation are separate responsibilities.

For example, decoding JSON:

```go
err := json.NewDecoder(
    c.Request.Body,
).Decode(&input)
```

only tells you whether the request body contains valid JSON.

It does not guarantee that the data is valid for your application.

After decoding, use Hawk's validation system:

```go
// Validate input here
```

For example, an application may require:

```text
name     → required
email    → required + valid email
password → required + minimum length
```

The request documentation and validation documentation should therefore be treated as separate concerns.

---

# Error Handling

When request data cannot be processed, return an appropriate HTTP status code.

For invalid JSON:

```go
c.JSON(400, map[string]any{
    "error": "Invalid JSON body",
})
```

For missing authentication:

```go
c.JSON(401, map[string]any{
    "error": "Unauthorized",
})
```

For invalid request data:

```go
c.JSON(422, map[string]any{
    "error": "Validation failed",
})
```

The exact error format should be standardized by the application.

---

# Best Practices

### Validate input

Never assume request data is valid.

### Check parsing errors

Always check errors when parsing:

* JSON
* forms
* multipart data
* cookies

### Validate Content-Type when necessary

If an endpoint expects JSON, verify that the request uses an appropriate content type.

### Limit large request bodies

Especially for endpoints that accept files or large JSON payloads.

### Don't trust headers or cookies

Headers and cookies come from the client and should not automatically be considered trustworthy.

### Separate parsing from business logic

A handler should generally:

```text
Receive request
      ↓
Parse request
      ↓
Validate input
      ↓
Call application logic
      ↓
Return response
```

rather than placing all application logic directly inside the request parser.

---

# Summary

Hawk exposes Go's standard HTTP request through:

```go
c.Request
```

This allows applications to work with:

```text
JSON
Forms
Headers
Cookies
Query parameters
Route parameters
Request bodies
HTTP methods
URLs
```

The main Hawk-specific helpers are:

```go
c.Param()
c.Query()
```

For other request data, Hawk works naturally with Go's `net/http` APIs.

This keeps request handling familiar to Go developers while providing convenient helpers for common Hawk functionality.
