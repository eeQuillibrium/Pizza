package models

type Order struct {
	UserId int            `json:"userid"`
	Price  float64        `json:"price"`
	Units  []PieceUnitnum `json:"units"`
}

type PieceUnitnum struct {
	Piece   int `json:"piece"`
	Unitnum int `json:"unitnum"`
}
