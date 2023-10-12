package demo1

import "github.com/starter-go/libgorm"

// Source ... 本组数据源
type Source interface {
	libgorm.Agent
}

// TaDao ...
type TaDao interface {
	Find(id int) (*TableA, error)
}

// TbDao ...
type TbDao interface {
	Find(id int) (*TableB, error)
}

// TcDao ...
type TcDao interface {
	Find(id int) (*TableC, error)
}
