package hawk

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func Logger(c *Context) {

    start := time.Now()

    c.Next()

    latency := time.Since(start)

    fmt.Printf(
        "%s %s %d %v\n",
        c.Request.Method,
        c.Request.URL.Path,
        c.statusCode,
        latency,
    )
}

func Recovery(c *Context) {
    defer func() {
        if err := recover(); err != nil {
            c.JSON(500, "error")
        }
    }()

    c.Next()
}

func Cors(config CorsConfig) HandlerFunc {
    return func(c *Context) {

        requestOrigin := c.Request.Header.Get("Origin")


        if len(config.AllowOrigins) > 0 {

            if requestOrigin != "" && requestOrigin != config.AllowOrigins[0] {

                c.JSON(http.StatusForbidden, map[string]string{
                    "error": "origin not allowed",
                })

                c.Abort()
                return
            }
        }

        if requestOrigin != "" {
            c.Response.Header().Set(
                "Access-Control-Allow-Origin",
                requestOrigin,
            )
        }

        c.Response.Header().Set(
            "Access-Control-Allow-Methods",
            strings.Join(config.AllowMethods, ", "),
        )

        c.Response.Header().Set(
            "Access-Control-Allow-Headers",
            strings.Join(config.AllowHeaders, ", "),
        )

        if config.AllowCredentials {
            c.Response.Header().Set(
                "Access-Control-Allow-Credentials",
                "true",
            )
        }

        if len(config.ExposeHeaders) > 0 {
            c.Response.Header().Set(
                "Access-Control-Expose-Headers",
                strings.Join(config.ExposeHeaders, ", "),
            )
        }

        if c.Request.Method == http.MethodOptions {
            c.Status(http.StatusNoContent)
            c.Abort()
            return
        }

        c.Next()
    }
}