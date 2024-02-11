package models

type Transaction struct {
	Id     int     `db:"Id" json:"id"`
	Text   string  `db:"Text" json:"text"`
	Amount float64 `db:"Amount" json:"amount"`
}
