package main

import (
	"chasier/data/db"
	"chasier/domain/categories"
	"chasier/domain/product"
	"chasier/domain/transaction"
	"chasier/presentasion/api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	DB, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/chasier?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err.Error())
	}

	CategoriesRepository := db.NewCategoryRepository(DB)
	CategoriesUsecase := categories.CategoryUsecase{Repository: CategoriesRepository}
	CategoriesHandler := api.CategoriesHandler{Usecase: CategoriesUsecase}

	ProductRepository := db.NewProductRepository(DB)
	ProductUsecase := product.ProductUsecase{Repository: ProductRepository}
	productHandler := api.ProductHandler{Usecase: ProductUsecase}

	TransactionRepository := db.NewTransactionRepository(DB)
	TransactionUsecase := transaction.TransactionUsecase{Repository: TransactionRepository}
	TransactionHandler := api.TransactionHandler{Usecase: TransactionUsecase}

	router := mux.NewRouter()

	router.HandleFunc("/categories", CategoriesHandler.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories", CategoriesHandler.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", CategoriesHandler.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", CategoriesHandler.DeleteCategory).Methods("DELETE")

	router.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.GetProductById).Methods("GET")
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/p_bycategory/{id}", productHandler.GetProductByCategoryId).Methods("GET")

	router.HandleFunc("/transactions", TransactionHandler.GetAllTransactions).Methods("GET")
	router.HandleFunc("/transactions", TransactionHandler.CreateTransaction).Methods("POST")

	err = http.ListenAndServe(":1010", router)
	fmt.Println(err)

}
