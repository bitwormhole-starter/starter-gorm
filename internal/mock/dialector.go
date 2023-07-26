package mock

import (
	"github.com/starter-go/vlog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type dialector struct{}

func (inst *dialector) _impl() gorm.Dialector {
	return inst
}

func (inst *dialector) Name() string {
	inst.log("name")
	return ""
}

func (inst *dialector) Initialize(*gorm.DB) error {
	inst.log("Initialize")
	return nil
}

func (inst *dialector) Migrator(db *gorm.DB) gorm.Migrator {
	inst.log("Migrator")
	return &migrator{}
}

func (inst *dialector) DataTypeOf(*schema.Field) string {
	inst.log("DataTypeOf")
	return ""
}

func (inst *dialector) DefaultValueOf(*schema.Field) clause.Expression {
	inst.log("DefaultValueOf")
	return nil
}

func (inst *dialector) BindVarTo(writer clause.Writer, stmt *gorm.Statement, v interface{}) {
	inst.log("BindVarTo")
}

func (inst *dialector) QuoteTo(clause.Writer, string) {
	inst.log("QuoteTo")
}

func (inst *dialector) Explain(sql string, vars ...interface{}) string {
	inst.log("Explain")
	return ""
}

func (inst *dialector) log(tag string) {
	vlog.Debug("dialector.invoke:" + tag)
}
