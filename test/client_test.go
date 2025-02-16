package test

import (
	"testing"

	"github.com/alexndr54/library-buyer-digiflazz/client"
	"github.com/stretchr/testify/assert"
)

var Digi = client.NewDigiflazz("your_username", "yourkey")

func TestCekSaldo(t *testing.T) {
	t.Run("Cek Saldo Berhasil", func(t *testing.T) {
		saldo, err := Digi.CekSaldo()
		assert.Nil(t, err)
		if err == nil {
			t.Log("Saldo: ", saldo.Data.Deposit)
		}
	})

	t.Run("Cek Saldo Gagal", func(t *testing.T) {
		Digi.Username = "ini_username_salah"
		_, err := Digi.CekSaldo()
		assert.Nil(t, err, "Gagal Pesan: "+err.Error())
	})
}
func TestDaftarHarga(t *testing.T) {
	t.Run("Daftar Harga Berhasil", func(t *testing.T) {

		daftarHarga, err := Digi.DaftarHarga()
		assert.Nil(t, err)
		if err == nil {
			t.Log("Product 1:", daftarHarga.Data[0])
		}
	})

	t.Run("Daftar Harga Gagal", func(t *testing.T) {
		Digi.Username = "ini_username_salah"
		_, err := Digi.DaftarHarga()
		assert.Nil(t, err, "Gagal Pesan: "+err.Error())
	})
}
func TestTopup(t *testing.T) {
	t.Run("Topup Berhasil", func(t *testing.T) {
		up, err := Digi.Topup("PLN20", "56604171910", "asfhvasfvha")
		assert.Nil(t, err)

		if err == nil {
			t.Logf("Refid: %s, CustomerNo: %s, BuyerSkuCode: %s, Message: %s, Status: %s, Rc: %s, Sn: %s, BuyerLastSaldo: %d, Price: %d, Tele: %s, Wa: %s", up.Data.RefId, up.Data.CustomerNo, up.Data.BuyerSkuCode, up.Data.Message, up.Data.Status, up.Data.Rc, up.Data.Sn, up.Data.BuyerLastSaldo, up.Data.Price, up.Data.Tele, up.Data.Wa)
		}
	})

	t.Run("Topup Gagal", func(t *testing.T) {
		_, err := Digi.Topup("Fucur", "56604171910", "asfhvasfvha")
		assert.Nil(t, err, "Error: "+err.Error())
	})
}
