package product

import "chasier/domain/categories"

type Product struct {
	Id          int
	Nama        string `json:"name"`
	Price       int
	Category_Id int
	Category    categories.Category
}
