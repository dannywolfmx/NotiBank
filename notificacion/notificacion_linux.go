// +build linux freebsd netbsd openbsd !windows
package notificacion

import (
	"github.com/gen2brain/beeep"
)

func DameNuevaNotificacion() *notificacion {
	return &notificacion{}
}

type notificacion struct {
	mensaje string
}

func (n *notificacion) Mostrar(tipoAlerta float32) error {
	if err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration); err != nil {
		return err
	}

	if err := beeep.Notify("Tipo de cambio", n.mensaje, ""); err != nil {
		return err
	}
	return nil
}

func (n *notificacion) FijaMensaje(mensaje string) {
	n.mensaje = mensaje
}
