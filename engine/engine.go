package hawk

import (
	"fmt"
	"net/http"
	"time"
)

type H map[string]any

func (h *Hawk) ServeHTTP(
	response http.ResponseWriter,
	request *http.Request,
) {

	c := &Context{
		Response: response,
		Request: request,
		index: -1,
	}

	handlers := append(
		[]HandlerFunc{},
		h.middleware...,
	)

	handlers = append(
		handlers,
		h.router(),
	)

	c.handlers = handlers

	c.startTime = time.Now()

	c.Next()
}
func (h *Hawk) router() HandlerFunc {

	return func(c *Context) {

		for _, route := range h.routes {

			if route.Method != c.Request.Method {
				continue
			}

			matched, params := match(
				route.Pattern,
				c.Request.URL.Path,
			)

			if !matched {
				continue
			}

			c.params = params

			// Route middleware + controller
			c.ReplaceHandlers(route.Handler...)

			c.Next()

			return
		}

		http.NotFound(c.Response, c.Request)
	}
}
func (h *Hawk) Use(handler ...HandlerFunc){
	h.middleware = append(h.middleware, handler...)
}

func (g *RouterGroup) Use(handler HandlerFunc){
    g.middleware = append(g.middleware, handler)
}

func (h *Hawk) Run(port string) error {
    url :=fmt.Sprintf("Hawk server running at http://127.0.0.1%s\n", port)
    fmt.Println(url)
	return http.ListenAndServe(port, h)
}


func New() *Hawk {
    return &Hawk{
        routes: []Route{},
    }
} 

func Default() *Hawk {
    h := &Hawk{
        routes: []Route{},
    }

    // Default middleware
    h.Use(Logger)
    h.Use(Recovery)

    return h
}
