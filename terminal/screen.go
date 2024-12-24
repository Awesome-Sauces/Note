package terminal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

/*
DO NOT TOUCH
*/

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

func CenterColumnPosition(s tcell.Screen, text string) int {
	screenWidth, _ := s.Size()
	textWidth := runewidth.StringWidth(text)
	col := (screenWidth - textWidth) / 2

	return col
}

func FormattedPuts(s tcell.Screen, style tcell.Style, x, y int, str string) {
	i := 0
	var deferred []rune
	dwidth := 0
	zwj := false
	for _, r := range str {
		if r == '\u200d' { // Zero width char
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
