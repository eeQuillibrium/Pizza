package models

type Order struct {
	OrderId int            `json:"orderid"`
	UserId  int            `json:"userid"`
	Price   int            `json:"price"`
	State   string         `json:"state"`
	Units   []PieceUnitnum `json:"units"`
}

type PieceUnitnum struct {
	Piece   int `json:"piece"`
	Unitnum int `json:"unitnum"`
}
