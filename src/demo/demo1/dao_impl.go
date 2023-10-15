package demo1

import "fmt"

////////////////////////////////////////////////////////////////////////////////

// TaDaoImpl ...
type TaDaoImpl struct {

	//starter:component

	Src MyAgent //starter:inject("#")
}

func (inst *TaDaoImpl) _impl() TaDao {
	return inst
}

// Find ...
func (inst *TaDaoImpl) Find(id int) (*TableA, error) {
	return nil, fmt.Errorf("no impl")
}

////////////////////////////////////////////////////////////////////////////////

// TbDaoImpl ...
type TbDaoImpl struct {
	//starter:component

	Src MyAgent //starter:inject("#")

}

func (inst *TbDaoImpl) _impl() TbDao {
	return inst
}

// Find ...
func (inst *TbDaoImpl) Find(id int) (*TableB, error) {
	return nil, fmt.Errorf("no impl")
}

////////////////////////////////////////////////////////////////////////////////

// TcDaoImpl ...
type TcDaoImpl struct {

	//starter:component

	Src MyAgent //starter:inject("#")

}

func (inst *TcDaoImpl) _impl() TcDao {
	return inst
}

// Find ...
func (inst *TcDaoImpl) Find(id int) (*TableC, error) {
	o := &TableC{}
	db := inst.Src.DB(nil)
	res := db.Find(o, id)
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

////////////////////////////////////////////////////////////////////////////////
