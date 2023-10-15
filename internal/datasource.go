package internal

import (
	"github.com/starter-go/libgorm"
	"gorm.io/gorm"
)

// defaultDatasource 默认的数据源
type defaultDatasource struct {
	config libgorm.Configuration
	driver libgorm.Driver
	db     *gorm.DB
}

func (inst *defaultDatasource) _impl() libgorm.DataSource {
	return inst
}

// DB ...
func (inst *defaultDatasource) DB() (*gorm.DB, error) {
	db := inst.db
	if db != nil {
		return db, nil
	}
	db, err := inst.load()
	if err != nil {
		return nil, err
	}
	inst.db = db
	return db, nil
}

func (inst *defaultDatasource) load() (*gorm.DB, error) {
	cfg := &inst.config
	holder, err := inst.driver.Open(cfg)
	if err != nil {
		return nil, err
	}
	db := holder.DB()
	return db, nil
}
