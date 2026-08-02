package hawk

// Don't edit or add anything here
import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/beyond3800/hawk/validation"
)

// context
func (c *Context) BindJSON(obj any) error{
    err := json.NewDecoder(c.Request.Body).Decode(obj)
    if err != nil{
        c.Abort()
        return fmt.Errorf("Unable to bind data")
    }
	return nil
}
func (c *Context) BindAndValidate(obj any) error {
    
    if err := c.BindJSON(obj); err != nil {
        return err
    }
    errors, err := validation.Validate(obj)
    if err != nil{
        c.ValidationError(errors)
        c.Abort()
        return err
    }
    return nil
}
func (c *Context) Query(key string) string{
	return c.Request.URL.Query().Get(key)
}
func (c *Context) JSON(status int, data any) error {
    switch value := data.(type) {

	case Resource:
		data = value.ToMap()

	case CollectionResource:
		data = value.ToSlice()

	}
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
func (c *Context) AbortWithError(status int, err any){
    c.JSON(status, err)
    c.Abort()
}
func (c *Context) Next() {
    c.index++

    for c.index < len(c.handlers) {
        c.handlers[c.index](c)
        c.index++
    }
}
func (c *Context) ValidationError(err any) {
    c.JSON(http.StatusUnprocessableEntity, H{
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
func (c *Context) Form(key string) string {
    _ = c.Request.ParseForm()
    return c.Request.FormValue(key)
}
func (c *Context) DefaultForm(key, defaultValue string) string {
    value := c.Form(key)
    if value == "" {
        return defaultValue
    }
    return value
}
func (c *Context) Forms() map[string][]string {
    _ = c.Request.ParseForm()
    return c.Request.Form
}
func (c *Context) File(name string) (*multipart.FileHeader, error) {
    err := c.Request.ParseMultipartForm(32 << 20)
    if err != nil {
        return nil, err
    }

    file, _, err := c.Request.FormFile(name)
    if file != nil {
        file.Close()
    }

    _, header, err := c.Request.FormFile(name)
    return header, err
}

func (c *Context) OpenFile(name string) (multipart.File, *multipart.FileHeader, error) {
    header, err := c.File(name)
    if err != nil {
        return nil, nil, err
    }
    file, _, _ := c.Request.FormFile(name)  
    return file, header, nil

}
func (c *Context) ReplaceHandlers(handlers ...HandlerFunc) {
	c.handlers = handlers
	c.index = -1
}
func (c *Context) QueryInt(key string, defaultValue int) int {

	value := c.Query(key)

	if value == "" {
		return defaultValue
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return number
}