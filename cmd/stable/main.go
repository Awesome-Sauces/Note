package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// CenterPosition calculates the center row and column (y, x) for a given string
// based on the screen's current size.
func CenterPosition(s tcell.Screen, text string) (int, int) {
	screenWidth, screenHeight := s.Size()
	textWidth := runewidth.StringWidth(text)

	// row = vertical center, col = horizontal center
	row := screenHeight / 2
	col := (screenWidth - textWidth) / 2

	return row, col
}

// puts is your original function for drawing a string on the screen
func puts(s tcell.Screen, style tcell.Style, x, y int, str string) {
	i := 0
	var deferred []rune
	dwidth := 0
	zwj := false
	for _, r := range str {
		if r == '\u200d' {
			if len(deferred) == 0 {
				deferred = append(deferred, ' ')
				dwidth = 1
			}
			deferred = append(deferred, r)
			zwj = true
			continue
		}
		if zwj {
			deferred = append(deferred, r)
			zwj = false
			continue
		}
		switch runewidth.RuneWidth(r) {
		case 0:
			if len(deferred) == 0 {
				deferred = append(deferred, ' ')
				dwidth = 1
			}
		case 1:
			if len(deferred) != 0 {
				s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 1
		case 2:
			if len(deferred) != 0 {
				s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 2
		}
		deferred = append(deferred, r)
	}
	if len(deferred) != 0 {
		s.SetContent(x+i, y, deferred[0], deferred[1:], style)
		i += dwidth
	}
}

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

	// Clear the screen
	screen.Clear()

	// Our message
	message := "NOTE v0.0.1"

	// Example style
	style := tcell.StyleDefault.
		Foreground(tcell.ColorWhite)

	row_resizer := 7

	// Compute the center row & col for the message
	centerRow, centerCol := CenterPosition(screen, message)

	centerRow -= row_resizer

	if centerRow < 0 {
		centerRow = 0
	}

	// Draw the message in the center (remember that in `puts`, the order is x, y)
	puts(screen, style, centerCol, centerRow, message)

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
			centerRow, centerCol := CenterPosition(screen, message)

			centerRow -= row_resizer

			if centerRow < 0 {
				centerRow = 0
			}

			puts(screen, style, centerCol, centerRow, message)
			screen.Show()
		}
	}
}
