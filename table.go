package libgorm

// // TableContext 是表格组的初始化参数
// type TableContext struct {
// 	Database Database
// }

// TableManager 是全局的表格管理器
type TableManager interface {
	ListAll() []*TableRegistration
}

// TableRegistration 表示一个表格的注册信息
type TableRegistration struct {
	SimpleName string
	Group      *Group
	Prototype  any // 原型
}

// EmptyTable 是一个用于占位的空表
type EmptyTable struct {
	ID int
}
