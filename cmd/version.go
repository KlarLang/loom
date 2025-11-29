package cmd

import (
	"fmt"
	"os"
)


func showVersion(){
	log := NewLog()

	if len(os.Args) - 1 >= 2 {
		if os.Args[2] == "--short" || os.Args[2] == "-st"{
			fmt.Println("loom --version")
			log.Line()

			fmt.Printf("loom %s-dev\n", log.LoomVersion)
			fmt.Printf("Klang core %s-dev\n", log.KlangVersion)
			return
		} 
		
		fmt.Println("Unexpected command:'", os.Args[3], "'")
		return
		
	}

	fmt.Println("loom --version")
	log.Line()
	fmt.Printf("loom %s-dev\n", log.LoomVersion)
	fmt.Printf("Klang core %s-dev\n", log.KlangVersion)
	fmt.Println("Target: JVM", )
	fmt.Println("Build: debug")
	log.Line()
}