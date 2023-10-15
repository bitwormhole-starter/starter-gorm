package main

import (
	"os"

	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(libgorm.Module())
	i.WithPanic(true).Run()
}
