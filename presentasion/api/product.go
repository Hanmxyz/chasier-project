package api

import (
	"chasier/domain/product"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	Usecase product.ProductUsecase
}

func (handler ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := handler.Usecase.GetAllProducts()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	productsbyte, err := json.Marshal(products)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(productsbyte)
}

func (handler ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	Id := Vars["id"]

	product, err := handler.Usecase.GetProductById(Id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	productbyte, err := json.Marshal(product)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(productbyte)
}

func (handler ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	bodybyte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var product product.Product
	err = json.Unmarshal(bodybyte, &product)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = handler.Usecase.CreateProduct(product)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("success: true"))
}

func (handler ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	bodybyte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var product product.Product
	err = json.Unmarshal(bodybyte, &product)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = handler.Usecase.UpdateProduct(product, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("success: true"))
}

func (handler ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := handler.Usecase.DeleteProduct(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("success: true"))
}

func (handler ProductHandler) GetProductByCategoryId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	productsbyte, err := handler.Usecase.GetProductByCategoryId(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	bodybyte, err := json.Marshal(productsbyte)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(bodybyte)
}
