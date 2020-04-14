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
			fmt.Print(err)
		}
		//30 minutos
		time.Sleep(30 * time.Minute)
	}
}

func notificaTipoDeCambio(banco bank.Bank, not notificacion.Tray) error {

	newExchangeRate, err := banco.GetExchangeRate()
	if err != nil {
		return err
	}

	//Determina si el tipo de cambio subio o bajo
	alertType := newExchangeRate - actualExchangeRate

	actualExchangeRate = newExchangeRate
	//Formatear el tipo de cambio para que sea un valor monetario
	not.SetMessage("$ " + fmt.Sprintf("%.2f", newExchangeRate))

	return not.Show(alertType)
}
