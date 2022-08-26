package main

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
)

var (
	stdinfd  = int(os.Stdin.Fd())
	stdoutfd = int(os.Stdout.Fd())
)

func CheckArgs() {

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		fmt.Printf("%s", NOTE_HELP)
		os.Exit(1)
	} else if len(os.Args) > 1 && os.Args[1] == "-deploy" {
		fmt.Printf("If permission is denied please\ncheck if note is a locked directory\n\n")
	
		if len(os.Args) <= 2 {
			fmt.Printf("To deploy your color theme to local\n"+
						"note installation please provide the file name\n" + 
						"that contains your theme.")
			os.Exit(1)
		}

		dirname, err := os.UserHomeDir()
		dir = dirname + "/note/themes.rocky"

		check(err)

		var dat, errr = os.ReadFile(os.Args[2])
		check(errr)

		d1 := []byte(string(dat))

		write := os.WriteFile(dir, d1, 0644)
		check(write)

		fmt.Printf("Success!")

		os.Exit(1)
	} else if len(os.Args) > 1 && os.Args[1] == "-script" {
		if len(os.Args) <= 2 {
			fmt.Printf("Please enter script filename\n")
			os.Exit(1)
		}
	
		fmt.Printf("%s\n\n", NOTE_VERSION)
		fmt.Printf("Running Rocky Runtime-Enviroment\n")

		var dat, err = os.ReadFile(os.Args[2])

		check(err)

		var txt string = string(dat)

		//	for i, s := range os.Args {if i > 1{txt+=s + " "}}

		lex := Lexer{text: txt,position: 0,}
		lex.lex()

		lex.eval()

		os.Exit(1)
	} else if len(os.Args) == 1 {
		fmt.Printf("%s", NOTE_HELP)
		os.Exit(1)
	} else if len(os.Args) >= 1 {

		dirname, err := os.UserHomeDir()
		dir = dirname + "/note/themes.rocky"

		check(err)
		var dat, errorr = os.ReadFile(dir)

		check(errorr)

		var txt string = string(dat)

		lex := Lexer{text: txt,
			position: 0,
		}
		lex.lex()

		lex.eval()

		for i, s := range os.Args {
			if i != 0 {

				var dat, err = os.ReadFile(os.Args[i])

				// If file doesn't exist
				// create a temporary file
				// Handled later
				if err != nil {
					noteFiles[i] = NoteFile{textView: tview.NewTextView().
						SetDynamicColors(true).
						SetRegions(true).
						SetChangedFunc(func() {
							app.Draw()
						}),
						filename: s + "(**|**)",
						order:    i}
				}

				// Creating instance of struct
				// if file exists
				edit := TextEditor{}
				edit.loadText(string(dat))
				noteFiles[i] = NoteFile{textView: tview.NewTextView().
					SetDynamicColors(true).
					SetRegions(true).
					SetChangedFunc(func() {
						app.Draw()
					}),
					editor:   edit,
					filename: s,
					order:    i}
			}
		}

	} else {
		fmt.Println("For help do: note --help")
		os.Exit(0)
	}

}
