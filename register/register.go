package register

type LWWRegister struct {
	state *State
}

func CreateLLWWRegister(state *State) *LWWRegister {
	return &LWWRegister{state}
}

func (register *LWWRegister) Merge(remoteState *State) *LWWRegister {
	if remoteState == nil {
		return register
	}

	if state := *register.state; state.isWin(remoteState) {
		state.SetColor(remoteState.color)
	}
	return register
}
