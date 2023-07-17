package internal

import (
	"sort"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/vlog"
)

// TableNameGetter 用于从实体原型获取表名
type TableNameGetter interface {
	TableName() string
}

////////////////////////////////////////////////////////////////////////////////

// DefaultTableManager 是默认的表格管理器
type DefaultTableManager struct {

	//starter:component
	_as func(libgorm.TableManager) //starter:as("#")

	TRs         []libgorm.TableRegistry   //starter:inject(".")
	DataSources libgorm.DataSourceManager //starter:inject("#")

	GlobalTableNamePrefix string //starter:inject("${libgorm.auto-migrate.table-name-prefix}")
	AutoMigrate           bool   //starter:inject("${libgorm.auto-migrate.enabled}")
	SourceName            string //starter:inject("${libgorm.auto-migrate.datasource}")

	cached []*libgorm.TableRegistration
}

func (inst *DefaultTableManager) _impl() application.Lifecycle {
	inst._as(inst)
	return inst
}

// Life ...
func (inst *DefaultTableManager) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.init,
	}
}

func (inst *DefaultTableManager) init() error {
	if inst.AutoMigrate {
		return inst.migrate()
	}
	return nil
}

func (inst *DefaultTableManager) migrate() error {

	all, err := inst.getAll()
	if err != nil {
		return err
	}

	ds, err := inst.DataSources.GetDataSource(inst.SourceName)
	if err != nil {
		return err
	}

	db, err := ds.DB()
	if err != nil {
		return err
	}

	dst := make([]any, 0)
	namelist := make([]string, 0)

	for _, tr := range all {
		pt := tr.PrototypeGetter()
		dst = append(dst, pt)
		name := inst.getTableName(tr, pt)
		namelist = append(namelist, name)
	}

	sort.Strings(namelist)
	for _, name := range namelist {
		vlog.Info("auto-migrate table: '%s'", name)
	}

	return db.AutoMigrate(dst...)
}

func (inst *DefaultTableManager) getTableName(tr *libgorm.TableRegistration, o any) string {
	getter, ok := o.(TableNameGetter)
	if ok {
		return getter.TableName()
	}
	return tr.SimpleName
}

// ListAll ...
func (inst *DefaultTableManager) ListAll() []*libgorm.TableRegistration {
	all, _ := inst.getAll()
	if all == nil {
		all = make([]*libgorm.TableRegistration, 0)
	}
	return all
}

func (inst *DefaultTableManager) getAll() ([]*libgorm.TableRegistration, error) {
	all := inst.cached
	if all == nil {
		list, err := inst.load()
		if err != nil {
			return nil, err
		}
		all = list
		inst.cached = all
	}
	return all, nil
}

func (inst *DefaultTableManager) load() ([]*libgorm.TableRegistration, error) {
	src := inst.TRs
	dst := make([]*libgorm.TableRegistration, 0)
	for _, r1 := range src {
		r2list := r1.ListTableRegistrations()
		for _, r2 := range r2list {
			inst.prepareGroup(r2.Group)
			dst = append(dst, r2)
		}
	}
	return dst, nil
}

func (inst *DefaultTableManager) prepareGroup(g *libgorm.TableGroup) {
	prefix := inst.GlobalTableNamePrefix + g.Name
	g.PrefixSetter(prefix)
}
