package internal

import (
	"sync"

	"github.com/starter-go/application"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/safe"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/vlog"
)

// DefaultTableNamer 是默认的数据表命名工具
type DefaultTableNamer struct {

	//starter:component
	_as func(libgorm.TableNamerRegistry) //starter:as(".")

	Context               application.Context //starter:inject("context")
	TableNamesPropResPath string              //starter:inject("${libgorm.table-group-namespaces.properties}")

	mappings map[string]string // map[namespace] table_name_prefix
	mutex    sync.Mutex
}

func (inst *DefaultTableNamer) _impl() (libgorm.TableNamer, libgorm.TableNamerRegistry) {
	return inst, inst
}

// Registration ...
func (inst *DefaultTableNamer) Registration() *libgorm.TableNamerRegistration {
	return &libgorm.TableNamerRegistration{
		Priority: 0,
		Namer:    inst,
	}
}

// GetName ...
func (inst *DefaultTableNamer) GetName(namespace, simpleName string) string {
	prefix := inst.findPrefix(namespace)
	return prefix + simpleName
}

func (inst *DefaultTableNamer) findPrefix(namespace string) string {

	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	mappings := inst.mappings
	if mappings == nil {
		mappings = inst.loadMappings()
		inst.mappings = mappings
	}

	prefix := mappings[namespace]
	if prefix == "" {
		prefix = "t_"
		vlog.Warn("no namespace-to-prefix mapping for namespace: '%s'", namespace)
	}

	return prefix
}

func (inst *DefaultTableNamer) loadMappings() map[string]string {
	path := inst.TableNamesPropResPath
	res := inst.Context.GetResources()
	props, err := properties.LoadFromResource(path, res, safe.Fast())
	if err != nil {
		panic(err)
	}
	return props.Export(nil)
}
