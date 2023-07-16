package repository

import (
	"context"
	"database/sql"
	"fmt"

	"strings"

	"github.com/doffy007/rest-api-makanan/config"
	"github.com/doffy007/rest-api-makanan/database/query"
	"github.com/doffy007/rest-api-makanan/internal/entity"
	"github.com/jmoiron/sqlx"
)

type ReceiptRepository interface {
	CreateReceipt(entity.CreateReceipt) error
	FindOneReceipt(entity.FindOneReceipt, []string) (*entity.Receipt, error)
	UpdateReceipt(entity.CreateReceipt, entity.FindOneReceipt) error
	GetAllReceipt() error
	DeleteReceipt(entity.DeleteReceipt) error
}

type receiptRepository struct {
	ctx  context.Context
	conf *config.Config
	db   *sqlx.DB
}

func (r repository) ReceiptRepository() ReceiptRepository {
	return &receiptRepository{
		ctx:  r.ctx,
		conf: r.conf,
		db:   r.db,
	}
}

// CreateReceipt implements.
func (r receiptRepository) CreateReceipt(payload entity.CreateReceipt) error {
	_, err := r.db.NamedExecContext(r.ctx, query.CreateReceipt, payload)
	if err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Repository CreateReceipt: %v\n", err.Error())
		return err
	}

	return nil
}

// FindOneReceipt implements.
func (r receiptRepository) FindOneReceipt(payload entity.FindOneReceipt, fields []string) (*entity.Receipt, error) {
	var columns string
	if len(fields) == 0 {
		columns = "*"
	} else {
		columns = strings.Join(fields, ", ")
	}

	dest := entity.Receipt{}

	var filterValues []interface{}
	var filterKeys []string

	if payload.ID == 0 {
		filterKeys = append(filterKeys, "id = ?")
		filterValues = append(filterValues, string(rune(payload.ID)))
	}

	if payload.ResepMakanan != "" {
		filterKeys = append(filterKeys, "resep = ?")
		filterValues = append(filterValues, payload.ResepMakanan)
	}

	if payload.Kategori != "" {
		filterKeys = append(filterKeys, "kategori = ?")
		filterValues = append(filterValues, payload.ResepMakanan)
	}

	query := fmt.Sprintf(query.FindOneReceipt, columns, strings.Join(filterKeys, " AND "))
	err := r.db.GetContext(r.ctx, dest, query, filterValues...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Printf("\033[1;31m [ERROR] \033[0m Repository FindOneReceipt: %v\n", err.Error())
		return nil, err
	}

	return &dest, nil
}

// UpdateReceipt implements ReceiptRepository.
func (r receiptRepository) UpdateReceipt(payload entity.CreateReceipt, filter entity.FindOneReceipt) error {
	var fieldKeys []string
	var fieldAndFilterValues []interface{}

	var filterKeys []string

	// Payload
	if payload.ResepMakanan == "" {
		fieldKeys = append(fieldKeys, "resep = ?")
		fieldAndFilterValues = append(fieldAndFilterValues, payload.ResepMakanan)
	}

	if payload.KategoriMakanan == "" {
		fieldKeys = append(fieldKeys, "kategori = ?")
		fieldAndFilterValues = append(fieldAndFilterValues, payload.KategoriMakanan)
	}

	if payload.BahanMakanan != nil {
		fieldKeys = append(fieldKeys, "bahan = ?")
		fieldAndFilterValues = append(fieldAndFilterValues, payload.BahanMakanan)
	}

	// Filter
	if filter.ID == 0 {
		fieldKeys = append(fieldKeys, "id = ?")
		fieldAndFilterValues = append(fieldAndFilterValues, filter.ID)
	}
	if filter.ResepMakanan == "" {
		fieldKeys = append(fieldKeys, "resep = ?")
		fieldAndFilterValues = append(fieldAndFilterValues, filter.ResepMakanan)
	}
	if filter.BahanMakanan != nil {
		filterKeys = append(filterKeys, "bahan = ?")
		fieldAndFilterValues = append(fieldAndFilterValues, filter.BahanMakanan)
	}
	if filter.Kategori == "" {
		fieldKeys = append(fieldKeys, "kategori = ?")
		fieldAndFilterValues = append(fieldAndFilterValues, filter.Kategori)
	}

	queryUpdate := fmt.Sprintf(query.UpdateReceipt, strings.Join(fieldKeys, ","), strings.Join(filterKeys, " AND "))

	_, err := r.db.ExecContext(r.ctx, queryUpdate, fieldAndFilterValues...)
	if err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Repository UpdateReceipt: %v\n", err.Error())
		return err
	}

	return nil
}

// GetAllReceipt implements.
func (r receiptRepository) GetAllReceipt() error {

	dest := entity.Receipt{}
	err := r.db.GetContext(r.ctx, dest, query.FindAllReceipt)
	if err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Repository GetAllReceipt: %v\n", err.Error())
		return err
	}

	return nil
}

// DeleteReceipt implements.
func (r receiptRepository) DeleteReceipt(payload entity.DeleteReceipt) error {
	_, err := r.db.NamedExecContext(r.ctx, query.DeleteReceipt, payload)
	if err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Repository DeleteReceipt: %v\n", err.Error())
		return err
	}

	return nil
}
