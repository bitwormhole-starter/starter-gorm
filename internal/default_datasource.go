package internal

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
	"gorm.io/gorm"
)

// DefaultDatasource 默认的数据源
type DefaultDatasource struct {
	//starter:component
	_as func(libgorm.DataSource) //starter:as(".")

	Drivers libgorm.DriverManager //starter:inject("#")

	Name     string //starter:inject("${datasource.default.name}")
	Driver   string //starter:inject("${datasource.default.driver}")
	Host     string //starter:inject("${datasource.default.host}")
	Port     int    //starter:inject("${datasource.default.port}")
	User     string //starter:inject("${datasource.default.username}")
	Password string //starter:inject("${datasource.default.password}")
	Database string //starter:inject("${datasource.default.database}")
	Enabled  bool   //starter:inject("${datasource.default.enabled}")

	db *gorm.DB
}

func (inst *DefaultDatasource) _impl() (libgorm.DataSource, application.Lifecycle) {
	return inst, inst
}

// Life ...
func (inst *DefaultDatasource) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.init,
	}
}

func (inst *DefaultDatasource) init() error {
	_, err := inst.DB()
	return err
}

// Registration ...
func (inst *DefaultDatasource) Registration() *libgorm.DataSourceRegistration {
	return &libgorm.DataSourceRegistration{
		Name:       inst.Name,
		DataSource: inst,
	}
}

// DB ...
func (inst *DefaultDatasource) DB() (*gorm.DB, error) {
	db := inst.db
	if db == nil {
		d, err := inst.load()
		if err != nil {
			return nil, err
		}
		db = d
		inst.db = d
	}
	return db, nil
}

func (inst *DefaultDatasource) load() (*gorm.DB, error) {

	cfg := inst.config()
	if !cfg.Enabled {
		return nil, fmt.Errorf("the data-source named '%s' is disabled", inst.Name)
	}

	driver, err := inst.Drivers.FindDriver(cfg.Driver)
	if err != nil {
		return nil, err
	}

	db, err := driver.Open(cfg)
	if err != nil {
		return nil, err
	}

	return db.DB(), nil
}

func (inst *DefaultDatasource) config() *libgorm.Configuration {
	c := &libgorm.Configuration{}

	c.Database = inst.Database
	c.Driver = inst.Driver
	c.Enabled = inst.Enabled
	c.Host = inst.Host
	c.Password = inst.Password
	c.Port = inst.Port
	c.User = inst.User

	return c
}
