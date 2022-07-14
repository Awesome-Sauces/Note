package main

import (
	"fmt"
	"os"
	"github.com/rivo/tview"
)

func CheckArgs() {

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		fmt.Printf("%s\n\n", NOTE_VERSION)
		fmt.Printf("Usage: note [-arg]\n")
		fmt.Printf("    note [--help]\n")
		fmt.Printf("    note [filename]\n")
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
	}else if len(os.Args) == 1{
		fmt.Printf("%s\n\n", NOTE_VERSION)
		fmt.Printf("Usage: note [-arg]\n")
		fmt.Printf("    note [--help]\n")
		fmt.Printf("    note [filename]\n")
		fmt.Printf("    note [-txt] [color]\n")
		fmt.Printf("    note [-bg] [color]\n")
		fmt.Printf("Color list:\n")
		fmt.Printf("    [white]\n    [red]\n    [black]\n    [blue]\n    [green]\n    [purple]\n    [yellow]\n")
		os.Exit(1)
	}else if len(os.Args) >= 1{ 
		for i, s := range os.Args {
			if i != 0{

				var dat, err = os.ReadFile(os.Args[i])

				// If file doesn't exist
				// create a temporary file
				// Handled later
				if err != nil{
					noteFiles[i] = NoteFile{textView:  tview.NewTextView().
							SetScrollable(true).
							SetDynamicColors(true).
							SetRegions(true).
							SetChangedFunc(func() {
								app.Draw()
							}), 
							position: 0,
							text: "",
							filename: s + "(**|**)",
							order: i}
				}

				// Creating instance of struct 
				// if file exists
				noteFiles[i] = NoteFile{textView:  tview.NewTextView().
									SetScrollable(true).
									SetDynamicColors(true).
									SetRegions(true).
									SetChangedFunc(func() {
										app.Draw()
									}), 
									position: 0,
									text: string(dat),
									filename: s,
									order: i}
			}
		}

		
	} else {
		fmt.Println("For help do: note --help")
		os.Exit(0)
	}

}
