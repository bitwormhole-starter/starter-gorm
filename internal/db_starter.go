package internal

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/vlog"
)

// DatabaseStarter 数据库模块启动器
type DatabaseStarter struct {

	//starter:component
	_as func(libgorm.TableManager) //starter:as("#")

	Sources libgorm.DataSourceManager //starter:inject("#")
	Groups  libgorm.GroupManager      //starter:inject("#")

	AutoMigrate bool //starter:inject("${data.enable-auto-migrate}")

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
	return inst.logAll()
}

func (inst *DatabaseStarter) logAll() error {
	err := inst.listSources()
	if err != nil {
		return err
	}
	return inst.listGroups()
}

func (inst *DatabaseStarter) listSources() error {
	list := inst.Sources.ListAliases()
	for i, alias := range list {
		_, err := inst.Sources.GetDataSource(alias)
		if err != nil {
			return err
		}
		vlog.Info("libgorm.data.source[%d].alias: %s", i, alias)
	}
	return nil
}

func (inst *DatabaseStarter) listGroups() error {
	list := inst.Groups.ListGroups()
	for i, g := range list {
		vlog.Info("libgorm.data.group[%d].alias: %s", i, g.Alias)
	}
	return nil
}
