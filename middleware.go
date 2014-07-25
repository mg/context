package context

import (
	"net/http"
)

// Store value object in context at key
func Store(key string, value interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := Access(r)
			c.Set(key, value)
			next.ServeHTTP(w, r)
		})
	}
}

/*
	Store keys/values on context.
	Params is assumed to be on the form
		["key1", "value1", "key2", "value2" ...]
	Panics if len(parmas) % 2 != 0
*/
func StoreMany(params ...string) func(http.Handler) http.Handler {
	if len(params)%2 != 0 {
		panic("context.SetKeys: odd number of params")
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := Access(r)
			for i := 0; i < len(params); i = i + 2 {
				c.Set(params[i], params[i+1])
			}
			next.ServeHTTP(w, r)
		})
	}
}
