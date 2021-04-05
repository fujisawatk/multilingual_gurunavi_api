package api

import (
	"encoding/json"
	"fmt"
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

	// 構造体に変換
	var storeItems storeItems
	decodeBody(resp, &storeItems)

	fmt.Fprint(w, storeItems)
}

// decodeBody 外部APIのレスポンスをデコード
func decodeBody(resp *http.Response, out *storeItems) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

type storeItems struct {
	Rest []rest `json:"rest"`
}

type rest struct {
	Name name `json:"name"`
}

type name struct {
	Name string `json:"name"`
}
