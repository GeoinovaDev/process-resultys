package process

import (
	"github.com/GeoinovaDev/service-resultys"
)

// Routine struct
type Routine struct {
	Name    string
	IsAsync bool
	Func    func(*service.Unit, ...interface{})
}
