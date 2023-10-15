package internal

import (
	"fmt"
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
)

// GroupManagerImpl ...
type GroupManagerImpl struct {

	//starter:component
	_as func(libgorm.GroupManager) //starter:as("#")

	AC                 application.Context       //starter:inject("context")
	GroupAliases       string                    //starter:inject("${data.groups}")
	AutoMigrateEnabled bool                      //starter:inject("${data.enable-auto-migrate}")
	DataSources        libgorm.DataSourceManager //starter:inject("#")
	GroupRegistries    []libgorm.GroupRegistry   //starter:inject(".")

	cached *groupsCache
}

func (inst *GroupManagerImpl) _impl() libgorm.GroupManager {
	return inst
}

// ListGroups ...
func (inst *GroupManagerImpl) ListGroups() []*libgorm.GroupRegistration {
	cache := inst.cached
	if cache == nil {
		c, err := inst.load()
		if err != nil {
			panic(err)
		}
		cache = c
		inst.cached = c
	}
	return cache.groups
}

func (inst *GroupManagerImpl) load() (*groupsCache, error) {

	r1list := inst.GroupRegistries
	r2list := make([]*libgorm.GroupRegistration, 0)

	for _, r1 := range r1list {
		list := r1.Groups()
		r2list = append(r2list, list...)
	}

	loader := &groupsLoader{}
	loader.dsManager = inst.DataSources
	loader.enableAutoMigrate = inst.AutoMigrateEnabled
	loader.context = inst.AC
	loader.groups = r2list
	loader.aliases = strings.Split(inst.GroupAliases, ",")

	return loader.load()
}

////////////////////////////////////////////////////////////////////////////////

type groupsLoader struct {
	context           application.Context
	dsManager         libgorm.DataSourceManager
	enableAutoMigrate bool
	aliases           []string
	groups            []*libgorm.GroupRegistration
}

func (inst *groupsLoader) load() (*groupsCache, error) {
	c := &groupsCache{}
	groups := c.groups
	for _, alias := range inst.aliases {
		alias = strings.TrimSpace(alias)
		if alias == "" {
			continue
		}
		g, err := inst.loadGroup(alias)
		if err != nil {
			return nil, err
		}
		if g.Enabled {
			groups = append(groups, g)
		}
	}
	c.groups = groups
	return c, nil
}

func (inst *groupsLoader) loadGroup(alias string) (*libgorm.GroupRegistration, error) {

	getter := inst.context.GetProperties().Getter().Required()
	prefix := "datagroup." + alias + "."

	g := &libgorm.GroupRegistration{}
	g.Alias = alias
	g.Enabled = getter.GetBool(prefix + "enabled")
	g.Prefix = getter.GetString(prefix + "table-name-prefix")
	g.URI = getter.GetString(prefix + "uri")
	g.Source = getter.GetString(prefix + "datasource")

	err := getter.Error()
	if err != nil {
		return nil, err
	}

	g2, err := inst.findGroupByURI(g.URI)
	if err != nil {
		return nil, err
	}

	g.Group = g2
	err = inst.autoMigrate(g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (inst *groupsLoader) autoMigrate(g *libgorm.GroupRegistration) error {

	if !g.Enabled {
		return nil
	}

	ds, err := inst.dsManager.GetDataSource(g.Source)
	if err != nil {
		return err
	}

	db, err := ds.DB()
	if err != nil {
		return err
	}

	if !inst.enableAutoMigrate {
		return nil
	}
	entities := g.Group.Prototypes()
	return db.AutoMigrate(entities...)
}

func (inst *groupsLoader) findGroupByURI(uri string) (libgorm.Group, error) {
	src := inst.groups
	for _, item := range src {
		if uri == "" {
			continue
		}
		if item.URI == uri && item.Group != nil {
			return item.Group, nil
		}
	}
	return nil, fmt.Errorf("libgorm: group is not found, by URI [%s]", uri)
}

////////////////////////////////////////////////////////////////////////////////

type groupsCache struct {
	groups []*libgorm.GroupRegistration
}
