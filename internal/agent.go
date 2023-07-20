package internal

import (
	"github.com/starter-go/libgorm"
	"gorm.io/gorm"
)

// DataSourceAgent 实现默认的数据源代理
type DataSourceAgent struct {
	//starter:component
	_as func(libgorm.Agent) //starter:as("#")

	DSM        libgorm.DataSourceManager //starter:inject("#")
	SourceName string                    //starter:inject("${libgorm.agent.source}")

	cached *gorm.DB
}

func (inst *DataSourceAgent) _impl() {
	inst._as(inst)
}

// DB ...
func (inst *DataSourceAgent) DB(db *gorm.DB) *gorm.DB {
	if db == nil {
		db = inst.get()
	}
	return db
}

func (inst *DataSourceAgent) get() *gorm.DB {
	db := inst.cached
	if db == nil {
		db2, err := inst.load()
		if err != nil {
			panic(err)
		}
		db = db2
		inst.cached = db2
	}
	return db
}

func (inst *DataSourceAgent) load() (*gorm.DB, error) {
	name := inst.SourceName
	ds, err := inst.DSM.GetDataSource(name)
	if err != nil {
		return nil, err
	}
	return ds.DB()
}
