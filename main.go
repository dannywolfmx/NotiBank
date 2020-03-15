package main

import (
	"fmt"
	"time"

	"github.com/dannywolfmx/tipocambio2/banco"
	"github.com/dannywolfmx/tipocambio2/notificacion"
)

func main() {
	run()
}

func run() {
	banco, err := banco.Factorybanco("banamex")

	if err != nil {
		panic(err)
	}

	not := notificacion.DameNuevaNotificacion()
	for {

		go notificaTipoDeCambio(banco, not)
		//30 minutos
		time.Sleep(1800000 * time.Second)
	}
}

func notificaTipoDeCambio(banco banco.Banco, not notificacion.Notificacion) {
	var tipoDeCambioViejo float32
	tipoDeCambio, err := banco.TipoDeCambio()
	if err != nil {
		fmt.Println(err)
	}
	//Determina si el tipo de cambio subio o bajo
	tipoAlerta := tipoDeCambio - tipoDeCambioViejo

	//Formatear el tipo de cambio para que sea un valor monetario
	not.FijaMensaje("$ " + fmt.Sprintf("%.2f", tipoDeCambio))

	if err := not.Mostrar(tipoAlerta); err != nil {
		fmt.Println(err)
	}
}
