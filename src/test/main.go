package main

import (
	"embed"
	"os"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm/gen/gen4demo1"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(theModule())
	i.WithPanic(true).Run()
}

//go:embed "resources"
var theModuleResFS embed.FS

func theModule() application.Module {

	parent := libgorm.Module()

	mb := application.ModuleBuilder{}
	mb.Name(parent.Name() + "#test")
	mb.Version(parent.Version())
	mb.Revision(parent.Revision())
	mb.EmbedResources(theModuleResFS, "resources")
	mb.Components(gen4demo1.ExportConfigForDemo1)
	mb.Depend(parent)

	return mb.Create()
}
