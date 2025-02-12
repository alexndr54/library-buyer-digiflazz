package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/alexndr54/library-buyer-digiflazz/helper"
	"github.com/alexndr54/library-buyer-digiflazz/model"
	"io"
	"net/http"
)

type Digiflazz struct {
	Username string
	Key      string
}

// NewDigiflazz Membuat object digiflazz baru, harus dilakukan sebelum melakukan request ke API digiflazz
// username: Username akun digiflazz anda
// key: ProductionKey ataupun DevelopmentKey akun digiflazz anda
func NewDigiflazz(username, key string) *Digiflazz {
	return &Digiflazz{
		Username: username,
		Key:      key,
	}

}

// CekSaldo Untuk melakukan cek saldo yang tersedia di akun anda,tanpa memberikan parameter apapun
func (d *Digiflazz) CekSaldo() (*model.CekSaldoResponse, error) {
	sign := helper.GenerateMD5Hash(d.Username + d.Key + "depo")
	requestBody, err := json.Marshal(model.CekSaldoRequest{
		Cmd:      "deposit",
		Username: d.Username,
		Sign:     sign,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://api.digiflazz.com/v1/cek-saldo", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response model.CekSaldoResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Data.Rc != "" {
		return nil, errors.New(response.Data.Message)
	}

	return &response, nil
}

// DaftarHarga Untuk mendapatkan daftar harga produk anda
func (d *Digiflazz) DaftarHarga() (*model.DaftarHargaResponse, error) {
	sign := helper.GenerateMD5Hash(d.Username + d.Key + "prepaid")
	requestBody, err := json.Marshal(model.DaftarHargaRequest{
		Cmd:      "prepaid",
		Username: d.Username,
		Sign:     sign,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://api.digiflazz.com/v1/price-list", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response model.DaftarHargaResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		var dherror model.DaftarHargaResponseError
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

// Topup untuk melakukan transaksi topup/pembelian
// buyerSkuCode: Kode produk yang akan dibeli
// customerNo: Nomor tujuan transaksi
// refId: ID transaksi yang digunakan untuk mengecek status transaksi.
func (d *Digiflazz) Topup(buyerSkuCode, customerNo, refId string) (*model.TopUpResponse, error) {
	sign := helper.GenerateMD5Hash(d.Username + d.Key + refId)
	requestBody, err := json.Marshal(model.TopUpRequest{
		Username:     d.Username,
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
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response model.TopUpResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Data.Rc != "" {
		return nil, errors.New(response.Data.Message)
	}

	return &response, nil
}
