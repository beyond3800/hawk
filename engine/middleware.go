package hawk

import (
	"fmt"
	"net/http"
	"strconv"
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

	// Defaults
	if len(config.AllowMethods) == 0 {
		config.AllowMethods = []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		}
	}

	if len(config.AllowHeaders) == 0 {
		config.AllowHeaders = []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		}
	}

	if config.MaxAge == 0 {
		config.MaxAge = 12 * time.Hour
	}

	// Invalid configuration
	if config.AllowCredentials &&
		len(config.AllowOrigins) == 1 &&
		config.AllowOrigins[0] == "*" {
		panic("cors: AllowCredentials cannot be true when AllowOrigins contains '*'")
	}

	return func(c *Context) {

		origin := c.Request.Header.Get("Origin")

		if origin == "" {
			c.Next()
			return
		}

		allowed := false
		allowAll := false

		for _, o := range config.AllowOrigins {

			if o == "*" {
				allowAll = true
				allowed = true
				break
			}

			if o == origin {
				allowed = true
				break
			}
		}

		if !allowed {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		// Allow Origin
		if allowAll {
			c.Response.Header().Set(
				"Access-Control-Allow-Origin",
				"*",
			)
		} else {
			c.Response.Header().Set(
				"Access-Control-Allow-Origin",
				origin,
			)

			c.Response.Header().Set(
				"Vary",
				"Origin",
			)
		}

		// Methods
		c.Response.Header().Set(
			"Access-Control-Allow-Methods",
			strings.Join(config.AllowMethods, ", "),
		)

		// Headers
		c.Response.Header().Set(
			"Access-Control-Allow-Headers",
			strings.Join(config.AllowHeaders, ", "),
		)

		// Exposed Headers
		if len(config.ExposeHeaders) > 0 {
			c.Response.Header().Set(
				"Access-Control-Expose-Headers",
				strings.Join(config.ExposeHeaders, ", "),
			)
		}

		// Credentials
		if config.AllowCredentials {
			c.Response.Header().Set(
				"Access-Control-Allow-Credentials",
				"true",
			)
		}

		// Max Age
		c.Response.Header().Set(
			"Access-Control-Max-Age",
			strconv.Itoa(int(config.MaxAge.Seconds())),
		)

		// Preflight
		if c.Request.Method == http.MethodOptions {
			c.Status(http.StatusNoContent)
			c.Abort()
			return
		}

		c.Next()
	}
}