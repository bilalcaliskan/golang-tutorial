package main

import (
	"log"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func hardcodedTimeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func closuredTimeHandlerFunc(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

func runRequestHandling() {
	/*
	Processing HTTP requests with Go is primarily about two things: ServeMuxes and Handlers.
	A ServeMux is essentially a HTTP request router (or multiplexor). It compares incoming requests against a list of
	predefined URL paths, and calls the associated handler for the path whenever a match is found.
	Handlers are responsible for writing response headers and bodies. Almost any object can be a handler, so long as it
	satisfies the http.Handler interface. In lay terms, that simply means it must have a ServeHTTP method with the
	following signature:
		ServeHTTP(http.ResponseWriter, *http.Request)
	Go's HTTP package ships with a few functions to generate common handlers, such as FileServer, NotFoundHandler and
	RedirectHandler.
	*/
	mux := http.NewServeMux()

	// Implementation of RedirectHandler
	rh := http.RedirectHandler("http://example.org", 307)
	mux.Handle("/foo", rh)

	// Implementation of custom handler 1
	th1123 := &timeHandler{format: time.RFC1123}
	mux.Handle("/time/rfc1123", th1123)

	th3339 := &timeHandler{format: time.RFC3339}
	mux.Handle("/time/rfc3339", th3339)

	/*
	Implementation of functions as handlers:
		1- Convert the timeHandler function to a HandlerFunc type
		2- And add it to the ServeMux
	*/
	thfunc := http.HandlerFunc(hardcodedTimeHandlerFunc)
	mux.Handle("/time/hardcodedTimeHandlerFunc", thfunc)

	// Implementation of mux.HandleFunc
	thfunc2 := closuredTimeHandlerFunc(time.RFC1123)
	mux.Handle("/time/closuredTimeHandlerFunc", thfunc2)

	/*
	DefaultServeMux
	Generally you shouldn't use the DefaultServeMux because it poses a security risk.
	Because the DefaultServeMux is stored in a global variable, any package is able to access it and register a
	route – including any third-party packages that your application imports. If one of those third-party packages is
	compromised, they could use the DefaultServeMux to expose a malicious handler to the web.
	So as a rule of thumb it's a good idea to avoid the DefaultServeMux, and instead use your own locally-scoped
	ServeMux, like we have been so far.

	ListenAndServe will fall back to using the DefaultServeMux if no other handler is provided (that is, the second
	parameter is set to nil).
		http.ListenAndServe(":3000", nil) // it will use the DefaultServeMux
	*/

	log.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", mux)
}