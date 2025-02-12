package model

type CekSaldoRequest struct {
	Cmd      string `json:"cmd"`
	Username string `json:"username"`
	Sign     string `json:"sign"`
}

type CekSaldoResponse struct {
	Data struct {
		Deposit int    `json:"deposit"`
		Rc      string `json:"rc"`
		Message string `json:"message"`
	} `json:"data"`
}
