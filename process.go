package process

import (
	"fmt"
	"sync"
	"time"

	"git.resultys.com.br/lib/lower/exec"
	"git.resultys.com.br/lib/lower/log"
	"git.resultys.com.br/lib/lower/net/loopback"
	"git.resultys.com.br/motor/service"
)

// Process struct
type Process struct {
	wg       *sync.WaitGroup
	Routines []*Routine

	Diagnostic Diagnostic
}

// Stats mostra estatistica de tempo
func (process *Process) Stats() time.Duration {
	var elapsed time.Duration

	for i := 0; i < len(process.Routines); i++ {
		routine := process.Routines[i]
		elapsed += routine.Diagnostic.Elapsed
		fmt.Printf("\n== routine = %s \n 1. start = %s \n 2. stop  = %s \n 3. elapsed = %s\n", routine.Name, routine.Diagnostic.StartTime, routine.Diagnostic.StopTime, routine.Diagnostic.Elapsed)
	}

	return elapsed
}

// New create process
func New() *Process {
	return &Process{wg: &sync.WaitGroup{}}
}

// AddRoutine adiciona core processamento
func (process *Process) AddRoutine(routine *Routine) {
	process.Routines = append(process.Routines, routine)
}

// Start executa as routines
func (process *Process) Start(unit *service.Unit) {
	go func() {
		process.Diagnostic.Start()

		rts := process.Prepare()
		process.wg.Add(len(rts))

		for i := 0; i < len(rts); i++ {
			go process.run(rts[i], unit)
		}

		process.wg.Wait()
		unit.Release()

		process.Diagnostic.Stop()
	}()
}

func (process *Process) run(routine *Routine, unit *service.Unit) {
	routine.wg = process.wg

	exec.Try(func() {
		routine.Reset()
		routine.Diagnostic.Start()
		routine.Func(routine, unit)
	}).Catch(func(message string) {
		routine.Done(false)
		log.Logger.Save(message, log.WARNING, loopback.IP())
	})
}

// Prepare inicia as routines
func (process *Process) Prepare() []*Routine {
	m := []*Routine{}

	for _, routine := range process.Routines {
		if routine.IsRun {
			m = append(m, routine)
		}
	}

	return m
}
