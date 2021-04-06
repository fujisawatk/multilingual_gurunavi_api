package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func StoresGetHandler(w http.ResponseWriter, r *http.Request) {
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
		// 対応する言語かどうか判別するバリデーション
		err := Validation(l)
		if err != nil {
			return []response{}, err
		}

		res, err := GnaviRequest(l)

		// レスポンス整形
		for _, r := range res.Rest {
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

type request struct {
	Langs []string `json:"langs"`
}

type response struct {
	Lang string `json:"lang"`
	Name string `json:"name"`
}
