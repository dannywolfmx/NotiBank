package notificacion

const (
	Inicial = iota
	Alta
	Baja
)

//Notificacion muestra una notificacion en base a la configuracion dada
type Notificacion interface {
	Mostrar(tipoAlerta float32) error
	FijaMensaje(mensaje string)
}
