package process

import (
	"git.resultys.com.br/motor/service"
)

// Routine struct
type Routine struct {
	Name    string
	IsAsync bool
	Func    func(*service.Unit, ...interface{})
}
