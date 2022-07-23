package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
//	"time"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	
)

// Cursor color
var color string = "[:#00aeff]"

// NoteFile, utilized to simply
// TextView and other management
type NoteFile struct {
	textView *tview.TextView
	position int
	text     string
	filename string
	order    int
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
	d1 := []byte(v.text)

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

// Update position by 1 or by -1
func (v *NoteFile) updatePos(e bool) {
	if e {
		v.position++
	} else if v.position > 0 {
		v.position--
	}
}

// Draw the cursor to the screen
func (v *NoteFile) drawCursor() {

	var vFile string

	var zFile string

	if v.position > len(v.text) {
		v.position = len(v.text)
	}

	lineColor := "[" + pColorTheme.lnColor + ":" + pColorTheme.lnbgColor + ":" + pColorTheme.lnstyleColor + "]"

	vFile = tview.Escape(v.text) + color + " [:-]"

	if v.position >= 0 {
		vFile = tview.Escape(v.text[0:Math.max(v.position-1, 0)]) + color + tview.Escape(v.text[Math.max(v.position-1, 0):v.position]) + "[:-]" + tview.Escape(v.text[v.position:])
	}

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

	v.textView.Clear()

	v.textView.Write([]byte(zFile))
}

// Handle arrow key movement, input
// Deletion, etc
func (v *NoteFile) NewInputManager() {
	v.textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// If the app isn't focus'd on this view then ignore input
		if !v.textView.HasFocus() {
			return event
		}
		// Skip one word left on left arrow + shift
		if event.Key() == tcell.KeyLeft && event.Modifiers() == tcell.ModShift {
			v.setPos(len(regexp.MustCompile(`\S+\s*$`).ReplaceAllString(v.text[:v.position], "")))
			v.drawCursor()
			return event
			// Skip one word right on right arrow + shift
		} else if event.Key() == tcell.KeyRight && event.Modifiers() == tcell.ModShift {
			v.setPos(len(v.text) - len(regexp.MustCompile(`^\s*\S+\s*`).ReplaceAllString(v.text[v.position:], "")))
			v.drawCursor()
			return event
			// Delete one word right
		} else if event.Key() == tcell.KeyCtrlW {
			lastWord := regexp.MustCompile(`\S+\s*$`)
			newText := lastWord.ReplaceAllString(v.text[:v.position], "") + v.text[v.position:]
			v.position -= len(v.text) - len(newText)

			if v.position < 0 {
				v.position = 0
			}

			v.text = newText
			v.drawCursor()
			return event
			// Left Arrow key
		} else if event.Key() == tcell.KeyLeft {
			v.updatePos(false)
			v.drawCursor()
			return event
			// Right Arrow key
		} else if event.Key() == tcell.KeyRight {
			v.updatePos(true)
			v.drawCursor()
			return event
			// Up Arrow key
		} else if event.Key() == tcell.KeyUp {
			for i := v.position; i > 0; i-- {
				if v.text[i-1:i] == "\n" || v.text[i-1:i] == "\r\n" || v.text[i-1:i] == "\r" {
					v.setPos(i - 1)
					v.drawCursor()
					return event
				}
			}
			return event
			// Down Arrow key
		} else if event.Key() == tcell.KeyDown {
			for i := v.position; i >= 0; i++ {
				if i+1 <= len(v.text) {
					if v.text[i:i+1] == "\n" || v.text[i:i+1] == "\r\n" || v.text[i:i+1] == "\r" {
						v.setPos(i + 1)
						v.drawCursor()
						return event
					}
				} else {
					return event
				}
			}
			return event
			// "Newline"/Enter key
		} else if event.Key() == tcell.KeyEnter {
			v.text = v.text[0:v.position] + "\n" + v.text[v.position:]
			v.updatePos(true)
			v.drawCursor()
			return event
			// Deleting characters at cursor position
		} else if event.Key() == tcell.KeyDEL {

			if len(v.text) > 0 && v.position-1 != -1 {
				v.text = v.text[0:v.position-1] + v.text[v.position:]
				v.updatePos(false)
				v.drawCursor()

				return event
			}

			return event
			// On escape key save and exit
			// To-do:
			// Make it only close the file that ESC was pressed on
		} else if event.Key() == tcell.KeyESC {
			v.saveFile()
			app.Stop()
			// Handling normal
			// Characters
		} else {

			if v.position == len(v.text) {
				v.text += string(event.Rune())
			} else {
				v.text = v.text[0:v.position] + string(event.Rune()) + v.text[v.position:]
			}

			v.updatePos(true)
			v.drawCursor()

			return event

		}
		return event
	})

	v.textView.SetMouseCapture(func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {

		if action == tview.MouseLeftClick {
			row, column := event.Position()

			v.textView.ScrollTo(event.Position())

			v.textView.SetText(strconv.Itoa(row) + ":" + strconv.Itoa(column))
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

	v.position = 0

	v.textView.SetText(v.getStringFormat())
}

// Get the file with color format excluding cursor
func (v NoteFile) getStringFormat() string {

	tempText := tview.Escape(v.text)

	stringReturn := ""

	lineColor := "[" + pColorTheme.lnColor + ":" + pColorTheme.lnbgColor + ":" + pColorTheme.lnstyleColor + "]"

	for index, element := range strings.Split(tempText, "\n") {
		space := ""

		mDigits := iterativeDigitsCount(len(strings.Split(tempText, "\n")))
		sDigits := iterativeDigitsCount(index+1)

		for i := 1; i < (mDigits-sDigits)+1; i++ {
		    space += " "
		}

		element = tview.Escape(element)

		for _, et := range pColorTheme.keywords{
			strings.ReplaceAll(element, et.name, et.color + tview.Escape(et.name))
		}
		
		stringReturn += lineColor + space + strconv.Itoa(index+1) + " [-:-:-] " + element + "\n"
	}

	return stringReturn
}

// Setting postion
// to specific location
func (v *NoteFile) setPos(e int) {
	if e > len(v.text) {
		return
	}
	v.position = e
}

func iterativeDigitsCount(number int) int {
    count := 0
    for number != 0 {
        number /= 10
        count += 1
    }
    return count

}

