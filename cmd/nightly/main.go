package main

import (
	"log"
	"os"
	"time"
	"unicode"

	"github.com/Awesome-Sauces/Note/terminal"
	"github.com/gdamore/tcell/v2"
)

var CURRENT_NOTE_MODE = 0
var CURSOR_POSITION_X = 0
var CURSOR_POSITION_Y = 0

const (
	version                 = "NOTE v0.0.1"
	version_position_center = 7 // for a full sized term
)

var TextBuffer = make([]rune, 0)

func IsValidRune(r rune) bool {
	// Check if the rune is a digit
	if unicode.IsDigit(r) {
		return true
	}

	// Check if the rune is a letter
	if unicode.IsLetter(r) {
		return true
	}

	// Check if the rune is a special ASCII character
	specialChars := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	for _, special := range specialChars {
		if r == special {
			return true
		}
	}

	// Otherwise, return false
	return false
}

func main() {

	args := os.Args[1:]
	var filename string

	if len(args) > 0 {
		filename = args[0]
	}

	DRAW_TO_SCREEN := func(view *terminal.TerminalEnvironment) error {
		view.Screen().Clear()
		for _, content := range view.GetElements() {
			terminal.FormattedPuts(view.Screen(), content.Style, content.X, content.Y, content.Text)
		}

		err := Commands_Draw(view, filename)

		if err != nil {
			return err
		}

		view.Screen().Show()

		return nil
	}

	//REGISTERED_EVENT_HANDLERS[0] = func(term terminal.TerminalEnvironment, event tcell.Event) error { return nil }

	log_file_name := time.Now().String() + ".log"

	log_file, err := os.OpenFile(log_file_name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer log_file.Close()

	log.SetOutput(log_file)

	// Initialize the terminal environment
	termEnv, err := terminal.TerminalEnvironment_new()
	if err != nil {
		log.Fatalf("Failed to initialize terminal: %v", err)
	}
	//defer termEnv.Fini()

	screen := termEnv.Screen()

	// Use the terminal’s default foreground and background colors
	defStyle := tcell.StyleDefault.
		Foreground(tcell.ColorDefault).
		Background(tcell.ColorDefault)
	screen.SetStyle(defStyle)

	_, screenHeight := screen.Size()
	// Add the message as an element
	command_line_content := &terminal.ElementContent{
		Style: tcell.StyleDefault.Foreground(tcell.ColorWhite),
		X:     0,
		Y:     max(screenHeight-1, 0),
		Text:  "",
	}

	version_centerRow, version_centerColumn := terminal.CenterPosition(screen, version)
	version_centerRow -= version_position_center

	version_tag := &terminal.ElementContent{
		Style: tcell.StyleDefault.Foreground(tcell.ColorWhite),
		X:     version_centerColumn,
		Y:     version_centerRow,
		Text:  version,
	}

	if err := termEnv.TerminalEnvironment_set_ElementContent(0, version_tag); err != nil {
		log.Fatalf("Failed to add element content: %v", err)
	}

	if err := termEnv.TerminalEnvironment_set_ElementContent(1, command_line_content); err != nil {
		log.Fatalf("Failed to add element content: %v", err)
	}

	if err := DRAW_TO_SCREEN(termEnv); err != nil {
		log.Fatalf("Failed to render screen: %v", err)
	}

	screen.ShowCursor(0, 0)
	screen.SetCursorStyle(tcell.CursorStyleSteadyBlock)
	screen.Show()

	// Java Centric
	/*
		UserKeyEvent() -> *terminal.TerminalEnvironment, *tcell.EventKey

	*/

MAIN_LOOP:
	if err := DRAW_TO_SCREEN(termEnv); err != nil {
		log.Fatalf("Failed to render screen: %v", err)
	}

	ev := screen.PollEvent()

	switch event := ev.(type) {
	case *tcell.EventKey:
		if IsValidRune(event.Rune()) {
			TextBuffer = append(TextBuffer, event.Rune())

		}

		commands_UserKeyd(termEnv, event)
	case *tcell.EventResize: // on application launch this is called
		if err := DRAW_TO_SCREEN(termEnv); err != nil {
			log.Fatalf("Failed to render screen: %v", err)
		}
		goto MAIN_LOOP
	}

	goto MAIN_LOOP

}

func EXIT_CURRENT_NOTE_MODE() {
	// safely exit note
	exit_initial_time := time.Now()

	// exit code
	time.Sleep(10 * time.Millisecond)

	exit_final_time := time.Now()

	log.Printf("safely exited mode in %vs", exit_final_time.Sub(exit_initial_time).Seconds())
}

func SAFE_EXIT(environment *terminal.TerminalEnvironment) {
	// safely exit note
	exit_initial_time := time.Now()

	// exit code
	environment.Fini()
	time.Sleep(32 * time.Millisecond)

	exit_final_time := time.Now()

	log.Printf("safely exited note in %vs", exit_final_time.Sub(exit_initial_time).Seconds())

	os.Exit(0)
}
