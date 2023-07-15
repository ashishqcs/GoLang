package client

import (
	"errors"
	"io"
	"net/http"
	"strconv"
)

type Client interface {
	GetBillById(id int) ([]byte, error)
}

type InternetProviderClient struct {
	Client http.Client
	Url    string
}

var ErrServerNotResponding = errors.New("server not responding")
var ErrFoundNoResponse = errors.New("no Response")

func NewInternetProviderClient() InternetProviderClient {
	return InternetProviderClient{
		Url: "http://localhost:8080/internet-bills/v1",
	}
}

func (ip InternetProviderClient) GetBillById(id int) ([]byte, error) {
	url := ip.Url + "/" + strconv.Itoa(id)
	ip.Client.Get(url)
	res, err := http.Get(url)

	if err != nil {
		// log.Printf("Error: %s", err.Error())
		return nil, ErrServerNotResponding
	}

	body, _ := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		// log.Printf("Error: Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil, ErrFoundNoResponse
	}

	return body, nil
}
