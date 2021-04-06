package api

import (
	"multilingual_gurunavi_api/config"
	"net/http"
)

type storeItems struct {
	Rest []rest `json:"rest"`
}

type rest struct {
	Name name `json:"name"`
}

type name struct {
	Name string `json:"name"`
}

// GnaviRequest ぐるなびAPIからデータ取得
func GnaviRequest(l string) (storeItems, error) {
	url := config.GNAVIURL + "?keyid=" + config.GNAVIID + "&lang=" + l

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return storeItems{}, err
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return storeItems{}, err
	}

	// 構造体に変換
	var storeItems storeItems
	decodeBody(resp, &storeItems)

	return storeItems, nil
}
