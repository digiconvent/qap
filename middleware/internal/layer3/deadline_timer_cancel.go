package layer3

func (t *DeadlineTimer) Cancel(id string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if timer, ok := t.timers[id]; ok {
		timer.Stop()
		delete(t.timers, id)
	}
}
