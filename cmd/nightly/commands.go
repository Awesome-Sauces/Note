package main

import (
	"fmt"
	"log"

	"github.com/Awesome-Sauces/Note/terminal"
	"github.com/gdamore/tcell/v2"
)

func Commands_Draw(screen *terminal.TerminalEnvironment, filename string) error {

	tcell_screen := screen.Screen()

	width, height := screen.Screen().Size()

	height = max(height-2, 0)

	for i := 0; i < width; i++ {
		tcell_screen.SetCell(i, height, tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack))
	}

	if len(filename) == 0 || filename == "" || filename == " " {
		terminal.FormattedPuts(screen.Screen(), tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack).Bold(true), 0, height, "[No Name]")
	} else {
		terminal.FormattedPuts(screen.Screen(), tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack).Bold(true), 0, height, filename)
	}

	terminal.FormattedPuts(screen.Screen(), tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack).Bold(true), width-10, height, fmt.Sprintf("%d,%d", CURSOR_POSITION_Y, CURSOR_POSITION_X))

	return nil
}

// DONUT call from main
func commands_UserKeyd(term *terminal.TerminalEnvironment, event *tcell.EventKey) error {
	log.Println(TextBuffer)
	if event.Key() == tcell.KeyESC {
		SAFE_EXIT(term)
	}
	return nil
}
