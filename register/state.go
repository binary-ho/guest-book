package register

import (
	entity "guestbook/entity/user"
	"guestbook/rgb"
)

type State struct {
	user      entity.User
	timestamp uint32
	color     rgb.RGB
}

func (state *State) SetColor(rgb rgb.RGB) {
	state.color = rgb
}

func (state *State) isWin(remoteState *State) bool {
	if isTimestampBigger(state, remoteState) ||
		isPlanHigher(state, remoteState) ||
		isIdLower(state, remoteState) {
		return true
	}
	return false
}

func isTimestampBigger(state *State, remoteState *State) bool {
	return state.timestamp > remoteState.timestamp
}

func isPlanHigher(state *State, remoteState *State) bool {
	plan := state.user.Plan()
	remotePlan := remoteState.user.Plan()
	return state.timestamp == remoteState.timestamp &&
		plan.IsHigher(&remotePlan)
}

func isIdLower(state *State, remoteState *State) bool {
	plan := state.user.Plan()
	remotePlan := remoteState.user.Plan()
	return state.timestamp == remoteState.timestamp &&
		plan.IsEquals(&remotePlan) &&
		state.user.Id() < remoteState.user.Id()
}
