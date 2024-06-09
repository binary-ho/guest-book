package entities

type plan struct {
	level int8
	name  string
}

const (
	FREE  = "FREE"
	JJANG = "JJANG"
)

var plans map[string]plan

func initPlans() {
	plans = map[string]plan{
		FREE:  {name: FREE, level: 1},
		JJANG: {name: JJANG, level: 2},
	}
}

func Free() plan {
	if plans == nil {
		initPlans()
	}
	return plans[FREE]
}

func Jjang() plan {
	if plans == nil {
		initPlans()
	}
	return plans[JJANG]
}

func (plan plan) IsBetter(another plan) bool {
	return plan.level > another.level
}
