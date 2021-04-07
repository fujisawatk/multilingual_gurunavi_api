package usecase

import (
	"errors"
	"multilingual_gurunavi_api/domain/repository"
	"multilingual_gurunavi_api/usecase/validations"
)

type StoreUsecase interface {
	GetStores([]string) ([]response, error)
}

type storeUsecase struct {
	storeRepository repository.StoreRepository
}

func NewStoreUsecase(sr repository.StoreRepository) StoreUsecase {
	return &storeUsecase{
		storeRepository: sr,
	}
}

// GetStores ぐるなびAPIからデータ取得〜整形まで
func (su *storeUsecase) GetStores(langs []string) ([]response, error) {
	var responses []response

	// 各言語ごとにぐるなびAPIにリクエストを出す
	for _, l := range langs {
		err := validations.LangCheck(&l)
		if err != nil {
			// 非対応言語は処理スキップ
			continue
		}

		res, err := su.storeRepository.GnaviRequest(l)

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

type response struct {
	Lang string `json:"lang"`
	Name string `json:"name"`
}
