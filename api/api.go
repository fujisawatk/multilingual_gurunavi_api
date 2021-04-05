package api

import (
	"fmt"
	"io/ioutil"
	"multilingual_gurunavi_api/config"
	"net/http"
)

func HandleRestsGet(w http.ResponseWriter, r *http.Request) {
	url := config.GNAVIURL + "?keyid=" + config.GNAVIID + "&lang=" + "en"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprint(w, err)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprint(w, err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(w, err)
	}

	fmt.Fprint(w, string(byteArray))
}
