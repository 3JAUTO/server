package timermgr

import (
	"fmt"
	"time"

	"github.com/JEDIAC/server/internal/util"
)

type TimerType = uint16

func NewTimerMgr() *TimerMgr {
	return &TimerMgr{
		timers:    make(map[uint16]ITimer),
		onTimers:  make(map[int64]*timerMeta),
		onTickers: make(map[int64]*tickerMeta),
	}
}

func (tm *TimerMgr) RegisterTimer(t ITimer) error {
	if tm.timers[t.Type()] != nil {
		return fmt.Errorf("timer %d already registered", t.Type())
	}
	tm.timers[t.Type()] = t
	return nil
}

func (tm *TimerMgr) DispatchTimer(typ TimerType, execTs int64, args util.M) int64 {
	id := time.Now().UnixNano()
	if execTs <= id {
		return 0
	}
	meta := &timerMeta{
		_type: typ,
		_time: execTs,
		_args: args,
	}

	_timer := time.NewTimer(time.Duration(execTs - id))
	if _timer == nil {
		return 0
	}
	meta._raw = _timer
	tm.onTimers[id] = meta

	go func(timerID int64) {
		inMeta := tm.onTimers[timerID]
		itimer := tm.timers[inMeta._type]
		now := <-inMeta._raw.C
		itimer.Call(now.UnixNano(), inMeta._args)
		delete(tm.onTimers, timerID)
	}(id)
	return id
}

func (tm *TimerMgr) DispatchTicker(typ TimerType, gapTs int64, args util.M) int64 {
	id := time.Now().UnixNano()

	meta := &tickerMeta{
		_type: typ,
		_time: gapTs,
		_args: args,
	}

	_ticker := time.NewTicker(time.Duration(gapTs))
	if _ticker == nil {
		return 0
	}
	meta._raw = _ticker
	tm.onTickers[id] = meta

	go func(tickerID int64) {
		inMeta := tm.onTickers[tickerID]
		itimer := tm.timers[inMeta._type]
		now := <-inMeta._raw.C
		itimer.Call(now.UnixNano(), inMeta._args)
	}(id)
	return id
}
