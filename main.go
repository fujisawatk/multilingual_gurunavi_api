package main

import (
	"fmt"
	"multilingual_gurunavi_api/config"
	"net/http"
)

func main() {
	config.EnvLoad()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, config.GNAVIURL, "\n", config.GNAVIID)
}
