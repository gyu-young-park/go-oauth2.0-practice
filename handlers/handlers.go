package handlers

import (
	"net/http"
)

func serveStaticFile(mux *http.ServeMux) {
	mux.Handle("/", http.FileServer(http.Dir("public")))
}

func serveAuthHanlder(mux *http.ServeMux) {
	mux.HandleFunc("/api/auth/google/login", oauthGoogleLoginHandler)
	mux.HandleFunc("/api/auth/google/callback", oauthGoogleCallbackHandler)
}

func ServeHandler(mux *http.ServeMux) {
	serveStaticFile(mux)
	serveAuthHanlder(mux)
}
