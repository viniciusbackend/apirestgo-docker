package models

type Produto struct {
	Id         int     `json:"id"`
	Nome       string  `json:"nome"`
	Tipo       string  `json:"tipo"`
	Valor      float64 `json:"valor"`
	Quantidade int     `json:"quantidade"`
}

var Produtos []Produto
