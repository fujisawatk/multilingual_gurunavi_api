package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"multilingual_gurunavi_api/config"
	"net/http"
)

func HandleRestsGet(w http.ResponseWriter, r *http.Request) {
	// リクエストをパース
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}

	var request request
	err = json.Unmarshal(body, &request)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := GetStores(request)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, res)
}

// GetStores ぐるなびAPIからデータ取得〜整形まで
func GetStores(req request) ([]response, error) {
	var responses []response

	// 各言語ごとにぐるなびAPIにリクエストを出す
	for _, l := range req.Langs {
		url := config.GNAVIURL + "?keyid=" + config.GNAVIID + "&lang=" + l

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return []response{}, err
		}

		client := new(http.Client)
		resp, err := client.Do(req)
		if err != nil {
			return []response{}, err
		}
		defer resp.Body.Close()

		// 構造体に変換
		var storeItems storeItems
		decodeBody(resp, &storeItems)

		// レスポンス整形
		for _, r := range storeItems.Rest {
			tmp := response{
				Lang: l,
				Name: r.Name.Name,
			}
			responses = append(responses, tmp)
		}
	}
	// 取得レコードがない場合のエラーハンドリング
	if len(responses) > 0 {
		return responses, nil
	} else {
		return []response{}, errors.New("record not found")
	}
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

type request struct {
	Langs []string `json:"langs"`
}

type response struct {
	Lang string `json:"lang"`
	Name string `json:"name"`
}
