package process

import (
	"sync"

	"git.resultys.com.br/motor/orchestrator"
	"git.resultys.com.br/motor/service"
)

// Routine struct
type Routine struct {
	Name    string
	Success bool
	IsBlock bool
	IsRun   bool
	Func    func(*Routine, *service.Unit)

	Param        interface{}
	Orchestrator *orchestrator.Orchestrator

	Diagnostic Diagnostic

	wg *sync.WaitGroup

	isDone bool
}

// Done ...
func (routine *Routine) Done(success bool) {
	if routine.isDone {
		return
	}
	
	routine.isDone = true
	routine.Success = success
	routine.wg.Done()
	routine.Diagnostic.Stop()
}
