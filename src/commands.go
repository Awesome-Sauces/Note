package main

import (
	"fmt"
	"os"
)

func CheckArgs() {

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		fmt.Printf("Note v1.1.1\n\n")
		fmt.Printf("Usage: note [-arg]\n")
		fmt.Printf("    note [--help]\n")
		fmt.Printf("    note [-f] [filename]\n")
		fmt.Printf("    note [-txt] [color]\n")
		fmt.Printf("    note [-bg] [color]\n")
		fmt.Printf("Color list:\n")
		fmt.Printf("    [white]\n    [red]\n    [black]\n    [blue]\n    [green]\n    [purple]\n    [yellow]\n")
		os.Exit(1)
	} else if len(os.Args) > 2 && os.Args[1] == "-txt" {
		textColorChange(os.Args[2])
		os.Exit(3)
	} else if len(os.Args) > 2 && os.Args[1] == "-bg" {
		backGroundChange(os.Args[2])
		os.Exit(3)
	} else if len(os.Args) >= 2 && os.Args[1] == "-f" {
		return
	} else {
		fmt.Println("For help do: note --help")
		os.Exit(0)
	}

}
