package main

import (
	"os"
//	"regexp"
	"strconv"
	"strings"
//	"time"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
_	"fmt"
)

// Cursor color
var color string = "[:#00aeff]"

// NoteFile, utilized to simply
// TextView and other management
type NoteFile struct {
	textView *tview.TextView
	filename string
	order    int
	editor TextEditor
}

type ColorTheme struct {
	name         string
	lnColor      string
	lnbgColor    string
	bgColor      string
	lnstyleColor string
	keywords     map[string]cKeyWord
}

type cKeyWord struct {
	extension string
	name      string
	color     string
}

func (ct *ColorTheme) setStyleColor(color string) {
	ct.lnstyleColor = color
}

func (ct *ColorTheme) setlnColor(color string) {
	ct.lnColor = color
}

func (ct *ColorTheme) setlnbgColor(color string) {
	ct.lnbgColor = color
}

func (ct *ColorTheme) changeBGcolor(color string) {
	ct.bgColor = color
}

func (ct *ColorTheme) NewKeyword(e cKeyWord) {
	ct.keywords[e.name] = e
}

func (ct *ColorTheme) changeName(name string) {
	ct.name = name
}

// Save file, if it has UNOWN_FILE flag
// then create the file
func (v *NoteFile) saveFile() {
	d1 := []byte(v.editor.getText())

	if strings.Contains(v.filename, "(**|**)") {
		myfile, e := os.Create(strings.ReplaceAll(v.filename, "(**|**)", ""))
		check(e)
		myfile.Close()
		err := os.WriteFile(v.filename, d1, 0644)
		check(err)
	}

	err := os.WriteFile(v.filename, d1, 0644)
	check(err)
}

// Draw the cursor to the screen
func (v *NoteFile) drawCursor() {
/*
	var vFile string

	var zFile string

//	if v.position > len(v.text) {
	//	v.position = len(v.text)
//	}

	lineColor := "[" + pColorTheme.lnColor + ":" + pColorTheme.lnbgColor + ":" + pColorTheme.lnstyleColor + "]"

	for _, et := range pColorTheme.keywords{
		vFile = strings.ReplaceAll(vFile, et.name, "[" + et.color + "]" + et.name + "[#ffffff]")
	}


	for index, element := range strings.Split(vFile, "\n") {
		space := ""

		mDigits := iterativeDigitsCount(len(strings.Split(vFile, "\n")))
		sDigits := iterativeDigitsCount(index+1)

		for i := 1; i < (mDigits-sDigits)+1; i++ {
		    space += " "
		}
		
		zFile += lineColor + space + strconv.Itoa(index+1) + " [-:-:-] " + element + "\n"
	}
*/

//	vFile = tview.Escape(v.text) + color + " [:-]"

//	if v.position >= 0 {
//		vFile = tview.Escape(v.text[0:Math.max(v.position-1, 0)]) + color + tview.Escape(v.text[Math.max(v.position-1, 0):v.position]) + "[:-]" + tview.Escape(v.text[v.position:])
//	}


	v.textView.Clear()

	v.textView.SetText(v.editor.finalize())

	//v.textView.Write([]byte(v.editor.finalize()))
}

func (v NoteFile) getFormat() string{
	return v.editor.getFormat()
}

// Handle arrow key movement, input
// Deletion, etc
func (v *NoteFile) NewInputManager() {
	v.textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// If the app isn't focus'd on this view then ignore input
		if !v.textView.HasFocus() {return event}
		
		if event.Key() == tcell.KeyLeft && event.Modifiers() == tcell.ModCtrl {
			// Skip one word left on left arrow + shift
			v.editor.moveWordLeft()
			v.drawCursor()
			return event
		} else if event.Key() == tcell.KeyRight && event.Modifiers() == tcell.ModCtrl {
			// Skip one word right on right arrow + shift
			v.editor.moveWordRight()
			v.drawCursor()
			return event
		} else if event.Key() == tcell.KeyCtrlW {
			// Delete one word right
			v.editor.deleteWord()
			v.drawCursor()
			return event
		} else if event.Key() == tcell.KeyLeft {
			// Left Arrow key
			v.editor.moveLeft()
			v.drawCursor()
			return event
		} else if event.Key() == tcell.KeyRight {
			// Right Arrow key
			v.editor.moveRight()
			v.drawCursor()
			return event
		} else if event.Key() == tcell.KeyUp {
			// Up Arrow key
			v.editor.moveUp()
			v.drawCursor()
			return event
		} else if event.Key() == tcell.KeyDown {
			// Down Arrow key
			v.editor.moveDown()
			v.drawCursor()
			return event
		} else if event.Key() == tcell.KeyEnter {
			// "Newline"/Enter key
			v.editor.newLine()
			v.drawCursor()
			return event
		} else if event.Key() == tcell.KeyDEL {
			// Deleting characters at cursor position
			v.editor.deleteChar()
			v.drawCursor()
			return event
		} else if event.Key() == tcell.KeyESC {
			// On escape key save and exit
			// To-do:
			// Make it only close the file that ESC was pressed on
			v.saveFile()
			app.Stop()
		} else {
			// Handling normal
			// Characters
			v.editor.addChar(string(event.Rune()))

			v.drawCursor()

			return event

		}
		return event
	})

	v.textView.SetMouseCapture(func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {

		if action == tview.MouseLeftClick {
			row, column := event.Position()

			v.textView.ScrollTo(event.Position())

			v.editor.setLocation(position{x: column, y: row}, true)

			v.drawCursor()
			
		}

		return action, event
	})
}

// Re-instantize the textview
func (v NoteFile) autoCreateTextView() {
	v.startTheme()
	v.NewInputManager()
}

// Load color scheme and other data
func (v *NoteFile) startTheme() {
	v.textView.SetBackgroundColor(tcell.GetColor(pColorTheme.bgColor))

	v.textView.SetTitleColor(tcell.ColorBlack)
	v.textView.SetTitleAlign(tview.AlignCenter)

	v.textView.SetBorder(true)
	v.textView.SetTitle(TextColorGlobal + "[" + v.filename + "]	 Click esc to quit!")
	v.textView.SetBorderColor(tcell.ColorBlack)
	v.textView.SetBorderAttributes(tcell.AttrDim)


	v.textView.SetText(v.editor.getFormat())
}

func iterativeDigitsCount(number int) int {

	str := strconv.Itoa(number)
//	str = strings.ReplaceAll(str, ".", "")

//	fmt.Println(len(str))

    return len(str)
}

