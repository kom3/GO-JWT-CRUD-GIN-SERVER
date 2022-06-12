package CommonUtils

import (
	"fmt"
	"gomodules/APIS/CRUDAPIS"
)

func ModuleInitializer() {
	fmt.Printf("\n----Module Initialization Starts----\n\n\n")

	// New module inits can be added below(calling module wise init functions)

	CRUDAPIS.Init() //initializes mongodb connection objet

	fmt.Printf("\n----Module Initialization Ends----\n\n\n")
}
