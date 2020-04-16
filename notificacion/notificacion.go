package notificacion

//Tipos de notificaciones status
const (
	ExchangeUp = iota
	ExchangeDown
	ExchangeSame
	ErrorConnection
)

type message struct {
	title, iconPath, text string
}

var dafaultMessage = map[int]message{
	ExchangeUp: {
		title:    "Subio el tipo de cambio",
		iconPath: "up-red.png",
	},
	ExchangeDown: {
		title:    "Bajo el tipo de cambio",
		iconPath: "down-green.png",
	},
	ExchangeSame: {
		title:    "Tipo de cambio igual",
		iconPath: "equal.png",
	},
	//TODO: Hacerle un icono
	ErrorConnection: {
		title: "Error conexion",
	},
}

//Tray muestra una notificacion en base a la configuracion dada
type Tray interface {
	Show() error
	SetMessage(message string, status int) Tray
}
