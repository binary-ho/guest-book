package user

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
	if plans != nil {
		return
	}

	plans = map[string]plan{
		FREE:  {name: FREE, level: 1},
		JJANG: {name: JJANG, level: 2},
	}
}

func FreePlan() plan {
	initPlans()
	return plans[FREE]
}

func JjangPlan() plan {
	initPlans()
	return plans[JJANG]
}

func (plan *plan) IsHigher(another *plan) bool {
	return plan.level > another.level
}

func (plan *plan) IsEquals(another *plan) bool {
	return plan.level == another.level
}
