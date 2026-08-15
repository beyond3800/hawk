package hawk

import (
	"net/http"
	"os"
	"fmt"
	"strings"
)

func (h *Hawk) addRoute(method string, path string, handler ...HandlerFunc){
	h.routes = append(h.routes, Route{
		Method: method,
		Pattern: path,
		Handler: handler,
	})
}
func (h *Hawk) Get(path string, handler ...HandlerFunc) {
    h.addRoute(
		"GET",
		path,
		handler...,
	)
}
func (h *Hawk) Post(path string, handler ...HandlerFunc) {
    h.addRoute(
		"POST",
		path,
		handler...,
	)
}
func (h *Hawk) Put(path string, handler ...HandlerFunc) {
    h.addRoute(
		"PUT",
		path,
		handler...,
	)	
}
func (h *Hawk) Delete(path string, handler ...HandlerFunc) {
    h.addRoute(
		"DELETE",
		path,
		handler...,
	)
}
func (h *Hawk) Static(prefix string, directory string) {
	h.Get(prefix+"/*filepath", func(c *Context) {
		filepath := c.Param("filepath")

		if filepath == "" {
			filepath = "index.html"
		}
		fmt.Println("filepath",filepath)
		fullPath, err := safeStaticPath(directory, filepath)
		if err != nil {
			http.NotFound(c.Response, c.Request)
			return
		}

		info, err := os.Stat(fullPath)
		if err != nil || info.IsDir() {
			http.NotFound(c.Response, c.Request)
			return
		}

		http.ServeFile(c.Response, c.Request, fullPath)
	})
}

func joinPaths(prefix, path string) string {
    if path == "" {
        return "/" + strings.Trim(prefix, "/")
    }
    return "/" + strings.Trim(prefix, "/") + "/" + strings.Trim(path, "/")
}
