package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Simple error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Initializing tview.Application and
// tview.Flex
var app = tview.NewApplication()
var flex = tview.NewFlex()

var pColorTheme *ColorTheme = &ColorTheme{name: "", keywords: make(map[string]cKeyWord)}

// NoteFile struct map
// Stores all essential data
var noteFiles = make(map[int]NoteFile)

// File window open in
// order of arguments
var current_file = 1

var ResetColor string
var dir string

// Note version for printing to terminal
var NOTE_VERSION string = "Note v1.3.2"
var NOTE_HELP string = NOTE_VERSION + "\n\n" +
	"Usage: note [-arg]\n" +
	"    note [--help]\n" +
	"    note [filename]\n" +
	"    note [-txt] [color]\n" +
	"    note [-bg] [color]\n" +
	"Experimental:\n" +
	"    note -script [SCRIPT_NAME]\n" +
	"Color list:\n" +
	"    [white]\n    [red]\n    [black]\n    [blue]\n    [green]\n    [purple]\n    [yellow]\n"

func main() {
	// Checks args given to cli, if none are
	// given then give a help message
	CheckArgs()

	// Start up Message
	// Sleep .5 seconds to give it
	// a "loading" effect
	fmt.Println("Running " + NOTE_VERSION + ":")
	time.Sleep(500 * time.Millisecond)

	// Loop through noteFiles map and
	// Initializing all the tview.TextView
	for i, s := range noteFiles {
		s.autoCreateTextView()
		flex.AddItem(noteFiles[i].textView, 0, 1, false)
	}

	// Basic User input setup
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if !noteFiles[current_file].textView.HasFocus() {
			app.SetFocus(noteFiles[current_file].textView)
		}

		// When user presses Right Arrow key + ctrl
		// Switch over to the file on right
		if event.Key() == tcell.KeyRight && event.Modifiers() == tcell.ModCtrl {
			if current_file+1 > len(noteFiles) {
				return event
			}
			noteFiles[current_file].textView.SetText(noteFiles[current_file].getFormat())
			current_file++
			app.SetFocus(noteFiles[current_file].textView)
			noteFiles[current_file].autoCreateTextView()
			return event
			// When user presses Left Arrow key + ctrl
			// Switch over to the file on left
		} else if event.Key() == tcell.KeyLeft && event.Modifiers() == tcell.ModCtrl {
			if current_file-1 < 1 {
				return event
			}
			noteFiles[current_file].textView.SetText(noteFiles[current_file].getFormat())
			current_file--
			app.SetFocus(noteFiles[current_file].textView)
			noteFiles[current_file].autoCreateTextView()
			return event
		}

		return event
	})

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
