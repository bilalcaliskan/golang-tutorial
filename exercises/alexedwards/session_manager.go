package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	"io"
	"log"
	"net/http"
)

var session *scs.Session

func putHandler(w http.ResponseWriter, r *http.Request) {
	// Store a new key and value in the session data.
	session.Put(r.Context(), "message", "Hello from a session!")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Use the GetString helper to retrieve the string value associated with a
	// key. The zero value is returned if the key does not exist.
	msg := session.GetString(r.Context(), "message")
	io.WriteString(w, msg)
}

func runSessionManager() {
	/*
	https://godoc.org/github.com/alexedwards/scs
	We will use https://github.com/alexedwards/scs
	Its design leverages Go’s context package to automatically load and save session data via middleware.
	Importantly, it also provides the security features that you need when using server-side session stores (like
	straightforward session token regeneration) and supports both absolute and inactivity timeouts. The session data
	is safe for concurrent use.
	SCS supports a variety of different session stores:
		boltstore	BoltDB based session store
		memstore	In-memory session store (default)
		mysqlstore	MySQL based session store
		postgresstore	PostgreSQL based session store
		redisstore	Redis based session store
	We will use memstore(the default one)
	*/

	// Initialize a new session manager and configure it to use memstore as
	// the session store.
	session = scs.NewSession()
	session.Store = memstore.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/put", putHandler)
	mux.HandleFunc("/get", getHandler)

	// Wrap your handlers with the LoadAndSave() middleware.
	log.Println("Listening on port 4000...")
	http.ListenAndServe(":4000", session.LoadAndSave(mux))
}