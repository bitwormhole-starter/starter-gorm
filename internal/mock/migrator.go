package mock

import (
	"github.com/starter-go/vlog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type migrator struct{}

func (inst *migrator) _impl() gorm.Migrator {
	return inst
}

func (inst *migrator) AutoMigrate(dst ...interface{}) error {
	inst.log("AutoMigrate")
	return nil
}

// Database
func (inst *migrator) CurrentDatabase() string {
	inst.log("CurrentDatabase")
	return ""
}

func (inst *migrator) FullDataTypeOf(*schema.Field) clause.Expr {
	inst.log("FullDataTypeOf")
	return clause.Expr{}
}

func (inst *migrator) GetTypeAliases(databaseTypeName string) []string {
	inst.log("GetTypeAliases")
	return nil
}

// Tables
func (inst *migrator) CreateTable(dst ...interface{}) error {
	inst.log("CreateTable")
	return nil
}

func (inst *migrator) DropTable(dst ...interface{}) error {
	inst.log("DropTable")
	return nil
}

func (inst *migrator) HasTable(dst interface{}) bool {
	inst.log("HasTable")
	return false
}

func (inst *migrator) RenameTable(oldName, newName interface{}) error {
	inst.log("RenameTable")
	return nil
}

func (inst *migrator) GetTables() (tableList []string, err error) {
	inst.log("GetTables")
	return nil, nil
}

func (inst *migrator) TableType(dst interface{}) (gorm.TableType, error) {
	inst.log("TableType")
	return nil, nil
}

// Columns
func (inst *migrator) AddColumn(dst interface{}, field string) error {
	inst.log("AddColumn")
	return nil
}

func (inst *migrator) DropColumn(dst interface{}, field string) error {
	inst.log("DropColumn")
	return nil
}

func (inst *migrator) AlterColumn(dst interface{}, field string) error {
	inst.log("AlterColumn")
	return nil
}

func (inst *migrator) MigrateColumn(dst interface{}, field *schema.Field, columnType gorm.ColumnType) error {
	inst.log("MigrateColumn")
	return nil
}

func (inst *migrator) HasColumn(dst interface{}, field string) bool {
	inst.log("HasColumn")
	return false
}

func (inst *migrator) RenameColumn(dst interface{}, oldName, field string) error {
	inst.log("RenameColumn")
	return nil
}

func (inst *migrator) ColumnTypes(dst interface{}) ([]gorm.ColumnType, error) {
	inst.log("ColumnTypes")
	return nil, nil
}

// Views
func (inst *migrator) CreateView(name string, option gorm.ViewOption) error {
	inst.log("CreateView")
	return nil
}

func (inst *migrator) DropView(name string) error {
	inst.log("DropView")
	return nil
}

// Constraints
func (inst *migrator) CreateConstraint(dst interface{}, name string) error {
	inst.log("CreateConstraint")
	return nil
}

func (inst *migrator) DropConstraint(dst interface{}, name string) error {
	inst.log("DropConstraint")
	return nil
}

func (inst *migrator) HasConstraint(dst interface{}, name string) bool {
	inst.log("HasConstraint")
	return false
}

// Indexes
func (inst *migrator) CreateIndex(dst interface{}, name string) error {
	inst.log("CreateIndex")
	return nil
}

func (inst *migrator) DropIndex(dst interface{}, name string) error {
	inst.log("DropIndex")
	return nil
}

func (inst *migrator) HasIndex(dst interface{}, name string) bool {
	inst.log("HasIndex")
	return false
}

func (inst *migrator) RenameIndex(dst interface{}, oldName, newName string) error {
	inst.log("RenameIndex")
	return nil
}

func (inst *migrator) GetIndexes(dst interface{}) ([]gorm.Index, error) {
	inst.log("GetIndexes")
	return nil, nil
}

func (inst *migrator) log(tag string) {
	vlog.Debug("migrator.invoke:" + tag)
}
