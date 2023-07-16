package handler

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/doffy007/rest-api-makanan/internal/constants"
	"github.com/doffy007/rest-api-makanan/internal/helper"
	"github.com/doffy007/rest-api-makanan/internal/request"
	"github.com/doffy007/rest-api-makanan/internal/usecase"
)

type ReceiptHandler interface {
	ReceiptRegister(w http.ResponseWriter, r *http.Request)
	FindOneReceipt(w http.ResponseWriter, r *http.Request)
	UpdateReceipt(w http.ResponseWriter, r *http.Request)
	DeleteReceipt(w http.ResponseWriter, r *http.Request)
	GetAllReceipt(w http.ResponseWriter, r *http.Request)
}

type receiptHandler struct {
	usecase usecase.ReceiptUsecase
}

func (h handler) ReceiptHandler() ReceiptHandler {
	return receiptHandler{
		usecase: h.usecase.ReceiptUsecase(),
	}
}

// ReceiptRegister implements.
func (h receiptHandler) ReceiptRegister(w http.ResponseWriter, r *http.Request) {
	var params request.ParamReceiptRegister
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		helper.ResponseWriter(w, http.StatusUnprocessableEntity, constants.ERROR_VALIDATION, nil)
		return
	}

	ok, res := h.usecase.ReceiptRegister(params)
	if !ok {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler Receipt From Usecase: %v\n", res.Message)
		helper.ResponseWriter(w, http.StatusBadRequest, res.Message, nil)
		return
	}

	helper.ResponseWriter(w, http.StatusCreated, "", res.Data)
}

// FindOneReceipt implements.
func (h receiptHandler) FindOneReceipt(w http.ResponseWriter, r *http.Request) {
	var params request.ParamFindReceipt
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		helper.ResponseWriter(w, http.StatusUnprocessableEntity, constants.ERROR_VALIDATION, nil)
		return
	}

	ok, res := h.usecase.ReceiptFindOne(params)
	if !ok {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler Receipt From Usecase: %v\n", res.Message)
		helper.ResponseWriter(w, http.StatusBadRequest, res.Message, nil)
		return
	}

	helper.ResponseWriter(w, http.StatusAccepted, "", res.Data)
}

// UpdateReceipt implements ReceiptHandler.
func (h receiptHandler) UpdateReceipt(w http.ResponseWriter, r *http.Request) {
	var params request.ParamUpdateReceipt
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		helper.ResponseWriter(w, http.StatusUnprocessableEntity, constants.ERROR_VALIDATION, nil)
		return
	}

	ok, res := h.usecase.ReceiptUpdate(params)
	if !ok {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler Receipt From Usecase: %v\n", res.Message)
		helper.ResponseWriter(w, http.StatusBadRequest, res.Message, nil)
		return
	}

	helper.ResponseWriter(w, http.StatusAccepted, "", res.Data)
}

// DeleteReceipt implements ReceiptHandler.
func (h receiptHandler) DeleteReceipt(w http.ResponseWriter, r *http.Request) {
	var params request.ParamDeleteReceipt
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		helper.ResponseWriter(w, http.StatusUnprocessableEntity, constants.ERROR_VALIDATION, nil)
		return
	}

	ok, res := h.usecase.ReceiptDelete(params)
	if !ok {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler Receipt From Usecase: %v\n", res.Message)
		helper.ResponseWriter(w, http.StatusBadRequest, res.Message, nil)
		return
	}

	helper.ResponseWriter(w, http.StatusAccepted, "", nil)
}

// GetAllReceipt implements ReceiptHandler.
func (h receiptHandler) GetAllReceipt(w http.ResponseWriter, r *http.Request) {
	ok, res := h.usecase.ReceiptGetAll()
	if !ok {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler Receipt From Usecase: %v\n", res.Message)
		helper.ResponseWriter(w, http.StatusBadRequest, res.Message, nil)
		return
	}

	helper.ResponseWriter(w, http.StatusAccepted, "", res.Data)
}
