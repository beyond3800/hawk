package hawk

import (
	"fmt"
	"os"
	"path/filepath"
	"net/http"
	"strings"
	"time"
)
type HandlerFunc func (*Context)

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
	middleware  []HandlerFunc
}

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

type ErrorResponse struct {
    Error    string `json:"error"`
    Message  any    `json:"message"`
    Code     int    `json:"code"`
}

type SuccessResponse struct {
    Success   string  `json:"success"`
    Message   any     `json:"message"`
    Code      int     `json:"code"`
}

// func match(pattern, path string) (bool, map[string]string) {
//     params := make(map[string]string)

//     patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
//     pathParts := strings.Split(strings.Trim(path, "/"), "/")
	
//     if len(patternParts) != len(pathParts) {
//         return false, nil
//     }

//     for i := range patternParts {

// 		if strings.HasPrefix(patternParts[i], "*") {
// 			key := patternParts[i][1:]
// 			params[key] = strings.Join(pathParts[i:], "/")
// 			return true, params
// 		}
//         if strings.HasPrefix(patternParts[i], ":") {

//             key := patternParts[i][1:]
//             params[key] = pathParts[i]

//             continue
//         }

//         if patternParts[i] != pathParts[i] {
//             return false, nil
//         }
//     }

//     return true, params
// }
func match(pattern, path string) (bool, map[string]string) {
	params := make(map[string]string)

	patternParts := strings.Split(
		strings.Trim(pattern, "/"),
		"/",
	)

	pathParts := strings.Split(
		strings.Trim(path, "/"),
		"/",
	)

	for i := range patternParts {

		// Wildcard parameter
		if strings.HasPrefix(patternParts[i], "*") {
			key := patternParts[i][1:]

			params[key] = strings.Join(pathParts[i:], "/")

			return true, params
		}

		// We don't have a corresponding path segment.
		if i >= len(pathParts) {
			return false, nil
		}

		// Normal parameter
		if strings.HasPrefix(patternParts[i], ":") {
			key := patternParts[i][1:]
			params[key] = pathParts[i]

			continue
		}

		// Static segment
		if patternParts[i] != pathParts[i] {
			return false, nil
		}
	}

	// The route must consume the entire path
	// unless a wildcard already returned above.
	if len(patternParts) != len(pathParts) {
		return false, nil
	}

	return true, params
}

func isStaticRoute(pattern string) bool {
	parts := strings.Split(strings.Trim(pattern, "/"), "/")

	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			return false
		}
	}

	return true
}

func safeStaticPath(directory, requested string) (string, error) {
	base, err := filepath.Abs(directory)
	if err != nil {
		return "", err
	}

	target, err := filepath.Abs(
		filepath.Join(base, requested),
	)
	if err != nil {
		return "", err
	}

	relative, err := filepath.Rel(base, target)
	if err != nil {
		return "", err
	}

	if relative == ".." ||
		strings.HasPrefix(
			relative,
			".."+string(os.PathSeparator),
		) {
		return "", fmt.Errorf("path escapes static directory")
	}

	fmt.Println(directory)
	fmt.Println(requested)
	fmt.Println(target)
	return target, nil
}

func isWildcardRoute(pattern string) bool {
	parts := strings.Split(
		strings.Trim(pattern, "/"),
		"/",
	)

	for _, part := range parts {
		if strings.HasPrefix(part, "*") {
			return true
		}
	}

	return false
}
func wildcardScore(pattern string) int {
    parts := strings.Split(
        strings.Trim(pattern, "/"),
        "/",
    )

    score := 0

    for _, part := range parts {
        if !strings.HasPrefix(part, "*") {
            score++
        }
    }

    return score
}