package hawk


import (
	"net/http"
	"time"
)

func (h *Hawk) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    for _, route := range h.routes {

        if route.Method != request.Method {
            continue
        }

		handlers := make([]HandlerFunc,0)
		handlers = append(handlers, h.middleware...)
		handlers = append(handlers, route.Handler...)

		matched, params := match(route.Pattern, request.URL.Path)
        if matched {
            c := &Context{
                Response: response,
                Request: request,
				params: params,
				handlers: handlers,
				index: -1,
            }

			c.startTime = time.Now()
            c.Next()
            return
        }
    }

    http.NotFound(response, request)
}

func (h *Hawk) Use(handler ...HandlerFunc){
	h.middleware = append(h.middleware, handler...)
}

func (h *Hawk) Run(addr string) error {
	return http.ListenAndServe(addr, h)
}

func (g *RouterGroup) Use(handler HandlerFunc){
	g.handler = append(g.handler, handler)
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
