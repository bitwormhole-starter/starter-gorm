package internal

import (
	"fmt"
	"strings"

	"github.com/starter-go/libgorm"
)

// DefaultDatasourceManager 默认的数据源管理器
type DefaultDatasourceManager struct {

	//starter:component
	_as func(libgorm.DataSourceManager) //starter:as("#")

	Sources []libgorm.DataSource //starter:inject(".")

}

func (inst *DefaultDatasourceManager) _impl() libgorm.DataSourceManager {
	return inst
}

// GetDataSource ...
func (inst *DefaultDatasourceManager) GetDataSource(name string) (libgorm.DataSource, error) {
	want := inst.normalizeName(name)
	src := inst.Sources
	for _, ds := range src {
		r := ds.Registration()
		have := inst.normalizeName(r.Name)
		if have == want {
			return ds, nil
		}
	}
	return nil, fmt.Errorf("no data-source with name: '%s'", want)
}

func (inst *DefaultDatasourceManager) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name
}
