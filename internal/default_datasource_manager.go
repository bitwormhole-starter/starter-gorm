package internal

import (
	"fmt"
	"strings"

	"github.com/starter-go/libgorm"
	"github.com/starter-go/vlog"
)

// DefaultDatasourceManager 默认的数据源管理器
type DefaultDatasourceManager struct {

	//starter:component
	_as func(libgorm.DataSourceManager) //starter:as("#")

	Sources []libgorm.DataSource //starter:inject(".")

	cached []*libgorm.DataSourceRegistration
}

func (inst *DefaultDatasourceManager) _impl() libgorm.DataSourceManager {
	return inst
}

func (inst *DefaultDatasourceManager) loadAll() []*libgorm.DataSourceRegistration {
	src := inst.Sources
	dst := make([]*libgorm.DataSourceRegistration, 0)
	for _, ds := range src {
		r := ds.Registration()
		r.Name = inst.normalizeName(r.Name)
		dst = append(dst, r)
		vlog.Info("load data source with name: '%s'", r.Name)
	}
	return dst
}

func (inst *DefaultDatasourceManager) listAll() []*libgorm.DataSourceRegistration {
	dst := inst.cached
	if dst == nil {
		dst = inst.loadAll()
		inst.cached = dst
	}
	return dst
}

// GetDataSource ...
func (inst *DefaultDatasourceManager) GetDataSource(name string) (libgorm.DataSource, error) {
	want := inst.normalizeName(name)
	all := inst.listAll()
	for _, r1 := range all {
		if want == r1.Name {
			return r1.DataSource, nil
		}
	}
	return nil, fmt.Errorf("no data-source with name: '%s'", want)
}

// ListNames ...
func (inst *DefaultDatasourceManager) ListNames() []string {
	dst := make([]string, 0)
	src := inst.listAll()
	for _, ds := range src {
		name := ds.Name
		dst = append(dst, name)
	}
	return dst
}

func (inst *DefaultDatasourceManager) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name
}
