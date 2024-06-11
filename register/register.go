package register

type LWWRegister struct {
	state *State
}

func (register *LWWRegister) Merge(remoteState *State) {
	if remoteState == nil {
		return
	}

	if state := *register.state; state.isWin(remoteState) {
		state.SetColor(remoteState.color)
	}
}
