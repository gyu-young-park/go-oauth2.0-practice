package main

import (
	"net/http"

	"github.com/gyu-young-park/go-oauth2.0-practice/config"
	"github.com/gyu-young-park/go-oauth2.0-practice/handlers.go"
)

func main() {
	config.SetupConfig()

	mux := http.NewServeMux()
	handlers.ServeHandler(mux)
	http.ListenAndServe(":8080", mux)
}
