package libgorm

// // TableGroupBuilder 用于创建一组相关的表格信息
// type TableGroupBuilder struct {
// 	group *TableGroup
// 	list  []*TableRegistration
// }

// // Group 设置分组信息
// func (inst *TableGroupBuilder) Group(g *TableGroup) *TableGroupBuilder {
// 	if g != nil {
// 		inst.group = g
// 	}
// 	return inst
// }

// // Table 添加一个表格
// func (inst *TableGroupBuilder) Table(t *TableRegistration) *TableGroupBuilder {
// 	if t != nil {
// 		inst.list = append(inst.list, t)
// 	}
// 	return inst
// }

// // Entity 以 Entity 的形式添加一个表格
// func (inst *TableGroupBuilder) Entity(ent any) *TableGroupBuilder {

// 	getter, ok := ent.(TableNameGetter)
// 	if !ok {
// 		vlog.Warn("no TableNameGetter for entity '%v', ignored", ent)
// 		return inst
// 	}

// 	name := getter.TableName()

// 	tr := &TableRegistration{}
// 	tr.Group = inst.group
// 	tr.NameGetter = getter.TableName
// 	tr.SimpleName = name
// 	tr.PrototypeGetter = func() any { return ent }

// 	return inst.Table(tr)
// }

// // Create 创建表格组
// func (inst *TableGroupBuilder) Create() []*TableRegistration {

// 	g := inst.group
// 	all := inst.list

// 	if g == nil {
// 		g = &TableGroup{}
// 	}

// 	if g.NamerSetter == nil {
// 		g.NamerSetter = inst.nopSetter()
// 	}

// 	for _, item := range all {
// 		if item.PrototypeGetter == nil {
// 			item.PrototypeGetter = inst.nopGetter()
// 		}
// 		item.Group = g
// 	}

// 	return all
// }

// func (inst *TableGroupBuilder) nopGetter() func() any {
// 	return func() any { return &EmptyTable{} }
// }

// func (inst *TableGroupBuilder) nopSetter() func(namer TableNamer) {
// 	return func(namer TableNamer) {}
// }
