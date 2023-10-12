package libgorm

import (
	"fmt"

	"gorm.io/gorm"
)

// DataSource 数据源
type DataSource interface {
	DB() (*gorm.DB, error)
	Registration() *DataSourceRegistration
}

// DataSourceRegistration 数据源注册信息
type DataSourceRegistration struct {
	Enabled    bool
	Name       string
	DataSource DataSource
	Groups     []string
}

// DataSourceManager 数据源管理器
type DataSourceManager interface {
	GetDataSource(name string) (DataSource, error)
	ListNames() []string
}

// Agent 是默认的数据源代理
type Agent interface {
	DB(db *gorm.DB) *gorm.DB
}

////////////////////////////////////////////////////////////////////////////////

// DatabaseAgent ...
type DatabaseAgent struct {
	db1 Database
	db2 *gorm.DB
}

func (inst *DatabaseAgent) _impl(a Agent) {
	a = inst
}

// Init 初始化
func (inst *DatabaseAgent) Init(db Database) {
	if inst.db1 == nil {
		inst.db1 = db
	}
}

func (inst *DatabaseAgent) load() (*gorm.DB, error) {
	d1 := inst.db1
	if d1 == nil {
		return nil, fmt.Errorf("DatabaseAgent without source")
	}
	d2 := inst.db1.DB()
	if d2 == nil {
		return nil, fmt.Errorf("DatabaseAgent without source")
	}
	return d2, nil
}

func (inst *DatabaseAgent) get() *gorm.DB {
	db := inst.db2
	if db == nil {
		want, err := inst.load()
		if err != nil {
			panic(err)
		}
		db = want
		inst.db2 = db
	}
	return db
}

// DB 如有需要，获取 DB
func (inst *DatabaseAgent) DB(db *gorm.DB) *gorm.DB {
	if db == nil {
		db = inst.get()
	}
	return db
}

////////////////////////////////////////////////////////////////////////////////
