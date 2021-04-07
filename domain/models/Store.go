package models

type Store struct {
	Rest []rest `json:"rest"`
}

type rest struct {
	Name name `json:"name"`
}

type name struct {
	Name string `json:"name"`
}
