package cmd

import "fmt"

func showVersion() {
	log := NewLog()
	width := log.getTerminalWidth()

	log.Header()

	if width < 80 {

		log.newLine("			   ")
		log.addNewLineToHealder(fmt.Sprintf("    %sEnv info:            %s", log.PRIMARY_COLOR, log.RESET_COLOR))
		log.addNewLineToHealder("       Target: JVM       ")
		log.addNewLineToHealder("       Build:  debug	   ")
		log.finalizeBottomheader()
		return
	}

	log.addNewLineToHealder(fmt.Sprintf("                         %sEnv info:%s                                      ", log.PRIMARY_COLOR, log.RESET_COLOR))
	log.addNewLineToHealder("                             Target: JVM                                ")
	log.addNewLineToHealder("                             Build:  debug                              ")
	log.finalizeBottomheader()
}
