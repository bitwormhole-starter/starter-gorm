package libgorm

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

const (
	theModuleName     = "github.com/starter-go/libgorm"
	theModuleVersion  = "v0.9.12"
	theModuleRevision = 13
	theModuleResPath  = "src/lib/resources"
)

//go:embed "src/lib/resources"
var theModuleResFS embed.FS

// ModuleBuilder 用于创建模块 ['github.com/starter-go/libgorm']
func ModuleBuilder() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Depend(starter.Module())

	return mb
}
