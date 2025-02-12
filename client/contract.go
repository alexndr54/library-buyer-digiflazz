package client

import "github.com/alexndr54/library-buyer-digiflazz/model"

// Kontrak yang harus diikuti oleh semua request ke Digiflazz
type DigiflazzRequest interface {
	CekSaldo() (*model.CekSaldoResponse, error)
	DaftarHarga() (*model.DaftarHargaResponse, error)
	Topup(buyerSkuCode, customerNo, refId string) (*model.TopUpResponse, error)
}
