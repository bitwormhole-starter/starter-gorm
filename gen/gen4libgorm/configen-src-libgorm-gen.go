package gen4libgorm
import (
    p0ef6f2938 "github.com/starter-go/application"
    p512a30914 "github.com/starter-go/libgorm"
    p6553b6c73 "github.com/starter-go/libgorm/internal"
    p273dbed52 "github.com/starter-go/libgorm/internal/mock"
     "github.com/starter-go/application"
)

// type p6553b6c73.DefaultDatasourceManager in package:github.com/starter-go/libgorm/internal
//
// id:com-6553b6c73b54fb77-internal-DefaultDatasourceManager
// class:
// alias:alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager
// scope:singleton
//
type p6553b6c73b_internal_DefaultDatasourceManager struct {
}

func (inst* p6553b6c73b_internal_DefaultDatasourceManager) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6553b6c73b54fb77-internal-DefaultDatasourceManager"
	r.Classes = ""
	r.Aliases = "alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6553b6c73b_internal_DefaultDatasourceManager) new() any {
    return &p6553b6c73.DefaultDatasourceManager{}
}

func (inst* p6553b6c73b_internal_DefaultDatasourceManager) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6553b6c73.DefaultDatasourceManager)
	nop(ie, com)

	
    com.Sources = inst.getSources(ie)


    return nil
}


func (inst*p6553b6c73b_internal_DefaultDatasourceManager) getSources(ie application.InjectionExt)[]p512a30914.DataSourceRegistry{
    dst := make([]p512a30914.DataSourceRegistry, 0)
    src := ie.ListComponents(".class-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceRegistry")
    for _, item1 := range src {
        item2 := item1.(p512a30914.DataSourceRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p6553b6c73.DefaultDatasourceRegistry in package:github.com/starter-go/libgorm/internal
//
// id:com-6553b6c73b54fb77-internal-DefaultDatasourceRegistry
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceRegistry
// alias:
// scope:singleton
//
type p6553b6c73b_internal_DefaultDatasourceRegistry struct {
}

func (inst* p6553b6c73b_internal_DefaultDatasourceRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6553b6c73b54fb77-internal-DefaultDatasourceRegistry"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6553b6c73b_internal_DefaultDatasourceRegistry) new() any {
    return &p6553b6c73.DefaultDatasourceRegistry{}
}

func (inst* p6553b6c73b_internal_DefaultDatasourceRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6553b6c73.DefaultDatasourceRegistry)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.Drivers = inst.getDrivers(ie)
    com.DataSourceNameList = inst.getDataSourceNameList(ie)


    return nil
}


func (inst*p6553b6c73b_internal_DefaultDatasourceRegistry) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p6553b6c73b_internal_DefaultDatasourceRegistry) getDrivers(ie application.InjectionExt)p512a30914.DriverManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DriverManager").(p512a30914.DriverManager)
}


func (inst*p6553b6c73b_internal_DefaultDatasourceRegistry) getDataSourceNameList(ie application.InjectionExt)string{
    return ie.GetString("${data.sources}")
}



// type p6553b6c73.DatabaseStarter in package:github.com/starter-go/libgorm/internal
//
// id:com-6553b6c73b54fb77-internal-DatabaseStarter
// class:
// alias:alias-512a309140d0ad99eb1c95c8dc0d02f9-TableManager
// scope:singleton
//
type p6553b6c73b_internal_DatabaseStarter struct {
}

func (inst* p6553b6c73b_internal_DatabaseStarter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6553b6c73b54fb77-internal-DatabaseStarter"
	r.Classes = ""
	r.Aliases = "alias-512a309140d0ad99eb1c95c8dc0d02f9-TableManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6553b6c73b_internal_DatabaseStarter) new() any {
    return &p6553b6c73.DatabaseStarter{}
}

func (inst* p6553b6c73b_internal_DatabaseStarter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6553b6c73.DatabaseStarter)
	nop(ie, com)

	
    com.Sources = inst.getSources(ie)
    com.Groups = inst.getGroups(ie)
    com.AutoMigrate = inst.getAutoMigrate(ie)


    return nil
}


func (inst*p6553b6c73b_internal_DatabaseStarter) getSources(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}


func (inst*p6553b6c73b_internal_DatabaseStarter) getGroups(ie application.InjectionExt)p512a30914.GroupManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-GroupManager").(p512a30914.GroupManager)
}


func (inst*p6553b6c73b_internal_DatabaseStarter) getAutoMigrate(ie application.InjectionExt)bool{
    return ie.GetBool("${data.enable-auto-migrate}")
}



// type p6553b6c73.DefaultDriverManager in package:github.com/starter-go/libgorm/internal
//
// id:com-6553b6c73b54fb77-internal-DefaultDriverManager
// class:
// alias:alias-512a309140d0ad99eb1c95c8dc0d02f9-DriverManager
// scope:singleton
//
type p6553b6c73b_internal_DefaultDriverManager struct {
}

func (inst* p6553b6c73b_internal_DefaultDriverManager) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6553b6c73b54fb77-internal-DefaultDriverManager"
	r.Classes = ""
	r.Aliases = "alias-512a309140d0ad99eb1c95c8dc0d02f9-DriverManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6553b6c73b_internal_DefaultDriverManager) new() any {
    return &p6553b6c73.DefaultDriverManager{}
}

func (inst* p6553b6c73b_internal_DefaultDriverManager) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6553b6c73.DefaultDriverManager)
	nop(ie, com)

	
    com.Drivers = inst.getDrivers(ie)


    return nil
}


func (inst*p6553b6c73b_internal_DefaultDriverManager) getDrivers(ie application.InjectionExt)[]p512a30914.Driver{
    dst := make([]p512a30914.Driver, 0)
    src := ie.ListComponents(".class-512a309140d0ad99eb1c95c8dc0d02f9-Driver")
    for _, item1 := range src {
        item2 := item1.(p512a30914.Driver)
        dst = append(dst, item2)
    }
    return dst
}



// type p6553b6c73.GroupManagerImpl in package:github.com/starter-go/libgorm/internal
//
// id:com-6553b6c73b54fb77-internal-GroupManagerImpl
// class:
// alias:alias-512a309140d0ad99eb1c95c8dc0d02f9-GroupManager
// scope:singleton
//
type p6553b6c73b_internal_GroupManagerImpl struct {
}

func (inst* p6553b6c73b_internal_GroupManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6553b6c73b54fb77-internal-GroupManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-512a309140d0ad99eb1c95c8dc0d02f9-GroupManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6553b6c73b_internal_GroupManagerImpl) new() any {
    return &p6553b6c73.GroupManagerImpl{}
}

func (inst* p6553b6c73b_internal_GroupManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6553b6c73.GroupManagerImpl)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.GroupAliases = inst.getGroupAliases(ie)
    com.AutoMigrateEnabled = inst.getAutoMigrateEnabled(ie)
    com.DataSources = inst.getDataSources(ie)
    com.GroupRegistries = inst.getGroupRegistries(ie)


    return nil
}


func (inst*p6553b6c73b_internal_GroupManagerImpl) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p6553b6c73b_internal_GroupManagerImpl) getGroupAliases(ie application.InjectionExt)string{
    return ie.GetString("${data.groups}")
}


func (inst*p6553b6c73b_internal_GroupManagerImpl) getAutoMigrateEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${data.enable-auto-migrate}")
}


func (inst*p6553b6c73b_internal_GroupManagerImpl) getDataSources(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}


func (inst*p6553b6c73b_internal_GroupManagerImpl) getGroupRegistries(ie application.InjectionExt)[]p512a30914.GroupRegistry{
    dst := make([]p512a30914.GroupRegistry, 0)
    src := ie.ListComponents(".class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry")
    for _, item1 := range src {
        item2 := item1.(p512a30914.GroupRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p273dbed52.Driver in package:github.com/starter-go/libgorm/internal/mock
//
// id:com-273dbed527f9f8b7-mock-Driver
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-Driver
// alias:
// scope:singleton
//
type p273dbed527_mock_Driver struct {
}

func (inst* p273dbed527_mock_Driver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-273dbed527f9f8b7-mock-Driver"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-Driver"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p273dbed527_mock_Driver) new() any {
    return &p273dbed52.Driver{}
}

func (inst* p273dbed527_mock_Driver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p273dbed52.Driver)
	nop(ie, com)

	


    return nil
}


