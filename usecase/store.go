package usecase

import (
	"errors"
	"multilingual_gurunavi_api/domain/models"
	"multilingual_gurunavi_api/domain/repository"
	"multilingual_gurunavi_api/usecase/validations"
	"sync"
)

type StoreUsecase interface {
	GetStores([]string) ([]response, error)
	GetStoresByWaitGroup([]string) ([]models.Response, error)
	GetStoresByChRange([]string) ([]models.Response, error)
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
		if err != nil {
			continue
		}

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

// GetStoresByWaitGroup WaitGroupによるGoroutine終了待ちパターン
func (su *storeUsecase) GetStoresByWaitGroup(langs []string) ([]models.Response, error) {
	wg := new(sync.WaitGroup)
	// cpus := runtime.NumCPU() // CPUの数
	// semaphore := make(chan int, cpus)
	responses := make([]models.Response, 0, len(langs))

	for i := 0; i < len(langs); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// semaphore <- 1
			err := validations.LangCheck(&langs[i])
			if err != nil {
				return
			}
			res, err := su.storeRepository.GnaviRequest(langs[i])
			if err != nil {
				return
			}
			// レスポンス整形
			for _, r := range res.Rest {
				responses = append(responses,
					models.Response{
						Lang: langs[i],
						Name: r.Name.Name,
					},
				)
			}
			// <-semaphore
		}(i)
	}
	wg.Wait()
	// 取得レコードがない場合のエラーハンドリング
	if len(responses) > 0 {
		return responses, nil
	} else {
		return []models.Response{}, errors.New("record not found")
	}
}

// GetStoresByChRange 外部APIのレスポンスをチャネルに送信、rangeで受信パターン
func (su *storeUsecase) GetStoresByChRange(langs []string) ([]models.Response, error) {
	responses := make([]models.Response, 0, len(langs))
	responsesCh := make(chan []models.Response)

	go func() {
		defer close(responsesCh)
		for i := 0; i < len(langs); i++ {
			err := validations.LangCheck(&langs[i])
			if err != nil {
				// 非対応言語は処理スキップ
				responsesCh <- []models.Response{}
				continue
			}
			su.storeRepository.GnaviRequestCh(langs[i], responsesCh)
		}
	}()
	// チャネルに格納された変数を移動
	for response := range responsesCh {
		responses = append(responses, response...)
	}
	// 取得レコードがない場合のエラーハンドリング
	if len(responses) > 0 {
		return responses, nil
	} else {
		return []models.Response{}, errors.New("record not found")
	}
}

type response struct {
	Lang string `json:"lang"`
	Name string `json:"name"`
}
