package gen4demo1

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportConfigForDemo1 ...
func ExportConfigForDemo1(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
