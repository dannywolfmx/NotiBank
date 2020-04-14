package bank

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type Bank interface {
	GetExchangeRate() (float32, error)
	URL() string
}

func FactoryBank(name string) (Bank, error) {
	var bank Bank
	switch name {
	case "banorte":
		bank = &banorte{
			url: "https://www.banorte.com/wps/portal/banorte/Home/indicadores/dolares-y-divisas",
		}
	case "banamex":
		bank = &banamex{
			url: "https://dolarenmexico.com/",
		}
	default:
		return nil, errors.New("Bank don't found")
	}

	return bank, nil
}

//getHTML Get a html response
func getHTML(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return html, nil
}
