package models

type Order struct {
	UserId int            `json:"userid"`
	Price  int            `json:"price"`
	Units  []PieceUnitnum `json:"units"`
}

type PieceUnitnum struct {
	Piece   int `json:"piece"`
	Unitnum int `json:"unitnum"`
}
