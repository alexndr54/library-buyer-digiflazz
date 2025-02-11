package test

import (
	"digiflazz/client"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	USERNAME        = ""
	DEVELOPMENT_KEY = ""
)

func TestCekSaldo(t *testing.T) {
	t.Run("Cek Saldo Berhasil", func(t *testing.T) {
		saldo, err := client.CekSaldo(USERNAME, DEVELOPMENT_KEY)
		assert.Nil(t, err)
		t.Log("Saldo: ", saldo.Data.Deposit)
	})

	t.Run("Cek Saldo Gagal", func(t *testing.T) {
		_, err := client.CekSaldo("golekworing", DEVELOPMENT_KEY)
		assert.Nil(t, err, "Gagal Pesan: "+err.Error())
	})
}
func TestDaftarHarga(t *testing.T) {
	t.Run("Daftar Harga Berhasil", func(t *testing.T) {

		daftarHarga, err := client.DaftarHarga(USERNAME, DEVELOPMENT_KEY)
		assert.Nil(t, err)
		if daftarHarga.Error.Data.Rc == "" {
			t.Log("Product 1:", daftarHarga.Data[0])
		}
	})

	t.Run("Daftar Harga Gagal", func(t *testing.T) {
		_, err := client.DaftarHarga("golekworing", DEVELOPMENT_KEY)
		assert.Nil(t, err, "Gagal Pesan: "+err.Error())
	})
}
func TestTopup(t *testing.T) {
	t.Run("Topup Berhasil", func(t *testing.T) {
		up, err := client.TopUp(USERNAME, DEVELOPMENT_KEY, "PLN20", "56604171910", "asfhvasfvha")
		assert.Nil(t, err)

		if err == nil {
			t.Log(fmt.Sprintf("Refid: %s, CustomerNo: %s, BuyerSkuCode: %s, Message: %s, Status: %s, Rc: %s, Sn: %s, BuyerLastSaldo: %d, Price: %d, Tele: %s, Wa: %s", up.Data.RefId, up.Data.CustomerNo, up.Data.BuyerSkuCode, up.Data.Message, up.Data.Status, up.Data.Rc, up.Data.Sn, up.Data.BuyerLastSaldo, up.Data.Price, up.Data.Tele, up.Data.Wa))
		}
	})

	t.Run("Gagal Topup", func(t *testing.T) {
		_, err := client.TopUp("", "", "Fucur", "56604171910", "asfhvasfvha")
		assert.Nil(t, err, "Error: "+err.Error())
	})
}
