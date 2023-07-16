package usecase

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/doffy007/rest-api-makanan/config"
	"github.com/doffy007/rest-api-makanan/internal/entity"
	"github.com/doffy007/rest-api-makanan/internal/repository"
	"github.com/doffy007/rest-api-makanan/internal/request"
	"github.com/doffy007/rest-api-makanan/internal/response"
)

type ReceiptUsecase interface {
	ReceiptRegister(request.ParamReceiptRegister) (bool, response.BaseResponse)
	ReceiptFindOne(request.ParamFindReceipt) (bool, response.BaseResponse)
	ReceiptUpdate(request.ParamUpdateReceipt) (bool, response.BaseResponse)
	ReceiptGetAll() (bool, response.BaseResponse)
	ReceiptDelete(request.ParamDeleteReceipt) (bool, response.BaseResponse)
}

type receiptUsecase struct {
	ctx         context.Context
	conf        *config.Config
	receiptRepo repository.ReceiptRepository
}

// ReceiptGetAll implements.
func (u receiptUsecase) ReceiptGetAll() (bool, response.BaseResponse) {
	err := u.receiptRepo.GetAllReceipt()
	if err != nil {
		return false, response.BaseResponse{}.InternalServerError(err.Error())
	}

	return true, response.BaseResponse{StatusCode: http.StatusCreated, Message: "success Get All receipt", Data: true}
}

// ReceiptFindOne implements ReceiptUsecase.
func (u receiptUsecase) ReceiptFindOne(params request.ParamFindReceipt) (bool, response.BaseResponse) {
	var filter entity.FindOneReceipt

	if params.KategoriMakanan == "" {
		filter.Kategori = params.KategoriMakanan
	} else if params.ResepMakanan == "" {
		filter.ResepMakanan = params.ResepMakanan
	}

	data, err := u.receiptRepo.FindOneReceipt(filter, []string{})
	if err != nil {
		return false, response.BaseResponse{}.InternalServerError(err.Error())
	}

	if data == nil {
		return false, response.BaseResponse{StatusCode: http.StatusUnauthorized, Message: "username or password is incorrect"}
	}

	return true, response.BaseResponse{StatusCode: http.StatusAccepted, Message: "Success load data", Data: data}
}

// ReceiptRegister implements ReceiptUsecase.
func (u receiptUsecase) ReceiptRegister(params request.ParamReceiptRegister) (bool, response.BaseResponse) {
	if params.ResepMakanan != "" {
		checkResep, err := u.receiptRepo.FindOneReceipt(entity.FindOneReceipt{ResepMakanan: params.ResepMakanan}, []string{})
		if err != nil {
			return false, response.BaseResponse{}.InternalServerError(err.Error())
		}

		if checkResep != nil {
			return false, response.BaseResponse{StatusCode: http.StatusUnprocessableEntity, Message: "username already used"}

		}
	}

	if params.KategoriMakanan != "" {
		checkKategori, err := u.receiptRepo.FindOneReceipt(entity.FindOneReceipt{Kategori: params.KategoriMakanan}, []string{})
		if err != nil {
			return false, response.BaseResponse{}.InternalServerError(err.Error())
		}

		if checkKategori != nil {
			return false, response.BaseResponse{StatusCode: http.StatusUnprocessableEntity, Message: "username already used"}

		}
	}

	var payload entity.CreateReceipt
	rec, _ := json.Marshal(params)
	json.Unmarshal(rec, &payload)

	err := u.receiptRepo.CreateReceipt(payload)
	if err != nil {
		return false, response.BaseResponse{}.InternalServerError(err.Error())
	}

	return true, response.BaseResponse{StatusCode: http.StatusCreated, Message: "Success create Receipt"}
}

// ReceiptUpdate implements ReceiptUsecase.
func (u receiptUsecase) ReceiptUpdate(params request.ParamUpdateReceipt) (bool, response.BaseResponse) {
	findData, _ := u.receiptRepo.FindOneReceipt(entity.FindOneReceipt{
		ResepMakanan: params.ResepMakanan,
		Kategori:     params.KategoriMakanan,
	}, []string{})

	if params.KategoriMakanan != findData.KategoriMakanan {
		return false, response.BaseResponse{StatusCode: http.StatusUnprocessableEntity, Message: "Kategori tidak Ada"}
	}

	if params.ResepMakanan != findData.ResepMakanan {
		return false, response.BaseResponse{StatusCode: http.StatusUnprocessableEntity, Message: "Bahan tidak Ada"}
	}

	var payload entity.UpdateReceipt
	rec, _ := json.Marshal(params)
	json.Unmarshal(rec, &payload)

	err := u.receiptRepo.UpdateReceipt(entity.CreateReceipt{
		ResepMakanan:    payload.ResepMakanan,
		BahanMakanan:    []string{payload.BahanMakanan},
		KategoriMakanan: payload.Kategori,
	}, entity.FindOneReceipt{
		ResepMakanan: params.ResepMakanan,
		Kategori:     params.KategoriMakanan,
	})

	if err != nil {
		return false, response.BaseResponse{}.InternalServerError(err.Error())
	}

	return true, response.BaseResponse{StatusCode: http.StatusCreated, Message: "success update receipt", Data: true}
}

// ReceiptDelete implements.
func (u receiptUsecase) ReceiptDelete(params request.ParamDeleteReceipt) (bool, response.BaseResponse) {
	if params.ResepMakanan == "" {
		return false, response.BaseResponse{}.InternalServerError("Resep Makanan harus di isi")
	}

	var payload entity.DeleteReceipt
	rec, _ := json.Marshal(params)
	json.Unmarshal(rec, &payload)

	err := u.receiptRepo.DeleteReceipt(payload)
	if err != nil {
		return false, response.BaseResponse{}.InternalServerError(err.Error())
	}
	return true, response.BaseResponse{StatusCode: http.StatusCreated, Message: "Success create Receipt"}
}
func (u usecase) ReceiptUsecase() ReceiptUsecase {
	return &receiptUsecase{
		ctx:         u.ctx,
		conf:        u.config,
		receiptRepo: u.repository.ReceiptRepository(),
	}
}
