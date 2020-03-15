package banco

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type Banco interface {
	TipoDeCambio() (float32, error)
	URL() string
}

func Factorybanco(nombre string) (Banco, error) {
	var banco Banco
	switch nombre {
	case "banorte":
		banco = &banorte{
			url: "https://www.banorte.com/wps/portal/banorte/Home/indicadores/dolares-y-divisas",
		}
	case "banamex":
		banco = &banamex{
			url: "https://portal.banamex.com.mx/c719_004/economiaFinanzas/es/home",
		}
	default:
		return nil, errors.New("Banco no localizado")
	}

	return banco, nil
}

//dameHTML recibe una URL y returna []byte con el contenido de la pagina o un error
func dameHTML(url string) ([]byte, error) {
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
