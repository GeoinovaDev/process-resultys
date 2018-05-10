package process

import "time"

// Diagnostic struct
type Diagnostic struct {
	Elapsed   time.Duration
	StartTime time.Time
	StopTime  time.Time
}

// Start inicia
func (d *Diagnostic) Start() {
	d.StartTime = time.Now()
}

// Stop ...
func (d *Diagnostic) Stop() {
	d.Elapsed = time.Since(d.StartTime)
	d.StopTime = time.Now()
}
