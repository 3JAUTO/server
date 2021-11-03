package timermgr

import (
	"time"

	"github.com/JEDIAC/server/internal/util"
)

type ITimer interface {
	Type() TimerType
	Call(ts int64, args util.M)
	SaveToDB() bool
}

type DBTimer struct {
	ID        int64
	TimerType TimerType
	Timer     int64
	Tick      bool
	Args      []byte
}

type timerMeta struct {
	_type TimerType
	_time int64
	_args util.M
	_raw  *time.Timer
}

type tickerMeta struct {
	_type TimerType
	_time int64
	_args util.M
	_raw  *time.Ticker
}
