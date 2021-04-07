package main

import (
	"multilingual_gurunavi_api/config"
	"multilingual_gurunavi_api/infrastructure/persistence"
	"multilingual_gurunavi_api/interface/handler"
	"multilingual_gurunavi_api/usecase"
	"net/http"
)

func main() {
	// 環境変数読み込み
	config.EnvLoad()
	// 依存関係注入
	storePersistence := persistence.NewStorePersistence(http.DefaultClient)
	storeUsecase := usecase.NewStoreUsecase(storePersistence)
	storeHandler := handler.NewStoreHandler(storeUsecase)
	// ルーティング
	http.HandleFunc("/", storeHandler.StoresGetHandler)
	// サーバ起動
	http.ListenAndServe(":8000", nil)
}
