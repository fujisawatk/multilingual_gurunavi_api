package persistence_test

import (
	"log"
	"multilingual_gurunavi_api/config"
	"multilingual_gurunavi_api/infrastructure/persistence"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// 環境変数読み込み
	EnvLoad()
	code := m.Run()
	os.Exit(code)
}

func EnvLoad() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	config.GNAVIURL = os.Getenv("GNAVI_API_URL")
	config.GNAVIID = os.Getenv("KEY_ID")
}

func TestGnaviRequest(t *testing.T) {
	tests := []struct {
		testCase     string
		lang         string
		wantSliceLen int
		wantErr      bool
	}{
		{
			testCase:     "ぐるなびの多言語版レストラン検索から指定の言語に関連する店舗情報を取得できること",
			lang:         "ja",
			wantSliceLen: 10,
			wantErr:      false,
		},
	}
	sr := persistence.NewStorePersistence(http.DefaultClient)

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			res, err := sr.GnaviRequest(tt.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("storePersistence.GnaviRequest() error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(res.Rest) != tt.wantSliceLen {
				t.Errorf("storePersistence.GnaviRequest() = %v, want %v", len(res.Rest), tt.wantSliceLen)
			}
		})
	}
}
