package main

import (
	"fmt"
	"os"
	"github.com/rivo/tview"
	"strings"
)

var (
	stdinfd  = int(os.Stdin.Fd())
	stdoutfd = int(os.Stdout.Fd())
)

type TextEditor struct{
	text string
	rows map[int]string
	rowSpacing int
}

type position struct{
	x int
	y int
}

func (txt *TextEditor) getRow(pos position) string {
	return txt.rows[Math.min(pos.y, len(txt.rows))]
}

func (txt *TextEditor) getChar(pos position) string {
	pos.x += txt.rowSpacing
	return txt.rows[Math.min(pos.y, len(txt.rows))][Math.max(pos.x-1, 0):Math.min(pos.x, len(txt.rows[Math.min(pos.y, len(txt.rows))]))]
}

func (txt *TextEditor) orderText(text string){
	txt.text = text
	txt.rows = make(map[int]string)
	for _, element := range strings.Split(text, "\n") {txt.rows[len(txt.rows)+1] = element}
}

func (txt *TextEditor) initText(text string){
	txt.text = text
	yMap := make(map[int]string)
	for _, element := range strings.Split(text, "\n") {yMap[len(yMap)+1] = element}
	txt.rows = yMap
}

func CheckArgs() {

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		fmt.Printf("%s", NOTE_HELP)
		os.Exit(1)
	} else if len(os.Args) > 2 && os.Args[1] == "-txt" {
		textColorChange(os.Args[2])
		os.Exit(3)
	} else if len(os.Args) > 2 && os.Args[1] == "-bg" {
		backGroundChange(os.Args[2])
		os.Exit(3)
	} else if len(os.Args) > 1 && os.Args[1] == "-test"{
		
		editor := TextEditor{rowSpacing: 3}

		editor.initText("Hello world this is my text" +
						"\n Very amazing i see")
		
		fmt.Println("VALUE: " + editor.getChar(position{y: 2, x: 5}))
		

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
	}else if len(os.Args) == 1{
		fmt.Printf("%s", NOTE_HELP)
		os.Exit(1)
	}else if len(os.Args) >= 1{ 
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
