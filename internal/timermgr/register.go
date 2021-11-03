package timermgr

import (
	"fmt"

	"github.com/JEDIAC/server/internal/util"
)

type Timer1 struct{}

func (t *Timer1) Type() TimerType {
	return 1
}

func (t *Timer1) Call(ts int64, args util.M) {
	fmt.Println("Good Morning")
	fmt.Printf("%d\n", args["message"].(int))
}

func (t *Timer1) SaveToDB() bool {
	return false
}
