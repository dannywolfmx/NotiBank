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
		toast: &toast.Notification{
			AppID: "Tipo de cambio",
		},
	}
}

func (n *tray) SetMessage(mensaje string, status int) {
	n.toast.Message = mensaje
	//Check if exist a messege
	if m, ok := dafaultMessage[status]; ok {
		n.toast.Title = m.title
		//TODO esta funcion no verifica que exista el icono
		//Check if the file exist
		if icon, err := filepath.Abs(m.iconPath); err == nil {
			n.toast.Icon = icon
		}
	}

}

func (n *tray) Show() error {
	return n.toast.Push()
}

func (n *tray) getIcon(tipoNotificacion int) (message string, icon string) {
	switch tipoNotificacion {
	case ExchangeUp:
		message, icon = "Subio el tipo de cambio", ""
	case ExchangeDown:
		message, icon = "Bajo el tipo de cambio", ""
	case ExchangeSame:
		message, icon = "Tipo de cambio igual", ""
	case ErrorConnection:
		message, icon = "Error conexion", "equal.png"
	}
	return
}
