package internal

import (
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
)

const (
	theDataSourcePropertyNamePrefix = "datasource."
	theDataSourcePropertyNameSuffix = ".database"
)

// DefaultDatasourceRegistry 默认的数据源
type DefaultDatasourceRegistry struct {

	//starter:component
	_as func(libgorm.DataSourceRegistry) //starter:as(".")

	AC                 application.Context   //starter:inject("context")
	Drivers            libgorm.DriverManager //starter:inject("#")
	DataSourceNameList string                //starter:inject("${data.sources}")

}

func (inst *DefaultDatasourceRegistry) _impl() {
	inst._as(inst)
}

// ListSources ...
func (inst *DefaultDatasourceRegistry) ListSources() []*libgorm.DataSourceRegistration {
	dst := make([]*libgorm.DataSourceRegistration, 0)
	aliases := inst.listDataSourceAliases()
	for _, alias := range aliases {
		ds, err := inst.loadSource(alias)
		if err != nil {
			panic(err)
		}
		if ds.Enabled {
			dst = append(dst, ds)
		}
	}
	return dst
}

func (inst *DefaultDatasourceRegistry) loadSource(alias string) (*libgorm.DataSourceRegistration, error) {

	prefix := theDataSourcePropertyNamePrefix + alias + "."
	getter := inst.AC.GetProperties().Getter().Required()
	enabled := getter.GetBool(prefix + "enabled")

	info := &libgorm.DataSourceRegistration{}
	info.Alias = alias
	info.Enabled = enabled

	cfg := &info.Configuration
	cfg.Driver = getter.GetString(prefix + "driver")
	cfg.User = getter.GetString(prefix + "username")
	cfg.Password = getter.GetString(prefix + "password")
	cfg.Host = getter.GetString(prefix + "host")
	cfg.Port = getter.GetInt(prefix + "port")
	cfg.Database = getter.GetString(prefix + "database")
	cfg.Enabled = enabled

	err := getter.Error()
	if err != nil {
		return nil, err
	}

	err = inst.openSource(info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (inst *DefaultDatasourceRegistry) openSource(reg *libgorm.DataSourceRegistration) error {
	if !reg.Enabled {
		return nil
	}
	driver, err := inst.Drivers.FindDriver(reg.Configuration.Driver)
	if err != nil {
		return err
	}
	source := &defaultDatasource{
		config: reg.Configuration,
		driver: driver,
	}
	_, err = source.DB()
	if err != nil {
		return err
	}
	reg.DataSource = source
	return nil
}

func (inst *DefaultDatasourceRegistry) listDataSourceAliases() []string {
	dst := make([]string, 0)
	src := strings.Split(inst.DataSourceNameList, ",")
	for _, name := range src {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		dst = append(dst, name)
	}
	return dst
}
