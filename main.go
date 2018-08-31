package main

import (
	"fmt"

	"github.com/mxrk/cli_task/cmds"
	"github.com/mxrk/cli_task/db"
)

func main() {
	//	fmt.Println("Starting init ...")
	err := db.Init()
	if err != nil {
		fmt.Println("Error @init")
	}
	//fmt.Println("Starting execution ...")
	cmds.Execute()

}
