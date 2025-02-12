package model

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
