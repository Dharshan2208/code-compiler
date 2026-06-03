package queue

import "sync/atomic"

type Stats struct {
	Submitted uint64
	Completed uint64
	Failed    uint64
}

func (s *Stats) IncSubmitted() {
	atomic.AddUint64(&s.Submitted, 1)
}

func (s *Stats) IncCompleted() {
	atomic.AddUint64(&s.Completed, 1)
}

func (s *Stats) IncFailed() {
	atomic.AddUint64(&s.Failed, 1)
}

// adding this so reads are atomic-safe
func (s *Stats) Snapshot() (submitted uint64, completed uint64, failed uint64) {
	return atomic.LoadUint64(&s.Submitted),
		atomic.LoadUint64(&s.Completed),
		atomic.LoadUint64(&s.Failed)
}
