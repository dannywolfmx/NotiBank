// +build linux freebsd netbsd openbsd !windows
package notificacion

import (
	"github.com/gen2brain/beeep"
)

func NewTray() *tray {
	return &tray{}
}

type tray struct {
	mensaje string
}

func (n *tray) Show(tipoAlerta float32) error {
	if err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration); err != nil {
		return err
	}

	if err := beeep.Notify("Tipo de cambio", n.mensaje, ""); err != nil {
		return err
	}
	return nil
}

func (n *tray) SetMessage(mensaje string) {
	n.mensaje = mensaje
}
