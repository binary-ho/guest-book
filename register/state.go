package register

import (
	entities "guestbook/entity/user"
	"guestbook/rgb"
)

type State struct {
	user      entities.User
	timestamp uint
	value     rgb.RGB
}
