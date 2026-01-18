package cmd

import (
	"fmt"
)

func showHelp() {
	log := NewLog()

	log.Header()
	log.finalizeBottomheader()
	fmt.Println()

	width := log.getTerminalWidth()

	if width < 80 {
		fmt.Printf("%sCommands%s\n", log.PRIMARY_COLOR, log.RESET_COLOR)
		fmt.Printf("  %snew%s   <project>     Create a new project\n", log.PRIMARY_COLOR, log.RESET_COLOR)
		fmt.Printf("  %supdate%s	      Update loom\n", log.PRIMARY_COLOR, log.RESET_COLOR)
		fmt.Printf("  %suninstall%s	      Uninstall loom", log.PRIMARY_COLOR, log.RESET_COLOR)

		fmt.Println()

		fmt.Printf("%sOptions%s\n", log.PRIMARY_COLOR, log.RESET_COLOR)
		fmt.Printf("  %s-V%s, %s--version%s   Versions\n", log.PRIMARY_COLOR, log.RESET_COLOR, log.PRIMARY_COLOR, log.RESET_COLOR)
		fmt.Printf("  %s-h%s, %s--help%s      This help\n", log.PRIMARY_COLOR, log.RESET_COLOR, log.PRIMARY_COLOR, log.RESET_COLOR)

		fmt.Println()

		fmt.Printf("%sVersions%s\n", log.PRIMARY_COLOR, log.RESET_COLOR)
		fmt.Printf("  loom:       %s-dev\n", log.LoomVersion)
		fmt.Printf("  Klar Core: %s-dev\n", log.KlarVersion)

		log.Line()

		return
	}

	fmt.Printf("%sCommands%s\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("  %snew%s   <project>     Create a new Klar project\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("  %supdate%s	      Update loom (in the future Klar)\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("  %suninstall%s	      Uninstall loom (in the future Klar and dependences)\n", log.PRIMARY_COLOR, log.RESET_COLOR)

	fmt.Println()

	fmt.Printf("%sOptions%s\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("  %s-V%s, %s--version%s   Show versions\n", log.PRIMARY_COLOR, log.RESET_COLOR, log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("  %s-h%s, %s--help%s      Show this help\n", log.PRIMARY_COLOR, log.RESET_COLOR, log.PRIMARY_COLOR, log.RESET_COLOR)

	fmt.Println()

	fmt.Printf("%sVersions%s\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("  loom:       %s-dev\n", log.LoomVersion)

	log.Line()

}
