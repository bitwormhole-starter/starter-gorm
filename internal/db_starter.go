package internal

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
)

type loading struct {
	db     libgorm.Database
	source *libgorm.DataSourceRegistration
	group  *libgorm.TableGroup
}

////////////////////////////////////////////////////////////////////////////////

// DatabaseStarter 数据库模块启动器
type DatabaseStarter struct {

	//starter:component
	_as func(libgorm.TableManager) //starter:as("#")

	Sources libgorm.DataSourceManager //starter:inject("#")
	Groups  libgorm.GroupManager      //starter:inject("#")

	AutoMigrate bool //starter:inject("${libgorm.auto-migrate.enabled}")

	// GlobalTableNamePrefix string //starter:inject("${libgorm.auto-migrate.table-name-prefix}")
	// SourceName            string //starter:inject("${libgorm.auto-migrate.datasource}")

}

func (inst *DatabaseStarter) _impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *DatabaseStarter) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.init,
	}
}

func (inst *DatabaseStarter) init() error {
	return inst.loadAll()
}

func (inst *DatabaseStarter) loadAll() error {
	l := &loading{}
	dsm := inst.Sources
	srcNameList := dsm.ListNames()
	for _, name := range srcNameList {
		ds, err := dsm.GetDataSource(name)
		if err != nil {
			return err
		}
		l.source = ds.Registration()
		l.group = nil
		l.db = nil
		err = inst.loadSource(l)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *DatabaseStarter) loadSource(l *loading) error {
	ds := l.source
	if !ds.Enabled {
		return nil
	}

	db, err := ds.DataSource.DB()
	if err != nil {
		return err
	}
	dbb := libgorm.DatabaseBuilder{DB: db}
	l.db = dbb.Create()

	groups := ds.Groups
	for _, groupName := range groups {
		group, err := inst.findGroup(groupName)
		if err != nil {
			return err
		}
		l.group = group
		err = inst.loadGroup(l)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *DatabaseStarter) loadGroup(l *loading) error {
	g := l.group
	if !g.Enabled {
		return nil
	}
	ctx := &libgorm.TableContext{}
	ctx.Database = l.db
	fnOnInit := g.OnInit
	if fnOnInit != nil {
		fnOnInit(ctx)
	}
	return inst.migrate(l)
}

func (inst *DatabaseStarter) findGroup(name1 string) (*libgorm.TableGroup, error) {
	glist := inst.Groups.ListGroups()
	for _, g := range glist {
		name2 := g.Name
		if name1 == name2 {
			return g, nil
		}
	}
	return nil, fmt.Errorf("no data-table-group with name: %s", name1)
}

func (inst *DatabaseStarter) migrate(l *loading) error {
	if !inst.AutoMigrate {
		return nil
	}
	group := l.group
	list := group.Prototypes
	db := l.db.DB()
	return db.AutoMigrate(list...)
}
