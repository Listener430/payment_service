package models

type PurchaseReq struct {
	ReceiptData string `json:"receiptData"`
	UserID      string `json:"userId"`
}
