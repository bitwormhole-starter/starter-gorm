package gen4demo1

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

    
    inst.register(&p1b6e41313c_demo1_TaDaoImpl{})
    inst.register(&p1b6e41313c_demo1_TableReg{})
    inst.register(&p1b6e41313c_demo1_TbDaoImpl{})
    inst.register(&p1b6e41313c_demo1_TcDaoImpl{})


    return nil
}
