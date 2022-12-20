package main

import (
	"fmt"
	"net/http"

	"github.com/gyu-young-park/go-oauth2.0-practice/config"
	"github.com/gyu-young-park/go-oauth2.0-practice/handlers"
)

func main() {
	config.SetupConfig()

	mux := http.NewServeMux()
	handlers.ServeHandler(mux)
	fmt.Println("Start server:", config.ConfigMap.Port)
	http.ListenAndServe(config.ConfigMap.Port, mux)
}
