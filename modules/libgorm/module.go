package libgorm

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/libgorm/gen/gen4libgorm"
)

// Module 导出模块 ['github.com/starter-go/libgorm']
func Module() application.Module {
	mb := libgorm.ModuleBuilder()
	mb.Components(gen4libgorm.ExportConfigForLibgorm)
	return mb.Create()
}
