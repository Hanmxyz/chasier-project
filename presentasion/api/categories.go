package api

import (
	"chasier/domain/categories"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type CategoriesHandler struct {
	Usecase categories.CategoryUsecase
}

func (handler CategoriesHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := handler.Usecase.GetAllCategories()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	categoriesbyte, err := json.Marshal(categories)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(categoriesbyte)
}

func (handler CategoriesHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category categories.Category
	bodybyte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(bodybyte, &category)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = handler.Usecase.CreateCategory(category)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("success : true"))
}

func (handler CategoriesHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var category categories.Category

	bodybyte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(bodybyte, &category)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = handler.Usecase.UpdateCategory(category, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("success: true"))
}

func (handler CategoriesHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := handler.Usecase.DeleteCategory(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("success: true"))
}
