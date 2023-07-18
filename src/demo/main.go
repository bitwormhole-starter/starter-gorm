package main

import (
	"github.com/starter-go/libgorm/modgorm"
	"github.com/starter-go/starter"
)

func main() {

	m := modgorm.Module()

	i := starter.Init(nil)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
