package libgorm

// TableGroup 表示一组相关的表格
type TableGroup struct {
	Name        string
	Namespace   string
	NamerSetter func(namer TableNamer) // 用于设置本分组的表格命名者
}

// TableRegistration 表示一个表格的注册信息
type TableRegistration struct {
	SimpleName      string
	Group           *TableGroup
	PrototypeGetter func() any    // 用于获取原型
	NameGetter      func() string // 用于获取完整名称
}

// TableRegistry 是一个自动化的表格注册器
type TableRegistry interface {
	ListTableRegistrations() []*TableRegistration
}

// TableManager 是全局的表格管理器
type TableManager interface {
	ListAll() []*TableRegistration
}

// EmptyTable 是一个用于占位的空表
type EmptyTable struct {
	ID int
}
