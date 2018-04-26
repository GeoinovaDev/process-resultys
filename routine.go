package process

import (
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
}
