package gen4libgorm

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p273dbed527_mock_Driver{})
    inst.register(&p6553b6c73b_internal_DataSourceAgent{})
    inst.register(&p6553b6c73b_internal_DefaultDatasource{})
    inst.register(&p6553b6c73b_internal_DefaultDatasourceManager{})
    inst.register(&p6553b6c73b_internal_DefaultDriverManager{})
    inst.register(&p6553b6c73b_internal_DefaultTableManager{})
    inst.register(&p6553b6c73b_internal_DefaultTableNamer{})


    return nil
}
