package register

import (
	entity "guestbook/entity/user"
	"guestbook/rgb"
)

type State struct {
	color          rgb.RGB
	timestamp      uint32
	lastAccessUser entity.User
}

func (state *State) SetColor(rgb rgb.RGB) {
	state.color = rgb
}

func (state *State) isWin(remoteState *State) bool {
	if isTimestampBigger(state, remoteState) ||
		isPlanHigher(state, remoteState) ||
		isIDLower(state, remoteState) {
		return true
	}
	return false
}

func isTimestampBigger(state *State, remoteState *State) bool {
	return state.timestamp > remoteState.timestamp
}

func isPlanHigher(state *State, remoteState *State) bool {
	plan := state.lastAccessUser.Plan()
	remotePlan := remoteState.lastAccessUser.Plan()
	return state.timestamp == remoteState.timestamp &&
		plan.IsHigher(&remotePlan)
}

func isIDLower(state *State, remoteState *State) bool {
	plan := state.lastAccessUser.Plan()
	remotePlan := remoteState.lastAccessUser.Plan()
	return state.timestamp == remoteState.timestamp &&
		plan.IsEquals(&remotePlan) &&
		state.lastAccessUser.ID() < remoteState.lastAccessUser.ID()
}
