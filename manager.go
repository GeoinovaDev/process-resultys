package process

import (
	"sync"
	"time"

	"github.com/GeoinovaDev/lower-resultys/exception"
	"github.com/GeoinovaDev/lower-resultys/exec"
	"github.com/GeoinovaDev/service-resultys"
)

// Manager struct
type Manager struct {
	Routines   []*Routine
	Diagnostic Diagnostic
}

// New ...
func New() *Manager {
	return &Manager{}
}

// AddRoutine ...
func (manager *Manager) AddRoutine(routine *Routine) {
	manager.Routines = append(manager.Routines, routine)
}

// Start ...
func (manager *Manager) Start(unit *service.Unit, params ...interface{}) {
	manager.Diagnostic.Start()

	routines := manager.prepare()
	total := len(routines)

	wg := &sync.WaitGroup{}
	wg.Add(total)

	for i := 0; i < total; i++ {
		process := &Process{
			wg: wg,
		}

		params = append(params, process)

		go manager.run(routines[i], unit, process, params...)
	}

	wg.Wait()
	unit.Release()

	manager.Diagnostic.Stop()
}

// Stats ...
func (manager *Manager) Stats() time.Duration {
	var elapsed time.Duration

	for i := 0; i < len(manager.Routines); i++ {
		elapsed += manager.Diagnostic.Elapsed
	}

	return elapsed
}

func (manager *Manager) prepare() []*Routine {
	m := []*Routine{}

	for _, routine := range manager.Routines {
		m = append(m, routine)
	}

	return m
}

func (manager *Manager) run(routine *Routine, unit *service.Unit, process *Process, params ...interface{}) {
	exec.Try(func() {
		routine.Func(unit, params...)

		if !routine.IsAsync {
			process.Finish()
		}
	}).Catch(func(message string) {
		exception.Raise(message, exception.WARNING)
	})
}
