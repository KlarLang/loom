package cmd

import (
	"fmt"
	"os"
	"strings"
)

func lexCommand(){
	if len(os.Args) < 3{
		fmt.Println("Usage: loom lex <.k file>")
		return
	}

	
	log := NewLog()

	log.Header()
	fmt.Println()

	file := os.Args[2]
	if !strings.HasSuffix(file, ".k"){
		indexOfDot := strings.LastIndex(file, ".")

		if indexOfDot == -1 {
			fmt.Println("No suffix/extension found in: ", file)
			return
		}

		sufixo := file[indexOfDot:]

		fmt.Println("Erro: Not a .k file")
		log.Line()
		
		fmt.Println("File: ", file, " is a ", sufixo, " file")
		fmt.Println("did you provide the wrong file?")
		
		return
	}

    data, err := os.ReadFile(file)

	if err != nil {
		fmt.Println("\n❌ Error when opening ", file, ": ", err)
		return
	}

	fmt.Println("loom lex ", file)
	log.Line()

	fmt.Println("Running lexer in (simulation)...")
	fmt.Println("\n✔ ", file, " successfully lexicalized!")

	fmt.Println("Conteúdo de ", file, " (primeiros 50 bytes):")
	log.Line()
	fmt.Println(string(data[:50])) 

	fmt.Println()

	fmt.Println("Tokens:")
	log.Line()
	fmt.Println("Tokens here...")

}