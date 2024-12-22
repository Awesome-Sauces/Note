package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	// Create a new screen
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Error creating screen: %v", err)
	}

	// Initialize the screen (enters alt-screen mode on most terminals)
	if err := screen.Init(); err != nil {
		log.Fatalf("Error initializing screen: %v", err)
	}
	defer screen.Fini()

	// Use the terminal’s default foreground and background colors
	defStyle := tcell.StyleDefault.
		Foreground(tcell.ColorDefault).
		Background(tcell.ColorDefault)
	screen.SetStyle(defStyle)

	// Clear the screen so everything is rendered in the default style
	screen.Clear()

	// Some text to display
	message := "Using terminal's default colors!"
	row, col := 5, 2

	// Draw the message with default style
	for i, r := range message {
		screen.SetContent(col+i, row, r, nil, defStyle)
	}

	// Flush changes to the screen
	screen.Show()

	// Event loop: wait for ESC key to exit
	for {
		ev := screen.PollEvent()
		switch event := ev.(type) {
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape {
				return
			}
		case *tcell.EventResize:
			// Clear & redraw on resize
			screen.Clear()
			for i, r := range message {
				screen.SetContent(col+i, row, r, nil, defStyle)
			}
			screen.Show()
		}
	}
}
