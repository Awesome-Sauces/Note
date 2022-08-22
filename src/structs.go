package main

import (
	"os"
	//	"regexp"
	"strconv"
	"strings"

	//	"time"
_	 "fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Cursor color
var color string = "[:#00aeff]"

// NoteFile, utilized to simply
// TextView and other management
type NoteFile struct {
	textView *tview.TextView
	filename string
	order    int
	editor   TextEditor
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
	d1 := []byte(v.editor.saveFile())

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
	v.textView.Clear()

	v.textView.SetText(v.editor.finalize())

	//v.textView.Write([]byte(v.editor.finalize()))
}

func (v NoteFile) getFormat() string {
	return v.editor.getFormat()
}

// Handle arrow key movement, input
// Deletion, etc
func (v *NoteFile) NewInputManager() {
	v.textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// If the app isn't focus'd on this view then ignore input
		key := event.Key()
		run := event.Rune()
		letter := string(run)

		if !v.textView.HasFocus() {
			return event
		}

		if key == tcell.KeyLeft && event.Modifiers() == tcell.ModCtrl {
			// Skip one word left on left arrow + shift
			v.editor.moveWordLeft()
			v.drawCursor()
			return event
		} else if key == tcell.KeyRight && event.Modifiers() == tcell.ModCtrl {
			// Skip one word right on right arrow + shift
			v.editor.moveWordRight()
			v.drawCursor()
			return event
		} else if key == tcell.KeyCtrlW {
			// Delete one word right
			v.editor.deleteWord()
			v.drawCursor()
			return event
		} else if key == tcell.KeyLeft {
			// Left Arrow key
			v.editor.moveLeft()
			v.drawCursor()
			return event
		} else if key == tcell.KeyRight {
			// Right Arrow key
			v.editor.moveRight()
			v.drawCursor()
			return event
		} else if key == tcell.KeyUp {
			// Up Arrow key
			v.editor.moveUp()
			v.drawCursor()
			return event
		} else if key == tcell.KeyDown {
			// Down Arrow key
			v.editor.moveDown()
			v.drawCursor()
			return event
		} else if key == tcell.KeyEnter {
			// "Newline"/Enter key
			v.editor.newLine()
			v.drawCursor()
			return event
		} else if key == tcell.KeyDEL {
			// Deleting characters at cursor position
			v.editor.deleteChar()
			v.drawCursor()
			return event
		} else if key == tcell.KeyESC {
			// On escape key save and exit
			// To-do:
			// Make it only close the file that ESC was pressed on
			app.Stop()

			v.saveFile()

			return event
		} else {
			// Handling normal
			// Characters
			if key == tcell.KeyRune || key == tcell.KeyTab{
				v.editor.addChar(string(run))
			}

			v.drawCursor()

			if letter == "h" ||
			letter == "l" ||
			letter == "j" ||
			letter == "k" ||
			letter == "g" ||
			letter == "G" {
				return nil
			}


			return event

		}
		return event
	})

	v.textView.SetMouseCapture(func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {

		if action == tview.MouseLeftClick {
			row, column := event.Position()

			v.textView.ScrollTo(row, column)

			v.editor.setLocation(position{x: row, y: column}, true)

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
	v.textView.SetTitle("[" + v.filename + "]	 Click esc to quit!")
	v.textView.SetBorderColor(tcell.ColorBlack)
	v.textView.SetBorderAttributes(tcell.AttrDim)

	v.textView.SetText(v.editor.getFormat())
}

func iterativeDigitsCount(number int) int {

	str := strconv.Itoa(number)

	return len(str)
}
