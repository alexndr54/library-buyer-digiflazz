# Unofficial Library Digiflazz
### Library ini digunakan untuk melakukan transaksi,cek saldo,hingga handle webhooks transaksi dari digiflazz,Library ini di khususkan hanya sebagai "buyer" bukan "seller"

# Instalasi
`go get github.com/alexndr54/library-buyer-digiflazz`   

# Pengunaan
`Digi := client.NewDigiflazz("your username", "production/development key")  `
1. Cek Saldo   
  `response,err := Digi.CekSaldo()` 

2. Mendapatkan Daftar Harga   
    `response, err := Digi.DaftarHarga()`

3. Melakukan Topup Baru
    `response, err := Digi.Topup("sku code", "target", "rfid")`

4. Handle Webhook   
`yourfibervariabel.Post("/webhook", webhook.GofiberWebhookHandler)`   
catatan: webhook.GofiberWebhookHandler adalah handler yang digunakan untuk menghandle webhook dari digiflazz, handler ini menggunakan library [Gofiber](https://github.com/gofiber/fiber)