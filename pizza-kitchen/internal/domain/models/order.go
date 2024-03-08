package models

type Order struct {
	UserId int64
	Price  int64
	Units  []*PieceUnitnum
}

type PieceUnitnum struct {
	Piece   int64
	Unitnum int64
}
