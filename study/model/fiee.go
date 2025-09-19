package model

type ArtistAccount struct {
	Name    string            `json:"name"`
	SubNum  string            `json:"subNum"`
	Account map[uint64]string `json:"account"`
}
