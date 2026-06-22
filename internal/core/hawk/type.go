package hawk

import (
	"net/http"
	"strings"
	"time"
)
type HandlerFunc func (*Context)

type Context struct {
	Response  http.ResponseWriter
	Request  *http.Request
	params   map[string]string

	handlers []HandlerFunc
	aborted  bool
	index    int

	statusCode int
	startTime  time.Time
}

type CorsConfig struct{
	AllowOrigins       []string
	AllowMethods       []string
	AllowHeaders       []string
	ExposeHeaders      []string
	AllowCredentials   bool
	MaxAge             time.Duration
}

type Route struct{
	Method     string
	Pattern    string
	Handler    []HandlerFunc
}
type Hawk struct {
	routes     []Route
	middleware []HandlerFunc
}

type RouterGroup struct {
    prefix      string
    parent      *RouterGroup
    engine      *Hawk
	handler  []HandlerFunc
}
type ErrorResponse struct {
    Error   string `json:"error"`
    Message string `json:"message"`
    Code    int    `json:"code"`
}
type SuccessResponse struct {
    Success   string `json:"success"`
    Message string `json:"message"`
    Code    int    `json:"code"`
}

func match(pattern, path string) (bool, map[string]string) {
    params := make(map[string]string)

    patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
    pathParts := strings.Split(strings.Trim(path, "/"), "/")
	
    if len(patternParts) != len(pathParts) {
        return false, nil
    }

    for i := range patternParts {

        if strings.HasPrefix(patternParts[i], ":") {

            key := patternParts[i][1:]
            params[key] = pathParts[i]

            continue
        }

        if patternParts[i] != pathParts[i] {
            return false, nil
        }
    }

    return true, params
}