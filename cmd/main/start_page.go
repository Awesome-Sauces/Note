package main

import (
	"github.com/Awesome-Sauces/Note/terminal"
	"github.com/gdamore/tcell/v2"
)

type start_page struct {
	blah int64
}

func StartPage_initialize() *start_page {
	return &start_page{blah: 30}
}

func (sp start_page) Draw(screen tcell.Screen) {

	terminal.FormattedPuts(screen, tcell.StyleDefault.Background(tcell.ColorWhite), 5, 5, "Hello")

}

func (sp *start_page) Event_handler(raw_event tcell.Event, shutdown_handler chan bool) {

}
