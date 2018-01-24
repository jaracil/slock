package slock

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	return &Mutex{ch: make(chan struct{}, 1)}
}

func (m *Mutex) Lock() {
	m.ch <- struct{}{}
}

func (m *Mutex) TryLock() bool {
	select {
	case m.ch <- struct{}{}:
		return true
	default:
		return false
	}

}

func (m *Mutex) Unlock() {
	select {
	case <-m.ch:
		return
	default:
		panic("Unlock unlocked mutex")
	}
}
