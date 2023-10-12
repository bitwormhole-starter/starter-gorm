package internal

// type tableNamerChain struct {
// 	chain []*libgorm.TableNamerRegistration
// }

// func (inst *tableNamerChain) _impl() libgorm.TableNamer {
// 	return inst
// }

// func (inst *tableNamerChain) init(src []libgorm.TableNamerRegistry) {
// 	inst.chain = nil
// 	for _, r1 := range src {
// 		r2 := r1.Registration()
// 		if r2 == nil {
// 			continue
// 		}
// 		inst.chain = append(inst.chain, r2)
// 	}
// 	inst.sort()
// }

// func (inst *tableNamerChain) sort() {
// 	list := inst.chain
// 	s := &util.Sorter{}
// 	s.OnLen = func() int { return len(list) }
// 	s.OnSwap = func(i1, i2 int) { list[i1], list[i2] = list[i2], list[i1] }
// 	s.OnLess = func(i1, i2 int) bool {
// 		p1 := list[i1].Priority
// 		p2 := list[i2].Priority
// 		return p1 > p2
// 	}
// 	s.Sort()
// 	inst.chain = list
// }

// func (inst *tableNamerChain) GetName(namespace, simpleName string) string {
// 	for _, r := range inst.chain {
// 		if r == nil {
// 			continue
// 		}
// 		namer := r.Namer
// 		if namer == nil {
// 			continue
// 		}
// 		fullname := namer.GetName(namespace, simpleName)
// 		if fullname != "" {
// 			return fullname
// 		}
// 	}
// 	return "table_" + simpleName
// }
