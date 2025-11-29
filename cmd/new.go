package cmd

import (
	"fmt"
	"os"
)


func newCommand(){
	if len(os.Args) < 3{
		fmt.Println("Usage: loom new <project_name>")
		return
	}

	log := NewLog()

	log.Header()
	fmt.Println()

	name := os.Args[2]
	fmt.Println("loom new ", name)
	log.Line()

	fmt.Println("◉ Creating project structure...")
	
	err := os.Mkdir(name, 0755)
	if err != nil {
		fmt.Println("\n❌ Error creating project:", err)
		return
	}
	
	fmt.Println("◉ Writing default manifest...")
	fmt.Println("◉ Initializing Klang module...")
	fmt.Println("◉ Ready.")


	fmt.Println("\n✔ Project ", name, " created.")
}