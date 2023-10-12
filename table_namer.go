package libgorm

////////////////////////////////////////////////////////////////////////////////

// // TableNameGetter 用于从实体原型获取表名
// type TableNameGetter interface {
// 	TableName() string
// }

// // PrefixSetter 用于设置分组的表名前缀
// type PrefixSetter func(prefix string)

////////////////////////////////////////////////////////////////////////////////

// // TableNamer 是一个名称映射器，它根据传入的 namespace & simpleName 返回全名
// type TableNamer interface {
// 	GetName(namespace, simpleName string) string
// }

// // TableNamerRegistration 是 TableNamer 的注册信息
// type TableNamerRegistration struct {
// 	Priority int // 是这个 TableNamer 的优先级，数值越大越先处理
// 	Namer    TableNamer
// }

// // TableNamerRegistry 表示一个 TableNamer 的注册器
// type TableNamerRegistry interface {
// 	Registration() *TableNamerRegistration
// }

////////////////////////////////////////////////////////////////////////////////

// // TableNameCache ...
// type TableNameCache struct {
// 	namer      TableNamer
// 	namespace  string
// 	fullname   string
// 	simpleName string
// }

// // Init ...
// func (inst *TableNameCache) Init(namer TableNamer, ns string, simpleName string) {
// 	inst.namer = namer
// 	inst.namespace = ns
// 	inst.simpleName = simpleName
// 	inst.fullname = ""
// }

// func (inst *TableNameCache) findName() string {
// 	namer := inst.namer
// 	key := inst.simpleName
// 	ns := inst.namespace
// 	name := ""
// 	if namer != nil {
// 		name = namer.GetName(ns, key)
// 	}
// 	if name == "" {
// 		name = "unnamed"
// 	}
// 	return name
// }

// // Name 取表名
// func (inst *TableNameCache) Name() string {
// 	name := inst.fullname
// 	if name == "" {
// 		name = inst.findName()
// 		inst.fullname = name
// 	}
// 	return name
// }
