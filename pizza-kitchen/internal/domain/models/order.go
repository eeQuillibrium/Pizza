package models

type Order struct {
	UserId int64          `json:"userid"`
	Price  int64          `json:"price"`
	Units  []PieceUnitnum `json:"units"`
}

type PieceUnitnum struct {
	Piece   int64 `json:"piece"`
	Unitnum int64 `json:"unitnum"`
}
