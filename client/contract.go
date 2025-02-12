package client

import "github.com/alexndr54/library-buyer-digiflazz/model"

type DigiflazzRequest interface {
	CekSaldo() (*model.CekSaldoResponse, error)
	DaftarHarga() (*model.DaftarHargaResponse, error)
	Topup(buyerSkuCode, customerNo, refId string) (*model.TopUpResponse, error)
}
