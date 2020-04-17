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

func (n *tray) SetMessage(mensaje string, status int) Tray {
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
	return n
}

func (n *tray) Show() error {
	return n.toast.Push()
}
