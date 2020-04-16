package main

import (
	"fmt"
	"time"

	"github.com/dannywolfmx/NotiBank/bank"
	"github.com/dannywolfmx/NotiBank/notificacion"
)

var actualExchangeRate float32

func main() {
	run()
}

func run() {
	banco, err := bank.FactoryBank("banamex")

	if err != nil {
		panic(err)
	}

	not := notificacion.NewTray()
	for {

		if err := notificaTipoDeCambio(banco, not); err != nil {
			notificarErrorConexion(not)
		}
		//30 minutos
		time.Sleep(10 * time.Second)
	}
}

func notificaTipoDeCambio(banco bank.Bank, not notificacion.Tray) error {

	newExchangeRate, err := banco.GetExchangeRate()
	if err != nil {
		return err
	}

	//Determina si el tipo de cambio subio o bajo
	diferencia := newExchangeRate - actualExchangeRate

	//Obten titulo de notificacion en base a su diferencia
	status := dameStatus(diferencia)

	//update the actualExchangeRate
	actualExchangeRate = newExchangeRate
	//Formatear el tipo de cambio para que sea un valor monetario
	message := fmt.Sprintf("$ %.2f", newExchangeRate)

	return not.SetMessage(message, status).Show()
}

func notificarErrorConexion(not notificacion.Tray) error {
	return not.SetMessage("No se pudo conectar al banco", notificacion.ErrorConnection).Show()
}

//Determina el tipo de mensaje si es que subio o bajo el tipo de cambio
func dameStatus(exchange float32) int {
	if exchange > 0 {
		return notificacion.ExchangeUp
	} else if exchange < 0 {
		return notificacion.ExchangeDown
	}
	return notificacion.ExchangeSame
}
