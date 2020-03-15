package banco

import (
	"bytes"
	"strconv"
)

type banorte struct {
	url string
}

//TipoDeCabio devuelve un string con el tipo de cambio o un error de existir
func (b *banorte) TipoDeCambio() (float32, error) {
	//banorte fija su contenido con javascript, no podemos parsear el html sin un engine
	//Por lo anterior prefiero buscarlo a mano
	body, err := dameHTML(b.url)
	if err != nil {
		return 0, err
	}
	//Tipo de cambio posicion 53 toma 5 caracteres
	cadena := []byte(`"nombreDolar":"VENTANILLA"`)
	//fmt.Print(string(body[271000:28000]))
	index := bytes.LastIndex(body, cadena)
	index += 53
	index2 := index + 5
	tipoDeCambio, err := strconv.ParseFloat(string(body[index:index2]), 32)
	if err != nil {
		return 0, err
	}
	return float32(tipoDeCambio), nil
}

func (b *banorte) URL() string {
	return b.url
}
