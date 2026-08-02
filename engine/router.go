package hawk

import "strings"

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
func joinPaths(prefix, path string) string {
    if path == "" {
        return "/" + strings.Trim(prefix, "/")
    }
    return "/" + strings.Trim(prefix, "/") + "/" + strings.Trim(path, "/")
}
