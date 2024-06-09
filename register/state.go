package register

import (
	entity "guestbook/entity/user"
	"guestbook/rgb"
)

type State struct {
	user      entity.User
	timestamp uint
	value     rgb.RGB
}
