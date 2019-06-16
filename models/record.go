package models

type Record struct {
	Items []Item `json:"items"`
}

type Item struct {
	Domain   string `json:"domain"`
	Sslgrade string `json:"sslgrade"`
}
