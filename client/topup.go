package client

import (
	"bytes"
	"digiflazz/helper"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type TopUpRequest struct {
	Username     string `json:"username"`
	BuyerSkuCode string `json:"buyer_sku_code"`
	CustomerNo   string `json:"customer_no"`
	RefId        string `json:"ref_id"`
	Sign         string `json:"sign"`
}

type TopUpResponse struct {
	Data struct {
		RefId          string `json:"ref_id"`
		CustomerNo     string `json:"customer_no"`
		BuyerSkuCode   string `json:"buyer_sku_code"`
		Message        string `json:"message"`
		Status         string `json:"status"`
		Rc             string `json:"rc"`
		Sn             string `json:"sn"`
		BuyerLastSaldo int    `json:"buyer_last_saldo"`
		Price          int    `json:"price"`
		Tele           string `json:"tele"`
		Wa             string `json:"wa"`
	} `json:"data"`
}

func TopUp(username, apiKey, buyerSkuCode, customerNo, refId string) (*TopUpResponse, error) {
	sign := helper.GenerateMD5Hash(username + apiKey + refId)
	requestBody, err := json.Marshal(TopUpRequest{
		Username:     username,
		BuyerSkuCode: buyerSkuCode,
		CustomerNo:   customerNo,
		RefId:        refId,
		Sign:         sign,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://api.digiflazz.com/v1/transaction", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response TopUpResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Data.Rc != "" {
		return nil, errors.New(response.Data.Message)
	}

	return &response, nil
}
