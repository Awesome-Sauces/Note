package pages

import "github.com/gdamore/tcell/v2"

type Page interface {
	Draw(tcell.Screen)
	Event_handler(tcell.Event, chan bool)
}

// on even call
// Page.Event_handler(tcell.Event)
// at the end of the main_loop
// call Page.Draw(screen tcell.Screen)
