package usecase_test

import (
	"log"
	"multilingual_gurunavi_api/config"
	"multilingual_gurunavi_api/infrastructure/persistence"
	"multilingual_gurunavi_api/usecase"
	"net/http"
	"os"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// 環境変数読み込み
	EnvLoad()
	code := m.Run()
	os.Exit(code)
}

func EnvLoad() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	config.GNAVIURL = os.Getenv("GNAVI_API_URL")
	config.GNAVIID = os.Getenv("KEY_ID")
}

func TestGetStores(t *testing.T) {
	tests := []struct {
		testCase     string
		langs        []string
		wantSliceLen int
		wantErr      bool
	}{
		{
			testCase:     "指定の言語が１つの場合、10個の要素を返すこと",
			langs:        []string{"ja"},
			wantSliceLen: 10,
			wantErr:      false,
		},
		{
			testCase:     "指定の言語が5つの場合、50個の要素を返すこと",
			langs:        []string{"ja", "zh_cn", "zh_tw", "ko", "en"},
			wantSliceLen: 50,
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			// go-vcr のレコーダを生成
			// 保存済みの通信内容からモック化
			r, _ := recorder.New("../utils/test_data/gnavi_data_01")
			defer r.Stop()

			customHTTPClient := &http.Client{
				Transport: r,
			}

			sr := persistence.NewStorePersistence(customHTTPClient)
			su := usecase.NewStoreUsecase(sr)
			res, err := su.GetStores(tt.langs)
			if (err != nil) != tt.wantErr {
				t.Errorf("storeUsecase.GetStores() error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(res) != tt.wantSliceLen {
				t.Errorf("storeUsecase.GetStores() = %v, want %v", len(res), tt.wantSliceLen)
			}
		})
	}
}
