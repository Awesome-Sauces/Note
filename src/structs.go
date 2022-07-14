package main

import(
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"regexp"
	"strings"
	"os"
	"strconv"
)

// Cursor color
var color string = "[#00aeff:#00aeff:lb]"


// NoteFile, utilized to simply 
// TextView and other management
type NoteFile struct{
	textView *tview.TextView
	position int
	text string
	filename string
	order int
}


// Save file, if it has UNOWN_FILE flag
// then create the file
func (v *NoteFile) saveFile(){
	d1 := []byte(v.text)

	if strings.Contains(v.text, "(**|**)"){
	    myfile, e := os.Create(strings.ReplaceAll(v.text, "(**|**)", ""))
    	check(e)
    	myfile.Close()
		err := os.WriteFile(v.filename, d1, 0644)
		check(err)
	}
	
	err := os.WriteFile(v.filename, d1, 0644)
	check(err)
}

// Update position by 1 or by -1
func (v *NoteFile) updatePos(e bool){if e{v.position++}else if v.position > 0{v.position--}}

// Draw the cursor to the screen
func (v *NoteFile) drawCursor(){
	
	var vFile string

	var zFile string

	if v.position > len(v.text) {v.position = len(v.text)}

	vFile = TextColorGlobal + tview.Escape(v.text) + color + " [-:-:-]" + TextColorGlobal

	if v.position >= 0 {vFile = TextColorGlobal + tview.Escape(v.text[0:v.position]) + color + " [-:-:-]" + TextColorGlobal + tview.Escape(v.text[v.position:])} 

	for index, element := range strings.Split(vFile, "\n") {zFile += "[red]" + strconv.Itoa(index+1) + TextColorGlobal + " " + element + "\n"}

	v.textView.Clear()

	v.textView.Write([]byte(zFile))
}

// Handle arrow key movement, input
// Deletion, etc
func (v *NoteFile) NewInputManager(){
	v.textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// If the app isn't focus'd on this view then ignore input
		if(!v.textView.HasFocus()) {return event}
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
		}else {

			if v.position == len(v.text) {
				v.text += string(event.Rune())
			} else {v.text = v.text[0:v.position] + string(event.Rune()) + v.text[v.position:]}

			v.updatePos(true)
			v.drawCursor()

			return event

		}
		return event
	})
}

// Re-instantize the textview
func (v NoteFile) autoCreateTextView(){
	v.startTheme()
	v.NewInputManager()
}

// Load color scheme and other data
func (v *NoteFile) startTheme(){
	switch BackGroundColorGlobal {
			case "[red]":
				v.textView.SetBackgroundColor(tcell.ColorRed)
			case "[green]":
				v.textView.SetBackgroundColor(tcell.ColorGreen)
			case "[yellow]":
				v.textView.SetBackgroundColor(tcell.ColorYellow)
			case "[blue]":
				v.textView.SetBackgroundColor(tcell.ColorBlue)
			case "[purple]":
				v.textView.SetBackgroundColor(tcell.ColorPurple)
			case "[black]":
				v.textView.SetBackgroundColor(tcell.NewRGBColor(30, 30, 30))
			case "[white]":
				v.textView.SetBackgroundColor(tcell.ColorWhite)
			}
	
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
func (v NoteFile) getStringFormat() string{

	tempText := tview.Escape(v.text)

	stringReturn := ""

	for index, element := range strings.Split(tempText, "\n"){
		stringReturn += "[red]" + strconv.Itoa(index+1) + TextColorGlobal + " " + element + "\n"
	}

	return stringReturn
}


// Setting postion
// to specific location
func (v *NoteFile) setPos(e int) {
	if e > len(v.text) {return}
	v.position = e
}
