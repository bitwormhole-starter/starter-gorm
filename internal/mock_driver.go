package internal

import (
	"github.com/starter-go/libgorm"
	"gorm.io/gorm"
)

// MockDriver 伪驱动，仅用于调试
type MockDriver struct {

	//starter:component
	_as func(libgorm.Driver) //starter:as(".")
}

func (inst *MockDriver) _impl() libgorm.Driver {
	return inst
}

// Registration 获取注册信息
func (inst *MockDriver) Registration() *libgorm.DriverRegistration {
	return &libgorm.DriverRegistration{
		Name:   "mock",
		Driver: inst,
	}
}

// Open 打开数据库
func (inst *MockDriver) Open(c *libgorm.Configuration) (libgorm.Database, error) {
	db := &gorm.DB{}
	builder := &libgorm.DatabaseBuilder{
		DB: db,
	}
	res := builder.Create()
	return res, nil
}
