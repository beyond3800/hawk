package middleware

import (
    "github.com/beyond3800/hawk/core/hawk"
)

func AdminMiddleware() hawk.HandlerFunc {
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