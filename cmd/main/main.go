package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

// DO NOT ABSTRACT AWAY THE tcell.Screen instance that is weird.

func main() {

	// Init screen Pretty self explanatory
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	if err := screen.Init(); err != nil {
		log.Fatal(err)
	}

	// By default the screen color is set to terminal default color
	// otherwise: screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorDefault).Background(tcell.ColorDefault))

	// Shutdown function
	note_program_shutdown := make(chan bool)

	go func() {
		if <-note_program_shutdown {
			screen.Fini()
			os.Exit(0)
		}
	}()

	goto MAIN_LOOP

MAIN_LOOP:

	// SHUTDOWN Channel, if message received, shutdown

	raw_event := screen.PollEvent()

	//go func() { note_program_shutdown <- false }()

	switch specified_event := raw_event.(type) {
	case *tcell.EventKey:
		if specified_event.Key() == tcell.KeyESC {
			note_program_shutdown <- true
		}
	}

	goto MAIN_LOOP

}
