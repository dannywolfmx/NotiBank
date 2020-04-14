package notificacion

const (
	Inicial = iota
	Alta
	Baja
)

//Tray muestra una notificacion en base a la configuracion dada
type Tray interface {
	Show(alertType float32) error
	SetMessage(string)
}
