package qtimer

const (
	StatusReady = 0 + iota
	StatusRunning
	StatusStopped
	StatusReset
	StatusClosed
)

type JobFn func()
