package libgorm

// Group 表示一组相关的表格
type Group interface {

	// Namespace  string                  // group 的命名空间
	// OnInit     func(ctx *TableContext) // OnInit 函数用于初始化这个表格组

	Prototypes() []any // 列出各种 entity 的原型

}

// GroupRegistration 分组的注册信息
type GroupRegistration struct {
	Enabled bool   // 指出是否启用该 group
	Alias   string // group 的别名 (短名称)
	URI     string // group 的ID名称 (长名称)
	Prefix  string // group 的表名前缀
	Source  string // 数据源的 alias

	Group Group
}

// GroupRegistry 是一个自动化的表格注册器
type GroupRegistry interface {

	// 取分组信息
	Groups() []*GroupRegistration
}

// GroupManager 分组管理器
type GroupManager interface {
	ListGroups() []*GroupRegistration
}
