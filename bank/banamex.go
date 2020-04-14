package bank

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
func (b *banamex) GetExchangeRate() (float32, error) {
	body, err := getHTML(b.url)
	if err != nil {
		return 0, err
	}
	reader := bytes.NewReader(body)

	content := getHTMLCells(reader)
	//Get the sell text, and drop the "$" sign
	dolaresCompra := strings.Replace(content[1], "$", "", 1)
	tipoDeCambio, err := strconv.ParseFloat(dolaresCompra, 32)
	if err != nil {
		return 0, err
	}
	return float32(tipoDeCambio), nil
}

func (b *banamex) URL() string {
	return b.url
}

//getHTMLCells obtiene las celdas de la tabla
func getHTMLCells(reader io.Reader) []string {

	z := html.NewTokenizer(reader)
	content := []string{}
	//Busca los elementos de una tabla en la pagina
	//Y los guarda en la variable content
	banamex := 9
	contador := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return content
		case html.StartTagToken:
			i := z.Token()
			switch i.Data {
			case "tr":
				contador++
			case "td":
				if contador == banamex {
					z.Next()
					switch z.Token().Data {
					case "span":
						if inner := z.Next(); inner == html.TextToken {
							text := (string)(z.Text())
							t := strings.TrimSpace(text)
							content = append(content, t)
						}

					}
				}
			}

		}
	}
}
