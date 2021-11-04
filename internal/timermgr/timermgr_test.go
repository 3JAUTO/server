package timermgr

import (
	"testing"
	"time"

	"github.com/JEDIAC/server/internal/timermgr/timers"
	"github.com/JEDIAC/server/internal/util"
)

func TestTimer(t *testing.T) {
	mgr := NewTimerMgr()
	mgr.RegisterTimer(&timers.Timer1{})
	id := mgr.DispatchTimer(1, time.Now().UnixNano()+3*int64(time.Second), util.M{"message": 1234})
	if id == 0 {
		t.Fatal("fail to dispatch timer")
	}
	time.Sleep(time.Second * 3)
}

func TestTicker(t *testing.T) {
	mgr := NewTimerMgr()
	mgr.RegisterTimer((*timers.Timer1)(nil))
	id := mgr.DispatchTicker(1, int64(time.Second), util.M{"message": 1234})
	if id == 0 {
		t.Fatal("fail to dispatch ticker")
	}
	time.Sleep(time.Second * 3)
	mgr.StopTicker(id)
	time.Sleep(time.Second * 2)
}
