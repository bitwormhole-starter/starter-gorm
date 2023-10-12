package gen4libgorm
import (
    p512a30914 "github.com/starter-go/libgorm"
    p6553b6c73 "github.com/starter-go/libgorm/internal"
    p273dbed52 "github.com/starter-go/libgorm/internal/mock"
     "github.com/starter-go/application"
)

// type p6553b6c73.DataSourceAgent in package:github.com/starter-go/libgorm/internal
//
// id:com-6553b6c73b54fb77-internal-DataSourceAgent
// class:
// alias:alias-512a309140d0ad99eb1c95c8dc0d02f9-Agent
// scope:singleton
//
type p6553b6c73b_internal_DataSourceAgent struct {
}

func (inst* p6553b6c73b_internal_DataSourceAgent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6553b6c73b54fb77-internal-DataSourceAgent"
	r.Classes = ""
	r.Aliases = "alias-512a309140d0ad99eb1c95c8dc0d02f9-Agent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6553b6c73b_internal_DataSourceAgent) new() any {
    return &p6553b6c73.DataSourceAgent{}
}

func (inst* p6553b6c73b_internal_DataSourceAgent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6553b6c73.DataSourceAgent)
	nop(ie, com)

	
    com.DSM = inst.getDSM(ie)
    com.SourceName = inst.getSourceName(ie)


    return nil
}


func (inst*p6553b6c73b_internal_DataSourceAgent) getDSM(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}


func (inst*p6553b6c73b_internal_DataSourceAgent) getSourceName(ie application.InjectionExt)string{
    return ie.GetString("${libgorm.agent.source}")
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
    return ie.GetBool("${libgorm.auto-migrate.enabled}")
}



// type p6553b6c73.DefaultDatasource in package:github.com/starter-go/libgorm/internal
//
// id:com-6553b6c73b54fb77-internal-DefaultDatasource
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-DataSource
// alias:
// scope:singleton
//
type p6553b6c73b_internal_DefaultDatasource struct {
}

func (inst* p6553b6c73b_internal_DefaultDatasource) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6553b6c73b54fb77-internal-DefaultDatasource"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-DataSource"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6553b6c73b_internal_DefaultDatasource) new() any {
    return &p6553b6c73.DefaultDatasource{}
}

func (inst* p6553b6c73b_internal_DefaultDatasource) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6553b6c73.DefaultDatasource)
	nop(ie, com)

	
    com.Drivers = inst.getDrivers(ie)
    com.Name = inst.getName(ie)
    com.Driver = inst.getDriver(ie)
    com.Host = inst.getHost(ie)
    com.Port = inst.getPort(ie)
    com.User = inst.getUser(ie)
    com.Password = inst.getPassword(ie)
    com.Database = inst.getDatabase(ie)
    com.Enabled = inst.getEnabled(ie)
    com.Groups = inst.getGroups(ie)


    return nil
}


func (inst*p6553b6c73b_internal_DefaultDatasource) getDrivers(ie application.InjectionExt)p512a30914.DriverManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DriverManager").(p512a30914.DriverManager)
}


func (inst*p6553b6c73b_internal_DefaultDatasource) getName(ie application.InjectionExt)string{
    return ie.GetString("${datasource.default.name}")
}


func (inst*p6553b6c73b_internal_DefaultDatasource) getDriver(ie application.InjectionExt)string{
    return ie.GetString("${datasource.default.driver}")
}


func (inst*p6553b6c73b_internal_DefaultDatasource) getHost(ie application.InjectionExt)string{
    return ie.GetString("${datasource.default.host}")
}


func (inst*p6553b6c73b_internal_DefaultDatasource) getPort(ie application.InjectionExt)int{
    return ie.GetInt("${datasource.default.port}")
}


func (inst*p6553b6c73b_internal_DefaultDatasource) getUser(ie application.InjectionExt)string{
    return ie.GetString("${datasource.default.username}")
}


func (inst*p6553b6c73b_internal_DefaultDatasource) getPassword(ie application.InjectionExt)string{
    return ie.GetString("${datasource.default.password}")
}


func (inst*p6553b6c73b_internal_DefaultDatasource) getDatabase(ie application.InjectionExt)string{
    return ie.GetString("${datasource.default.database}")
}


func (inst*p6553b6c73b_internal_DefaultDatasource) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datasource.default.enabled}")
}


func (inst*p6553b6c73b_internal_DefaultDatasource) getGroups(ie application.InjectionExt)string{
    return ie.GetString("${datasource.default.groups}")
}



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


func (inst*p6553b6c73b_internal_DefaultDatasourceManager) getSources(ie application.InjectionExt)[]p512a30914.DataSource{
    dst := make([]p512a30914.DataSource, 0)
    src := ie.ListComponents(".class-512a309140d0ad99eb1c95c8dc0d02f9-DataSource")
    for _, item1 := range src {
        item2 := item1.(p512a30914.DataSource)
        dst = append(dst, item2)
    }
    return dst
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

	
    com.Groups = inst.getGroups(ie)


    return nil
}


func (inst*p6553b6c73b_internal_GroupManagerImpl) getGroups(ie application.InjectionExt)[]p512a30914.TableGroupRegistry{
    dst := make([]p512a30914.TableGroupRegistry, 0)
    src := ie.ListComponents(".class-512a309140d0ad99eb1c95c8dc0d02f9-TableGroupRegistry")
    for _, item1 := range src {
        item2 := item1.(p512a30914.TableGroupRegistry)
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


