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
func (h *Hawk) Group(prefix string) *RouterGroup{
	return &RouterGroup{
		prefix: prefix,
		engine: h,
	}
}
func (g *RouterGroup) addRoute(method string, fullPath string, handler ...HandlerFunc){

	g.engine.routes = append(g.engine.routes, Route{
        Method:     method,
        Pattern:    fullPath,
        Handler:    handler,
    })
}
func joinPaths(prefix, path string) string {
    if path == "" {
        return "/" + strings.Trim(prefix, "/")
    }
    return "/" + strings.Trim(prefix, "/") + "/" + strings.Trim(path, "/")
}
func (g *RouterGroup) Get(path string, handler HandlerFunc) {
    fullPath := joinPaths(g.prefix , path) 
	g.addRoute("GET", fullPath, handler)
}
func (g *RouterGroup) Post(path string, handler HandlerFunc) {
    fullPath := joinPaths(g.prefix , path) 
	g.addRoute("POST",fullPath,handler)
}
func (g *RouterGroup) Put(path string, handler HandlerFunc) {
    fullPath := joinPaths(g.prefix , path) 
	g.addRoute("PUT",fullPath,handler)
}
func (g *RouterGroup) Delete(path string, handler HandlerFunc) {
    fullPath := joinPaths(g.prefix , path) 
	g.addRoute("DELETE",fullPath,handler)
}