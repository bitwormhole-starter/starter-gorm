package libgorm

// TableGroup 表示一组相关的表格
type TableGroup struct {

	// Namespace  string                  // group 的命名空间

	Enabled    bool                    // 指出是否启用该 group
	Name       string                  // group 的名称
	Prefix     string                  // group 的表名前缀
	OnInit     func(ctx *TableContext) // OnInit 函数用于初始化这个表格组
	Prototypes []any                   // 列出各种 entity 的原型
}

// TableGroupRegistry 是一个自动化的表格注册器
type TableGroupRegistry interface {

	// 取分组信息
	Group() *TableGroup
}

// GroupManager 分组管理器
type GroupManager interface {
	ListGroups() []*TableGroup
}
