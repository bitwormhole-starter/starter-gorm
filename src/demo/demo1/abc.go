package demo1

import "github.com/starter-go/libgorm"

// TableA ...
type TableA struct {
	ID       int
	ValueInt int
}

// TableB ...
type TableB struct {
	ID        int
	ValueBool bool
}

// TableC ...
type TableC struct {
	ID       int
	ValueStr string
}

////////////////////////////////////////////////////////////////////////////////

// TableName ...
func (t TableA) TableName() string {
	return theTableNames.ta.Name()
}

// TableName ...
func (t TableB) TableName() string {
	return theTableNames.tb.Name()
}

// TableName ...
func (t TableC) TableName() string {
	return theTableNames.tc.Name()
}

////////////////////////////////////////////////////////////////////////////////

type names struct {
	namer libgorm.TableNamer
	ns    string

	ta libgorm.TableNameCache
	tb libgorm.TableNameCache
	tc libgorm.TableNameCache
}

var theTableNames names

const theTableNamespace = "libgorm/demo"
const theTableGroupName = "demo"

func (inst *names) initAll(namer libgorm.TableNamer) {

	inst.ns = theTableNamespace
	inst.namer = namer

	inst.initItem("a", &inst.ta)
	inst.initItem("b", &inst.tb)
	inst.initItem("c", &inst.tc)
}

func (inst *names) initItem(simpleName string, cache *libgorm.TableNameCache) {
	cache.Init(inst.namer, inst.ns, simpleName)
}

////////////////////////////////////////////////////////////////////////////////

// TableReg ...
type TableReg struct {

	//starter:component
	_as func(libgorm.TableRegistry) //starter:as(".")

}

func (inst *TableReg) _impl() {
	inst._as(inst)
}

// ListTableRegistrations ...
func (inst *TableReg) ListTableRegistrations() []*libgorm.TableRegistration {

	b := libgorm.TableGroupBuilder{}
	namex := &theTableNames

	b.Group(&libgorm.TableGroup{
		Name:        theTableGroupName,
		Namespace:   theTableNamespace,
		NamerSetter: namex.initAll,
	})

	b.Entity(&TableA{})
	b.Entity(&TableB{})
	b.Entity(&TableC{})

	return b.Create()
}
