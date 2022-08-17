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
	} else if len(os.Args) > 1 && os.Args[1] == "-test" {

		/*
				editor := TextEditor{rowSpacing: 3}

				var det, er = os.ReadFile(os.Args[2])
				check(er)

				//fmt.Println(string(det))
				editor.loadText(string(det)"print('Hello World')\ndef main:\nprint('Main')")
				//editor.moveRight()


				var dat, err = os.ReadFile("note.rocky")

				check(err)

				var txt string = string(dat)

			//	for i, s := range os.Args {if i > 1{txt+=s + " "}}

				lex := Lexer{text: txt,
							 position: 0,
							 }
				lex.lex()

				lex.eval()




				//editor.deleteChar()
				//editor.deleteChar()

				fmt.Println(editor.finalize())
				editor.newLine()
				fmt.Println(editor.finalize())
		*/
		mapy := make(map[int]string)

		mapy[1] = "Hello"
		mapy[2] = "World"
		fmt.Println(len(mapy))
		delete(mapy, 2)
		fmt.Println(len(mapy))

		os.Exit(1)
	} else if len(os.Args) > 1 && os.Args[1] == "-script" {
		fmt.Printf("%s\n\n", NOTE_VERSION)
		fmt.Printf("Running Rocky Runtime-Enviroment\n")

		var dat, err = os.ReadFile("note.rocky")

		check(err)

		var txt string = string(dat)

		//	for i, s := range os.Args {if i > 1{txt+=s + " "}}

		lex := Lexer{text: txt,
			position: 0,
		}
		lex.lex()

		lex.eval()

		os.Exit(1)
	} else if len(os.Args) == 1 {
		fmt.Printf("%s", NOTE_HELP)
		os.Exit(1)
	} else if len(os.Args) >= 1 {
		var dat, err = os.ReadFile("note.rocky")

		check(err)

		var txt string = string(dat)

		//	for i, s := range os.Args {if i > 1{txt+=s + " "}}

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
