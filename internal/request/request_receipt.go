package request

type ParamReceiptRegister struct {
	ResepMakanan    string   `json:"resep" `
	BahanMakanan    []string `json:"bahan" `
	KategoriMakanan string   `json:"kategori" `
}

type ParamFindReceipt struct {
	ResepMakanan    string `json:"resep" `
	KategoriMakanan string `json:"kategori" `
}

type ParamUpdateReceipt struct {
	ResepMakanan    string   `json:"resep" `
	BahanMakanan    []string `json:"bahan" `
	KategoriMakanan string   `json:"kategori" `
}

type ParamDeleteReceipt struct {
	ResepMakanan string `json:"resep" `
}
