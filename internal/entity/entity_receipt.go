package entity

import "time"

type Receipt struct {
	ID              int        `json:"id" db:"id"`
	ResepMakanan    string     `json:"resep" db:"resep"`
	BahanMakanan    []string   `json:"bahan" db:"bahan"`
	KategoriMakanan string     `json:"kategori" db:"kategori"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updated_at"`
}

type CreateReceipt struct {
	ResepMakanan    string     `json:"resep" db:"resep"`
	BahanMakanan    []string   `json:"bahan" db:"bahan"`
	KategoriMakanan string     `json:"kategori" db:"kategori"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updated_at"`
}

type FindOneReceipt struct {
	ID           int        `json:"id" db:"id"`
	ResepMakanan string     `json:"resep" db:"resep"`
	BahanMakanan []string   `json:"bahan" db:"bahan"`
	Kategori     string     `json:"kategori" db:"kategori"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
}

type UpdateReceipt struct {
	ID           int    `json:"id" db:"id"`
	ResepMakanan string `json:"resep" db:"resep"`
	BahanMakanan string `json:"bahan" db:"bahan"`
	Kategori     string `json:"kategori" db:"kategori"`
}

type DeleteReceipt struct {
	ResepMakanan string `json:"resep" db:"resep"`
}
