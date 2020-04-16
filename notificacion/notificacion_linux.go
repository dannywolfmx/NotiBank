// +build linux freebsd netbsd openbsd !windows
package notificacion

import (
	"github.com/gen2brain/beeep"
)

func NewTray() *tray {
	return &tray{}
}

type tray struct {
	message, title string
}

func (t *tray) Show() error {
	if err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration); err != nil {
		return err
	}

	if err := beeep.Notify(t.title, t.message, ""); err != nil {
		return err
	}
	return nil
}

func (t *tray) SetMessage(message string, status int) Tray {
	t.message = message

	if m, ok := dafaultMessage[status]; ok {
		t.title = m.title
	} else {
		t.title = "Tipo de cambio"
	}
	return t
}
