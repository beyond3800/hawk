

# `docs/middleware.md`

```md
# Middleware

## Global Middleware

```go

app.Use(Logger())
app.Use(Recovery())

api := app.Group("/api")

api.Use(Auth())

api.Get("/profile", ProfileController.Show)


Create middleware

func Auth() hawk.HandlerFunc {
    return func(c *hawk.Context) {

        token := c.Request.Header.Get("Authorization")

        if token == "" {
            c.JSON(401, map[string]string{
                "error": "Unauthorized",
            })

            c.Abort()
            return
        }

        c.Next()
    }
}