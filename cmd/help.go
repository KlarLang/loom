package cmd

import (
	"fmt"
)

func showHelp() {	
	log := NewLog()

	log.Header()
	fmt.Println()

	fmt.Println("loom --help")
	log.Line()
	fmt.Println()

	fmt.Println("Commands")
	fmt.Println("	new <project_name>		Create a new Klang project")
	fmt.Println("	lex <.k file>	     		Lexicalize a .k file")

	fmt.Println()

	fmt.Println("Options")
	fmt.Println("	-V, --version   		Show versions")
	fmt.Println("	-h, --help       		Show this help log")

	fmt.Println()

	fmt.Println("Versions")
	fmt.Printf("	loom: 				%s-dev\n", log.LoomVersion)
	fmt.Printf("	Klang core: 			%s-dev\n", log.KlangVersion)

}