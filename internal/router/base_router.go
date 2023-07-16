package router

func (r router) BaseRouter() {
	receiptHandler := r.handler.ReceiptHandler()

	r.route.HandleFunc("/register", receiptHandler.ReceiptRegister).Methods("POST")
	r.route.HandleFunc("/find", receiptHandler.FindOneReceipt).Methods("GET")
	r.route.HandleFunc("/all", receiptHandler.GetAllReceipt).Methods("GET")
	r.route.HandleFunc("/delete", receiptHandler.DeleteReceipt).Methods("DELETE")
	r.route.HandleFunc("/update", receiptHandler.UpdateReceipt).Methods("PATCH")
}
