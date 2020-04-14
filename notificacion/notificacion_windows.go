// +build windows,!linux,!freebsd,!netbsd,!openbsd,!darwin,!js
package notificacion

import (
	"path/filepath"

	"gopkg.in/toast.v1"
)

type tray struct {
	toast *toast.Notification
}

func NewTray() *tray {
	return &tray{
		toast: crearToast(),
	}
}

func (n *tray) SetMessage(mensaje string) {
	n.toast.Message = mensaje
}

func (n *tray) Show(tipoAlerta float32) error {
	icon := ""
	var err error
	if tipoAlerta > 0 {
		n.toast.Title = "Subio el tipo de cambio"
		icon, err = filepath.Abs("up-red.png")
	} else if tipoAlerta == 0 {
		n.toast.Title = "Tipo de cambio igual"
		icon, err = filepath.Abs("equal.png")
	} else {
		n.toast.Title = "Bajo el tipo de cambio"
		icon, err = filepath.Abs("down-green.png")
	}
	if err == nil {
		n.toast.Icon = icon
	}
	return n.toast.Push()
}

func crearToast() *toast.Notification {
	return &toast.Notification{
		AppID: "Tipo de cambio",
	}
}
