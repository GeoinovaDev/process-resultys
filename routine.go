package process

import (
	"git.resultys.com.br/motor/orchestrator"
)

// Routine struct
type Routine struct {
	Name         string
	IsBlock      bool
	IsRun        bool
	Func         func(*Routine, *Process)
	Orchestrator *orchestrator.Orchestrator
	Param        interface{}
}
