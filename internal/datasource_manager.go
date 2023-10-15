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

	Sources []libgorm.DataSourceRegistry //starter:inject(".")

	cached []*libgorm.DataSourceRegistration
}

func (inst *DefaultDatasourceManager) _impl() libgorm.DataSourceManager {
	return inst
}

// GetDataSource ...
func (inst *DefaultDatasourceManager) GetDataSource(name string) (libgorm.DataSource, error) {
	want := inst.normalizeName(name)
	all := inst.getAll()
	for _, r1 := range all {
		if want == r1.Alias {
			return r1.DataSource, nil
		}
	}
	return nil, fmt.Errorf("no data-source with name: '%s'", want)
}

// ListAliases ...
func (inst *DefaultDatasourceManager) ListAliases() []string {
	dst := make([]string, 0)
	src := inst.getAll()
	for _, ds := range src {
		alias := ds.Alias
		dst = append(dst, alias)
	}
	return dst
}

func (inst *DefaultDatasourceManager) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name
}

func (inst *DefaultDatasourceManager) loadAll() []*libgorm.DataSourceRegistration {
	src := inst.Sources
	dst := make([]*libgorm.DataSourceRegistration, 0)
	for _, r1 := range src {
		list2 := r1.ListSources()
		// r.Name = inst.normalizeName(r.Name)
		// vlog.Info("load data source with name: '%s'", r.Name)
		dst = append(dst, list2...)
	}
	return dst
}

func (inst *DefaultDatasourceManager) getAll() []*libgorm.DataSourceRegistration {
	dst := inst.cached
	if dst == nil {
		dst = inst.loadAll()
		inst.cached = dst
	}
	return dst
}
