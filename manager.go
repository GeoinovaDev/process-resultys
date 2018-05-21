package process

import (
	"sync"
	"time"

	"git.resultys.com.br/lib/lower/exec"
	"git.resultys.com.br/lib/lower/log"
	"git.resultys.com.br/lib/lower/net/loopback"
	"git.resultys.com.br/motor/service"
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

// Prepare ...
func (manager *Manager) Prepare() []*Routine {
	m := []*Routine{}

	for _, routine := range manager.Routines {
		if routine.IsRun {
			m = append(m, routine)
		}
	}

	return m
}

// Start ...
func (manager *Manager) Start(unit *service.Unit) {
	go func() {
		manager.Diagnostic.Start()

		routines := manager.Prepare()
		total := len(routines)

		wg := &sync.WaitGroup{}
		wg.Add(total)

		for i := 0; i < total; i++ {
			go manager.run(routines[i], &Process{
				Unit: unit,
				wg:   wg,
			})
		}

		wg.Wait()
		unit.Release()

		manager.Diagnostic.Stop()
	}()
}

func (manager *Manager) run(routine *Routine, process *Process) {
	exec.Try(func() {
		process.Diagnostic.Start()
		routine.Func(routine, process)
	}).Catch(func(message string) {
		process.Done(false)
		log.Logger.Save(message, log.WARNING, loopback.IP())
	})
}

// Stats ...
func (manager *Manager) Stats() time.Duration {
	var elapsed time.Duration

	for i := 0; i < len(manager.Routines); i++ {
		// routine := manager.Routines[i]
		// elapsed += routine.Diagnostic.Elapsed
		// fmt.Printf("\n== routine = %s \n 1. start = %s \n 2. stop  = %s \n 3. elapsed = %s\n", routine.Name, routine.Diagnostic.StartTime, routine.Diagnostic.StopTime, routine.Diagnostic.Elapsed)
	}

	return elapsed
}
