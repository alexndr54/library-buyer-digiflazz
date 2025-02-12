package model

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
