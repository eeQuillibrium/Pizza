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

// unitNums is a string consist of order position number separated by ","
// amount is a string consist of order position units separated by ","
// order has amount[i] unitNums[i] 
type OrderDB struct {
	OrderId  int    `db:"order_id"`
	Price    int    `db:"price"`
	UnitNums string `db:"unit_nums"`
	Amount   string `db:"amount"`
	State    string `db:"state"`
	UserId   int    `db:"user_id"`
}