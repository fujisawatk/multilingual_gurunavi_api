package api

import "errors"

func Validation(req *string) error {
	var langs = []string{"ja", "zh_cn", "zh_tw", "ko", "en"}
	for _, l := range langs {
		if (*req) == l {
			return nil
		}
	}
	return errors.New("not supported languages")
}
