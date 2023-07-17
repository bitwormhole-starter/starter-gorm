package internal

import (
	"fmt"
	"strings"

	"github.com/starter-go/libgorm"
)

// DefaultDriverManager 默认的驱动管理器
type DefaultDriverManager struct {
	//starter:component
	_as func(libgorm.DriverManager) //starter:as("#")

	Drivers []libgorm.Driver //starter:inject(".")
}

func (inst *DefaultDriverManager) _impl() libgorm.DriverManager {
	return inst
}

// FindDriver 根据名称查找驱动
func (inst *DefaultDriverManager) FindDriver(name string) (libgorm.Driver, error) {
	src := inst.Drivers
	want := inst.normalizeName(name)
	for _, d := range src {
		r := d.Registration()
		have := inst.normalizeName(r.Name)
		if want == have {
			return d, nil
		}
	}
	return nil, fmt.Errorf("no data-source driver with name: '%s'", name)
}

func (inst *DefaultDriverManager) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name
}
