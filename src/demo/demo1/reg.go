package demo1

import (
	"github.com/starter-go/libgorm"
	"gorm.io/gorm"
)

// TableReg ...
type TableReg struct {

	//starter:component
	_as func(libgorm.GroupRegistry, MyAgent) //starter:as(".","#")

	DSMan      libgorm.DataSourceManager //starter:inject("#")
	Prefix     string                    //starter:inject("datagroup.demo1.table-name-prefix")
	URI        string                    //starter:inject("datagroup.demo1.uri")
	SourceName string                    //starter:inject("datagroup.demo1.datasource")

	agent libgorm.DataSourceAgent
}

func (inst *TableReg) _impl() {
	inst._as(inst, inst)
}

// Groups ...
func (inst *TableReg) Groups() []*libgorm.GroupRegistration {

	prefix := inst.Prefix
	theTableNamePrefix = prefix

	r1 := &libgorm.GroupRegistration{
		URI:    "uri:datagroup:demo1",
		Prefix: prefix,
		Group:  inst,
		Source: inst.SourceName,
	}

	return []*libgorm.GroupRegistration{r1}
}

// Prototypes ...
func (inst *TableReg) Prototypes() []any {

	// tables
	list := make([]any, 0)
	list = append(list, &TableA{})
	list = append(list, &TableB{})
	list = append(list, &TableC{})

	return list
}

// DB ...
func (inst *TableReg) DB(db *gorm.DB) *gorm.DB {
	agent := &inst.agent
	if !agent.Ready() {
		agent.Init(inst.DSMan, inst.SourceName)
	}
	return agent.DB(db)
}
