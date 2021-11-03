package timermgr

import "sync"

type TimerMgr struct {
	timers map[TimerType]ITimer // registered timer

	onTimers  map[int64]*timerMeta // starting timer
	onTickers map[int64]*tickerMeta

	mtx sync.Mutex
}
