package main

import (
	"os"

	"github.com/starter-go/libgorm/modgorm"
	"github.com/starter-go/starter"
)

func main() {

	m := modgorm.ModuleForTest()

	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
