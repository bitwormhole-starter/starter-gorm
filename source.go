package libgorm

import "gorm.io/gorm"

// DataSource 数据源
type DataSource interface {
	DB() (*gorm.DB, error)
	Registration() *DataSourceRegistration
}

// DataSourceRegistration 数据源注册信息
type DataSourceRegistration struct {
	Name       string
	DataSource DataSource
}

// DataSourceManager 数据源管理器
type DataSourceManager interface {
	GetDataSource(name string) (DataSource, error)
}

// Agent 是默认的数据源代理
type Agent interface {
	DB(db *gorm.DB) *gorm.DB
}
