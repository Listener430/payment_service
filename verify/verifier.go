package verify

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type VReq struct {
	ReceiptData string `json:"receipt-data"`
}

type VResp struct {
	Status int `json:"status"`
}

type Verifier interface {
	Check(receipt string) bool
}

type verifier struct {
	url string
}

func NewVerifier(u string) Verifier {
	return &verifier{url: u}
}

func (v *verifier) Check(rd string) bool {
	b, _ := json.Marshal(VReq{ReceiptData: rd})
	res, err := http.Post(v.url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return false
	}
	defer res.Body.Close()
	var r VResp
	json.NewDecoder(res.Body).Decode(&r)
	return r.Status == 0
}
