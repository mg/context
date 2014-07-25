context
=======

A library to provide context for http functions in Go. Has the following properties:

    - Does not rely on a global lock to protect access to the context.
    - Works with the plain http interface, no need for custom http function signatures.
    - Works with any middleware library that depends on the standard http interface.

Example using [alice](https://github.com/justinas/alice) for middleware:

    import (
        "net/http"
        "github.com/justinas/alice"
        "github.com/mg/context"
    )
    
    func handler1(w http.ResponseWriter, r *http.Request) {
        // access context and store value on it
        c := context.Access(r)
        c.Set("Key1", "Value") // Value can be any object
    }
    
    func handler2(w http.ResponseWriter, r *http.Request) {
        // access context and retrieve value from it
        c := context.Access(r)
        v := c.GetAsString("Key1") // v contains "Value"
    }
    
    func main() {
        http.Handle("/", alice.New(handler1, handler2))
        http.ListenAndServe(":8080", nil)
    }

*context.Access()* fetch the context associated with *r*. If there is none, *context.Access()* will create one and bind it to *r*.

In the above example *GetAsString* was used to fetch the value. The context contains the method *Get(key)* that simply returns an interface{} value. *GetAsString* is a convenience function that simply typecast the value to string before returning, returning "" if it fails. 

The library provides two middlewares to store values on the context.

    import (
        "net/http"
        "github.com/justinas/alice"
        "github.com/mg/context"
    )
    
    func handler(w http.ResponseWriter, r *http.Request) {
        c := context.Access(r)
        // c contains values at keys k1, k2 and k3
    }
    
    func main() {
        http.Handle("/", 
            alice.New(
                context.Store("k1", "val")
                context.StoreMany("k2", "val1", "k3", "val2"),
                handler
            ))
        http.ListenAndServe(":8080", nil)
    }

This technique of binding the context to the request is lifted from the context found in bradrydzewki [routes](https://github.com/bradrydzewski/routes/blob/master/exp/context/context.go) library.
