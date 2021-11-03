package timermgr

import (
	"testing"
	"time"

	"github.com/JEDIAC/server/internal/util"
)

func TestTimerMgr(t *testing.T) {
	mgr := NewTimerMgr()
	mgr.RegisterTimer(&Timer1{})
	id := mgr.DispatchTimer(1, time.Now().UnixNano()+3*int64(time.Second), util.M{"message": 1234})
	if id == 0 {
		panic("fail dispatch timer")
	}
	time.Sleep(time.Second * 3)
}
