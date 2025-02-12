package webook

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type Payload struct {
	Event string `json:"event"`
	Data  struct {
		TrxID        string `json:"trx_id"`
		RefID        string `json:"ref_id"`
		CustomerNo   string `json:"customer_no"`
		BuyerSKUCode string `json:"buyer_sku_code"`
		Message      string `json:"message"`
		Status       string `json:"status"`
		RC           string `json:"rc"`
		BuyerSaldo   int    `json:"buyer_last_saldo"`
		SN           string `json:"sn"`
		Price        int    `json:"price"`
		Tele         string `json:"tele"`
		WA           string `json:"wa"`
	} `json:"data"`
}

// GofiberWebhookHandler untuk meng handle webhook yang dikirim dari digiflazz, menggunakan framework gofiber
// *fiber.Ctx adalah parameter yang diperlukan untuk mengambil data dari request
func GofiberWebhookHandler(c *fiber.Ctx) (error, *Payload) {
	event := c.GetReqHeaders()["X-Digiflazz-Event"][0]
	var payload *Payload

	if err := c.BodyParser(&payload); err != nil {
		return errors.New("Error parsing JSON"), nil
	}

	payload.Event = event
	return nil, payload
}
