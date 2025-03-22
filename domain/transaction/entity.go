package transaction

import "time"

type Transaction struct {
	Id                int
	CreatedAt         time.Time
	Transaction_items []Transaction_items
}

type Transaction_items struct {
	Id            int
	Product_Id    int
	Product       Product
	Amount        int
	TransactionId int `json:"-"`
}

type Product struct {
	Id    int
	Nama  string `json:"name"`
	Price int
}

type Request struct {
	Items []struct {
		Product_Id int
		Amount     int
	}
}
