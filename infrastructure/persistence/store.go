package persistence

import (
	"encoding/json"
	"multilingual_gurunavi_api/config"
	"multilingual_gurunavi_api/domain/models"
	"multilingual_gurunavi_api/domain/repository"
	"net/http"
)

type storePersistence struct {
	*http.Client
}

func NewStorePersistence(c *http.Client) repository.StoreRepository {
	return &storePersistence{c}
}

// GnaviRequest ぐるなびAPIからデータ取得
func (sp *storePersistence) GnaviRequest(l string) (*models.Store, error) {
	url := config.GNAVIURL + "?keyid=" + config.GNAVIID + "&lang=" + l

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &models.Store{}, err
	}

	resp, err := sp.Do(req)
	if err != nil {
		return &models.Store{}, err
	}

	// 構造体に変換
	stores := &models.Store{}
	decodeBody(resp, stores)

	return stores, nil
}

// decodeBody 外部APIのレスポンスをデコード
func decodeBody(resp *http.Response, out *models.Store) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
