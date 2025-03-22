package api

import (
	"chasier/domain/transaction"
	"encoding/json"
	"io/ioutil"

	"net/http"
)

type TransactionHandler struct {
	Usecase transaction.TransactionUsecase
}

func (handler TransactionHandler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transaction, err := handler.Usecase.GetAllTransactions()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	TransactionByte, err := json.Marshal(transaction)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(TransactionByte)
}

func (handler TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	BodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var ReqByte transaction.Request
	err = json.Unmarshal(BodyByte, &ReqByte)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = handler.Usecase.CreateTransaction(ReqByte)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("success"))
}
