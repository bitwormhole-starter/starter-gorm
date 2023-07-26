package mock

import (
	"github.com/starter-go/libgorm"
	"gorm.io/gorm"
)

// Driver 伪驱动，仅用于调试
type Driver struct {

	//starter:component
	_as func(libgorm.Driver) //starter:as(".")
}

func (inst *Driver) _impl() libgorm.Driver {
	return inst
}

// Registration 获取注册信息
func (inst *Driver) Registration() *libgorm.DriverRegistration {
	return &libgorm.DriverRegistration{
		Name:   "mock",
		Driver: inst,
	}
}

// Open 打开数据库
func (inst *Driver) Open(c *libgorm.Configuration) (libgorm.Database, error) {
	db, err := inst.openGormDB(c)
	if err != nil {
		return nil, err
	}
	dbb := &libgorm.DatabaseBuilder{DB: db}
	return dbb.Create(), nil
}

func (inst *Driver) openGormDB(c *libgorm.Configuration) (*gorm.DB, error) {
	d := &dialector{}
	return gorm.Open(d)
}
