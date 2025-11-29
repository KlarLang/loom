package cmd

import "fmt"

type Log struct {
	LoomVersion string
	KlangVersion string
}

func (l Log) Line() {
	fmt.Println("────────────────────────────────────────────────────────")
}

func (l Log) Header() {
	fmt.Println("╭────────────────────────────────────────────────────────╮")
	fmt.Println("│  loom — Klang Project Manager                          │")
	fmt.Printf("│  version %s-dev  •  Klang Core %s             │\n", l.LoomVersion, l.KlangVersion)
	fmt.Println("╰────────────────────────────────────────────────────────╯")
}

func NewLog() Log {
	return Log{
		LoomVersion: "v0.1.1",
		KlangVersion: "v0.1.10",
	}
}