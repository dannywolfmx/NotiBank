package banco

import (
	"bytes"
	"io"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type banamex struct {
	url string
}

//TipoDeCabio devuelve un string con el tipo de cambio o un error de existir
func (b *banamex) TipoDeCambio() (float32, error) {
	body, err := dameHTML(b.url)
	if err != nil {
		return 0, err
	}
	reader := bytes.NewReader(body)

	content := dameCeltasTablaHDML(reader)

	//El elemento numero 4 son los dolares
	dolaresCompra := content[4]
	tipoDeCambio, err := strconv.ParseFloat(dolaresCompra, 32)
	if err != nil {
		return 0, err
	}
	return float32(tipoDeCambio), nil
}

func (b *banamex) URL() string {
	return b.url
}

//dameCeltasTablaHDML obtiene las celdas de la tabla
func dameCeltasTablaHDML(reader io.Reader) []string {

	z := html.NewTokenizer(reader)
	content := []string{}
	//Busca los elementos de una tabla en la pagina
	//Y los guarda en la variable content
	for z.Token().Data != "html" {
		if tt := z.Next(); tt == html.StartTagToken {
			if t := z.Token(); t.Data == "td" {
				if inner := z.Next(); inner == html.TextToken {
					text := (string)(z.Text())
					t := strings.TrimSpace(text)
					content = append(content, t)
				}
			}
		}
	}
	return content
}
