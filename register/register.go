package register

type LWWRegister struct {
	id    string
	state *State
}

func (register *LWWRegister) Merge(remoteState *State) {
	state := *register.state
	if state.isWin(remoteState) {
		state.SetColor(remoteState.color)
	}
}
