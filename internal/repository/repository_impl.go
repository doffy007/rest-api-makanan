package repository

type Repository interface {
	ReceiptRepository() ReceiptRepository
}
