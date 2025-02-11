package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/alexndr54/library-buyer-digiflazz/helper"
	"io"
	"net/http"
)

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

func CekSaldo(username, apiKey string) (*CekSaldoResponse, error) {
	sign := helper.GenerateMD5Hash(username + apiKey + "depo")
	requestBody, err := json.Marshal(CekSaldoRequest{
		Cmd:      "deposit",
		Username: username,
		Sign:     sign,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://api.digiflazz.com/v1/cek-saldo", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response CekSaldoResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Data.Rc != "" {
		return nil, errors.New(response.Data.Message)
	}

	return &response, nil
}
