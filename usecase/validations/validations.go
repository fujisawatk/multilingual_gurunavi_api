package validations

import "errors"

// func LangCheck(req *string) error {
// 	var langs = []string{"ja", "zh_cn", "zh_tw", "ko", "en"}
// 	for _, l := range langs {
// 		if (*req) == l {
// 			return nil
// 		}
// 	}
// 	return errors.New("not supported languages")
// }

func LangCheck(req *string) error {
	var langs = []string{"ja", "zh_cn", "zh_tw", "ko", "en"}
	for i := 0; i < len(langs); i++ {
		if (*req) == langs[i] {
			return nil
		}
	}
	return errors.New("not supported languages")
}
