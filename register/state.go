package register

import "guestbook/rgb"

type State struct {
	peer      string
	timestamp uint
	value     rgb.RGB
}
