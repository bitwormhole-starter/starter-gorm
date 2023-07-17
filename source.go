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
