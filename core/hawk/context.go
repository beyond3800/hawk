package hawk

// Don't edit or add anything here
import (
	"encoding/json"
	"net/http"

	"github.com/beyond3800/hawk/core/validation"
)

// context
func (c *Context) BindJSON(obj any) error{
	return json.NewDecoder(c.Request.Body).Decode(obj)
}
func (c *Context) BindAndValidate(obj any) error {
    
    if err := c.BindJSON(obj); err != nil {
        return err
    }
    c.ValidationError(validation.Validate(obj))
    return nil
}
func (c *Context) Query(key string) string{
	return c.Request.URL.Query().Get(key)
}
func (c *Context) JSON(status int, data any) error {
    c.Response.Header().Set("Content-Type", "application/json")
    c.Status(status)

    return json.NewEncoder(c.Response).Encode(data)
}
func (c *Context) String(status int, data string) error {
    c.Response.Header().Set("Content-Type", "text/plain")
    c.Status(status)
	_,err := c.Response.Write([]byte(data))
	return err
}
func (c *Context) Param(key string) string{
    return c.params[key]
}
func (c *Context) HTML(status int, html string) error {
    c.Response.Header().Set("Content-Type", "text/html")
    c.Status(status)
    _, err := c.Response.Write([]byte(html))
    return err
}
func (c *Context) Status(code int) {
	c.statusCode = code
    c.Response.WriteHeader(code)
}
func (c *Context) Abort(){
	c.index = len(c.handlers) 
}
func (c *Context) Next() {
    c.index++

    for c.index < len(c.handlers) {
        c.handlers[c.index](c)
        c.index++
    }
}
func (c *Context) ValidationError(err any) {
    c.JSON(http.StatusUnprocessableEntity, map[string]any{
        "message": "validation failed",
        "errors":  err,
    })
}
func (c *Context) Cookie(key string) (string, error) {
    cookie, err := c.Request.Cookie(key)
    if err != nil {
        return "", err
    }
    return cookie.Value, nil
}
func (c *Context) SetCookie(
    name string,
	value string,
	maxAge int,
	path string,
	domain string,
	secure bool,
	httpOnly bool,
    ) {
        http.SetCookie(c.Response, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Domain:   domain,
		MaxAge:   maxAge,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
}
func (c *Context) Cookies() []*http.Cookie {
	return c.Request.Cookies()
}
func (c *Context) DeleteCookie(name string) {
	http.SetCookie(c.Response, &http.Cookie{
		Name:   name,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}