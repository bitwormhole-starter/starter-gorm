package libgorm

import "gorm.io/gorm"

// Database 数据库实例
type Database interface {
	DB() *gorm.DB
}

////////////////////////////////////////////////////////////////////////////////

// DatabaseBuilder 用来创建 Database 实例
type DatabaseBuilder struct {
	DB *gorm.DB
}

// Create 创建 Database 实例
func (inst *DatabaseBuilder) Create() Database {
	return &database{db: inst.DB}
}

////////////////////////////////////////////////////////////////////////////////

type database struct {
	db *gorm.DB
}

func (inst *database) DB() *gorm.DB {
	return inst.db
}

////////////////////////////////////////////////////////////////////////////////
