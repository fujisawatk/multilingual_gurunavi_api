package handler

import (
	"encoding/json"
	"io/ioutil"
	"multilingual_gurunavi_api/interface/responses"
	"multilingual_gurunavi_api/usecase"
	"net/http"
)

type StoreHandler interface {
	StoresGetHandler(w http.ResponseWriter, r *http.Request)
}

type storeHandler struct {
	storeUsecase usecase.StoreUsecase
}

func NewStoreHandler(su usecase.StoreUsecase) StoreHandler {
	return &storeHandler{
		storeUsecase: su,
	}
}

func (sh *storeHandler) StoresGetHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストをパース
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var request request
	err = json.Unmarshal(body, &request)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// res, err := sh.storeUsecase.GetStores(request.Langs)
	res, err := sh.storeUsecase.GetStoresByWaitGroup(request.Langs)
	// res, err := sh.storeUsecase.GetStoresByChRange(request.Langs)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, res)
}

type request struct {
	Langs []string `json:"langs"`
}
