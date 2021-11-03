package timermgr

func (tm *TimerMgr) SaveToDB() {
	tm.mtx.Lock()
	defer tm.mtx.Unlock()

}

func (tm *TimerMgr) LoadFromDB() {
	tm.mtx.Lock()
	defer tm.mtx.Unlock()
}
