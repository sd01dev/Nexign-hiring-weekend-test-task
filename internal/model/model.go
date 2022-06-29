package model

// SpellerApiResponse Yandex speller return exactly this struct
type SpellerApiResponse []struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

type TextToCheck struct {
	Texts []string `json:"Texts"`
}
