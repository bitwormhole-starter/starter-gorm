package libgorm

// TableGroup 表示一组相关的表格
type TableGroup struct {
	Name         string
	PrefixSetter func(prefix string) // 用于设置本分组的表格名称前缀
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

////////////////////////////////////////////////////////////////////////////////

// TableGroupBuilder 用于创建一组相关的表格信息
type TableGroupBuilder struct {
	group *TableGroup
	list  []*TableRegistration
}

// Group 设置分组信息
func (inst *TableGroupBuilder) Group(g *TableGroup) *TableGroupBuilder {
	if g != nil {
		inst.group = g
	}
	return inst
}

// Table 添加一个表格
func (inst *TableGroupBuilder) Table(t *TableRegistration) *TableGroupBuilder {
	if t != nil {
		inst.list = append(inst.list, t)
	}
	return inst
}

// Create 创建表格组
func (inst *TableGroupBuilder) Create() []*TableRegistration {

	g := inst.group
	all := inst.list

	if g == nil {
		g = &TableGroup{}
	}

	if g.PrefixSetter == nil {
		g.PrefixSetter = inst.nopSetter()
	}

	for _, item := range all {
		if item.PrototypeGetter == nil {
			item.PrototypeGetter = inst.nopGetter()
		}
		item.Group = g
	}

	return all
}

func (inst *TableGroupBuilder) nopGetter() func() any {
	return func() any { return &EmptyTable{} }
}

func (inst *TableGroupBuilder) nopSetter() func(x string) {
	return func(x string) {}
}

////////////////////////////////////////////////////////////////////////////////

// EmptyTable 是一个用于占位的空表
type EmptyTable struct {
	ID int
}
