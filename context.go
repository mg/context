package context

import (
	"io"
	"net/http"
)

type (
	values map[interface{}]interface{}

	Context struct {
		values values
		req    *http.Request
	}
)

// Access context associated with request. Create new if none available.
func Access(r *http.Request) *Context {

	// Get the context bound to the http.Request
	if v, ok := r.Body.(*wrapper); ok {
		return v.context
	}

	// Create a new context
	c := Context{}
	c.values = make(values)
	c.req = r

	// Wrap the request and bind the context
	wrapper := wrap(r)
	wrapper.context = &c
	return &c
}

// Wrapper for body to pass context along with request
type wrapper struct {
	body    io.ReadCloser // the original message body
	context *Context
}

func wrap(r *http.Request) *wrapper {
	w := wrapper{body: r.Body}
	r.Body = &w
	return &w
}

func (w *wrapper) Read(p []byte) (n int, err error) {
	return w.body.Read(p)
}

func (w *wrapper) Close() error {
	return w.body.Close()
}

// Return value at key
func (c *Context) Get(key interface{}) interface{} {
	return c.values[key]
}

// Return value at key as string
func (c *Context) GetAsString(key interface{}) string {
	val := c.Get(key)
	if val == nil {
		return ""
	}

	if str, ok := val.(string); ok {
		return str
	}
	return ""
}

// Set key to value
func (c *Context) Set(key, value interface{}) {
	c.values[key] = value
}

// Delete value at key
func (c *Context) Del(key interface{}) {
	delete(c.values, key)
}
