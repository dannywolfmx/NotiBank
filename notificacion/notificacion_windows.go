// +build windows,!linux,!freebsd,!netbsd,!openbsd,!darwin,!js
package notificacion

import (
	"gopkg.in/toast.v1"
)

type notificacion struct {
	toast *toast.Notification
}

func dameNuevaNotificacion() *notificacion {
	return &notificacion{
		toast: crearToast(),
	}
}

func (n *notificacion) FijaMensaje(mensaje string) {
	n.toast.Message = mensaje
}

func (n *notificacion) Mostrar(tipoAlerta float32) error {
	if tipoAlerta > 0 {
		n.toast.Title = "Subio el tipo de cambio"
	} else if tipoAlerta == 0 {
		n.toast.Title = "Tipo de cambio igual"
	} else {
		n.toast.Title = "Bajo el tipo de cambio"
	}
	return n.toast.Push()
}

func crearToast() *toast.Notification {
	notification := &toast.Notification{
		AppID: "Tipo de cambio",
	}
	return notification
}
