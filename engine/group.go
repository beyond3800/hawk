package hawk


func (h *Hawk) Group(prefix string) *RouterGroup{
	return &RouterGroup{
		prefix: prefix,
		engine: h,
	}
}
func (g *RouterGroup) addRoute(method string, fullPath string, handler ...HandlerFunc){

	handlers := append([]HandlerFunc{}, g.middleware...)
	handlers = append(handlers, handler...)

	g.engine.routes = append(g.engine.routes, Route{
        Method:     method,
        Pattern:    fullPath,
        Handler:    handlers,
    })
}
func (g *RouterGroup) Get(path string, handler HandlerFunc) {
    fullPath := joinPaths(g.prefix , path) 
	g.addRoute("GET", fullPath, handler)
}
func (g *RouterGroup) Post(path string, handler HandlerFunc) {
    fullPath := joinPaths(g.prefix , path) 
	g.addRoute("POST", fullPath, handler)
}
func (g *RouterGroup) Put(path string, handler HandlerFunc) {
    fullPath := joinPaths(g.prefix , path) 
	g.addRoute("PUT", fullPath, handler)
}
func (g *RouterGroup) Delete(path string, handler HandlerFunc) {
    fullPath := joinPaths(g.prefix , path) 
	g.addRoute("DELETE", fullPath, handler)
}