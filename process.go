package process

import (
	"sync"

	"git.resultys.com.br/motor/service"
)

// Process struct
type Process struct {
	Success    bool
	Diagnostic Diagnostic
	wg         *sync.WaitGroup
	isDone     bool
	Unit       *service.Unit
}

// Finish ...
func (process *Process) Finish() {
	if process.isDone {
		return
	}

	process.isDone = true
	process.wg.Done()
	process.Diagnostic.Stop()
}

// Done ...
func (process *Process) Done(isSuccess bool) {
	if process.isDone {
		return
	}

	process.isDone = true
	process.Success = isSuccess
	process.wg.Done()
	process.Diagnostic.Stop()
}
