package demo1

import (
	"github.com/starter-go/libgorm"
	"gorm.io/gorm"
)

////////////////////////////////////////////////////////////////////////////////

// TableReg ...
type TableReg struct {

	//starter:component
	_as func(libgorm.GroupRegistry, Source) //starter:as(".","#")

	Prefix string //starter:inject("table-group.demo1.table-name-prefix")

	agent libgorm.DatabaseAgent
}

func (inst *TableReg) _impl() {
	inst._as(inst, inst)
}

func (inst *TableReg) init(c *libgorm.TableContext) {
	inst.agent.Init(c.Database)
}

// Group ...
func (inst *TableReg) Group() *libgorm.Group {

	list := make([]any, 0)
	list = append(list, &TableA{})
	list = append(list, &TableB{})
	list = append(list, &TableC{})

	// prefix
	prefix := inst.Prefix
	theTableNamePrefix = prefix

	return &libgorm.Group{
		Enabled: true,
		Name:    "default",
		//	Namespace:  "src/demo/demo1",
		Prefix:     prefix,
		Prototypes: list,
		OnInit:     inst.init,
	}
}

// DB ...
func (inst *TableReg) DB(db *gorm.DB) *gorm.DB {
	return inst.agent.DB(db)
}
