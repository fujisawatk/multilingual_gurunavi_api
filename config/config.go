package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// グローバル変数
var (
	GNAVIURL = ""
	GNAVIID  = ""
)

// EnvLoad 環境変数の読み込み
func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	GNAVIURL = os.Getenv("GNAVI_API_URL")
	GNAVIID = os.Getenv("KEY_ID")
}
