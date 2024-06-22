package user

import (
	"database/sql/driver"
	"errors"
	"guestbook/common/util"
)

type plan struct {
	level int8
	name  string
}

const (
	FREE  = "FREE"
	JJANG = "JJANG"
)

var plans map[string]plan

func init() {
	if plans != nil {
		return
	}

	plans = map[string]plan{
		FREE:  {name: FREE, level: 1},
		JJANG: {name: JJANG, level: 2},
	}
}

func FreePlan() plan {
	return plans[FREE]
}

func JjangPlan() plan {
	return plans[JJANG]
}

func (plan *plan) IsHigher(another *plan) bool {
	return plan.level > another.level
}

func (plan *plan) IsEquals(another *plan) bool {
	return plan.level == another.level
}

func (plan plan) Value() (driver.Value, error) {
	if util.IsBlank(plan.name) {
		return nil, errors.New("plan.name is empty")
	}
	return plan.name, nil
}

func (plan *plan) Scan(value interface{}) error {
	result, ok := value.([]byte)
	if !ok {
		return errors.New("plan을 []byte type으로 형변환 불가능")
	}

	planName := string(result[:])
	planByName, err := getPlanByName(planName)
	if err != nil {
		return err
	}
	*plan = *planByName
	return nil
}

func getPlanByName(name string) (*plan, error) {
	for _, plan := range plans {
		if plan.name != name {
			continue
		}
		return &plan, nil
	}
	return nil, errors.New("'" + name + "'라는 이름을 가진 plan 없음")
}
