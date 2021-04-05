package main

import (
	"multilingual_gurunavi_api/api"
	"multilingual_gurunavi_api/config"
	"net/http"
)

func main() {
	config.EnvLoad()
	http.HandleFunc("/", api.HandleRestsGet)
	http.ListenAndServe(":8000", nil)
}
