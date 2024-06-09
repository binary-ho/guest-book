package register

import "guestbook/rgb"

type LWWRegister struct {
	id    string
	state *State
}

func (register *LWWRegister) Merge(remoteState State) {
	localState := *register.state
	if localState.timestamp > remoteState.timestamp {
		return
	}
	if localState.timestamp == remoteState.timestamp && localState.peer > remoteState.peer {
		return
	}

	localState.value = remoteState.value
}

func (register *LWWRegister) Value() rgb.RGB {
	return register.state.value
}

func (register *LWWRegister) SetValue(rgb rgb.RGB) {
	register.state.value = rgb
}
