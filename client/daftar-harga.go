package client

import (
	"bytes"
	"digiflazz/helper"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type DaftarHargaRequest struct {
	Cmd      string `json:"cmd"`
	Username string `json:"username"`
	Sign     string `json:"sign"`
}

type DaftarHargaResponse struct {
	Error DaftarHargaResponseError `json:"error"`
	Data  []struct {
		ProductName         string `json:"product_name"`
		Category            string `json:"category"`
		Brand               string `json:"brand"`
		Type                string `json:"type"`
		SellerName          string `json:"seller_name"`
		Price               int    `json:"price"`
		BuyerSkuCode        string `json:"buyer_sku_code"`
		BuyerProductStatus  bool   `json:"buyer_product_status"`
		SellerProductStatus bool   `json:"seller_product_status"`
		UnlimitedStock      bool   `json:"unlimited_stock"`
		Stock               int    `json:"stock"`
		Multi               bool   `json:"multi"`
		StartCutOff         string `json:"start_cut_off"`
		EndCutOff           string `json:"end_cut_off"`
		Desc                string `json:"desc"`
	} `json:"data"`
}

type DaftarHargaResponseError struct {
	Data struct {
		Rc      string `json:"rc"`
		Message string `json:"message"`
	} `json:"data"`
}

func DaftarHarga(username, apiKey string) (*DaftarHargaResponse, error) {
	sign := helper.GenerateMD5Hash(username + apiKey + "prepaid")
	requestBody, err := json.Marshal(DaftarHargaRequest{
		Cmd:      "prepaid",
		Username: username,
		Sign:     sign,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://api.digiflazz.com/v1/price-list", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response DaftarHargaResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		var dherror DaftarHargaResponseError
		err = json.Unmarshal(body, &dherror)
		if err != nil {
			return nil, err
		}

		response.Error = dherror
	}

	if response.Error.Data.Rc != "" {
		return nil, errors.New(response.Error.Data.Message)
	}

	return &response, nil
}
