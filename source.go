package libgorm

import (
	"fmt"

	"gorm.io/gorm"
)

// DataSource 数据源
type DataSource interface {
	DB() (*gorm.DB, error)
}

// DataSourceRegistration 数据源注册信息
type DataSourceRegistration struct {
	Alias         string
	Configuration Configuration
	DataSource    DataSource
	Enabled       bool

	// Groups []string // alias of groups; [已废弃]使用 ： group.datasource 代替
}

// DataSourceRegistry 数据源注册器
type DataSourceRegistry interface {
	ListSources() []*DataSourceRegistration
}

// DataSourceManager 数据源管理器
type DataSourceManager interface {
	GetDataSource(alias string) (DataSource, error)
	ListAliases() []string
}

// Agent 是默认的数据源代理
type Agent interface {
	DB(db *gorm.DB) *gorm.DB
}

////////////////////////////////////////////////////////////////////////////////

// DataSourceAgent ...
type DataSourceAgent struct {
	cached            *gorm.DB
	dataSourceManager DataSourceManager
	dataSourceName    string
}

func (inst *DataSourceAgent) _impl(a Agent) {
	a = inst
}

// Init 初始化
func (inst *DataSourceAgent) Init(dataSourceManager DataSourceManager, dataSourceName string) {
	inst.dataSourceManager = dataSourceManager
	inst.dataSourceName = dataSourceName
}

// Ready 确认是否已就绪
func (inst *DataSourceAgent) Ready() bool {
	ok1 := inst.dataSourceManager != nil
	ok2 := inst.dataSourceName != ""
	return ok1 && ok2
}

func (inst *DataSourceAgent) load() (*gorm.DB, error) {
	if !inst.Ready() {
		return nil, fmt.Errorf("DataSourceAgent is not ready")
	}
	name := inst.dataSourceName
	ds, err := inst.dataSourceManager.GetDataSource(name)
	if err != nil {
		return nil, err
	}
	return ds.DB()
}

func (inst *DataSourceAgent) get() *gorm.DB {
	db := inst.cached
	if db == nil {
		result, err := inst.load()
		if err != nil {
			panic(err)
		}
		db = result
		inst.cached = result
	}
	return db
}

// DB 如有需要，获取 DB
func (inst *DataSourceAgent) DB(db *gorm.DB) *gorm.DB {
	if db == nil {
		db = inst.get()
	}
	return db
}

////////////////////////////////////////////////////////////////////////////////
