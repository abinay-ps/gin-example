package models

type Movie struct {
	ID       int32  `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
}
