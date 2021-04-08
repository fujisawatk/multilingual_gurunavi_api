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

func (sp *storePersistence) GnaviRequestCh(l string, responsesCh chan []models.Response) {
	url := config.GNAVIURL + "?keyid=" + config.GNAVIID + "&lang=" + l

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		responsesCh <- []models.Response{}
		return
	}

	resp, err := sp.Do(req)
	if err != nil {
		responsesCh <- []models.Response{}
		return
	}

	// 構造体に変換
	stores := &models.Store{}
	decodeBody(resp, stores)

	// レスポンス整形
	responses := make([]models.Response, 0, len(stores.Rest))
	for i := 0; i < len(stores.Rest); i++ {
		tmp := models.Response{
			Lang: l,
			Name: stores.Rest[i].Name.Name,
		}
		responses = append(responses, tmp)
	}
	responsesCh <- responses
}

// decodeBody 外部APIのレスポンスをデコード
func decodeBody(resp *http.Response, out *models.Store) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
