package gen4demo1
import (
    p512a30914 "github.com/starter-go/libgorm"
    p1b6e41313 "github.com/starter-go/libgorm/src/demo/demo1"
     "github.com/starter-go/application"
)

// type p1b6e41313.TaDaoImpl in package:github.com/starter-go/libgorm/src/demo/demo1
//
// id:com-1b6e41313c90dcc6-demo1-TaDaoImpl
// class:
// alias:
// scope:singleton
//
type p1b6e41313c_demo1_TaDaoImpl struct {
}

func (inst* p1b6e41313c_demo1_TaDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1b6e41313c90dcc6-demo1-TaDaoImpl"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1b6e41313c_demo1_TaDaoImpl) new() any {
    return &p1b6e41313.TaDaoImpl{}
}

func (inst* p1b6e41313c_demo1_TaDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1b6e41313.TaDaoImpl)
	nop(ie, com)

	
    com.Src = inst.getSrc(ie)


    return nil
}


func (inst*p1b6e41313c_demo1_TaDaoImpl) getSrc(ie application.InjectionExt)p1b6e41313.MyAgent{
    return ie.GetComponent("#alias-1b6e41313c90dcc6ce0464e19d3529ce-MyAgent").(p1b6e41313.MyAgent)
}



// type p1b6e41313.TbDaoImpl in package:github.com/starter-go/libgorm/src/demo/demo1
//
// id:com-1b6e41313c90dcc6-demo1-TbDaoImpl
// class:
// alias:
// scope:singleton
//
type p1b6e41313c_demo1_TbDaoImpl struct {
}

func (inst* p1b6e41313c_demo1_TbDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1b6e41313c90dcc6-demo1-TbDaoImpl"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1b6e41313c_demo1_TbDaoImpl) new() any {
    return &p1b6e41313.TbDaoImpl{}
}

func (inst* p1b6e41313c_demo1_TbDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1b6e41313.TbDaoImpl)
	nop(ie, com)

	
    com.Src = inst.getSrc(ie)


    return nil
}


func (inst*p1b6e41313c_demo1_TbDaoImpl) getSrc(ie application.InjectionExt)p1b6e41313.MyAgent{
    return ie.GetComponent("#alias-1b6e41313c90dcc6ce0464e19d3529ce-MyAgent").(p1b6e41313.MyAgent)
}



// type p1b6e41313.TcDaoImpl in package:github.com/starter-go/libgorm/src/demo/demo1
//
// id:com-1b6e41313c90dcc6-demo1-TcDaoImpl
// class:
// alias:
// scope:singleton
//
type p1b6e41313c_demo1_TcDaoImpl struct {
}

func (inst* p1b6e41313c_demo1_TcDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1b6e41313c90dcc6-demo1-TcDaoImpl"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1b6e41313c_demo1_TcDaoImpl) new() any {
    return &p1b6e41313.TcDaoImpl{}
}

func (inst* p1b6e41313c_demo1_TcDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1b6e41313.TcDaoImpl)
	nop(ie, com)

	
    com.Src = inst.getSrc(ie)


    return nil
}


func (inst*p1b6e41313c_demo1_TcDaoImpl) getSrc(ie application.InjectionExt)p1b6e41313.MyAgent{
    return ie.GetComponent("#alias-1b6e41313c90dcc6ce0464e19d3529ce-MyAgent").(p1b6e41313.MyAgent)
}



// type p1b6e41313.TableReg in package:github.com/starter-go/libgorm/src/demo/demo1
//
// id:com-1b6e41313c90dcc6-demo1-TableReg
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:alias-1b6e41313c90dcc6ce0464e19d3529ce-MyAgent
// scope:singleton
//
type p1b6e41313c_demo1_TableReg struct {
}

func (inst* p1b6e41313c_demo1_TableReg) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1b6e41313c90dcc6-demo1-TableReg"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = "alias-1b6e41313c90dcc6ce0464e19d3529ce-MyAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1b6e41313c_demo1_TableReg) new() any {
    return &p1b6e41313.TableReg{}
}

func (inst* p1b6e41313c_demo1_TableReg) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1b6e41313.TableReg)
	nop(ie, com)

	
    com.DSMan = inst.getDSMan(ie)
    com.Prefix = inst.getPrefix(ie)
    com.URI = inst.getURI(ie)
    com.SourceName = inst.getSourceName(ie)


    return nil
}


func (inst*p1b6e41313c_demo1_TableReg) getDSMan(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}


func (inst*p1b6e41313c_demo1_TableReg) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("datagroup.demo1.table-name-prefix")
}


func (inst*p1b6e41313c_demo1_TableReg) getURI(ie application.InjectionExt)string{
    return ie.GetString("datagroup.demo1.uri")
}


func (inst*p1b6e41313c_demo1_TableReg) getSourceName(ie application.InjectionExt)string{
    return ie.GetString("datagroup.demo1.datasource")
}


