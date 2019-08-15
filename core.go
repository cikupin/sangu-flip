package flip

import (
	"io"
	"net/url"
	"strings"
)

const (
	BankListURL = "/banks"
)

type CoreGateway struct {
	Client Client
}

func (gateway *CoreGateway) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = gateway.Client.BaseURL + path

	return gateway.Client.Call(method, path, header, body, v)
}

func (gateway *CoreGateway) GetBanks(bankCode string) (resp []Banks, err error) {
	data := url.Values{}
	if bankCode != "" {
		data.Set("code", bankCode)
	}

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	err = gateway.Call("GET", BankListURL, headers, strings.NewReader(data.Encode()), &resp)
	if err != nil {
		return
	}

	return
}