package terminal

import (
	"github.com/gdamore/tcell/v2"
)

type TerminalEnvironment struct {
	screen   tcell.Screen
	elements map[int]*ElementContent // map ID to content
}

type ElementContent struct {
	Style tcell.Style
	X     int
	Y     int
	Text  string
}

type TerminalEnvironmentError struct {
	message string
}

func (e *TerminalEnvironmentError) Error() string {
	return e.message
}

func (e *TerminalEnvironment) GetElements() map[int]*ElementContent {
	return e.elements
}

func (e *TerminalEnvironment) GetElementContent(ID int) *ElementContent {
	return e.elements[ID]
}

func TerminalEnvironment_new() (*TerminalEnvironment, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err := screen.Init(); err != nil {
		return nil, err
	}

	return &TerminalEnvironment{
		screen:   screen,
		elements: make(map[int]*ElementContent),
	}, nil
}

func (te *TerminalEnvironment) Screen() tcell.Screen {
	return te.screen
}

func (te *TerminalEnvironment) TerminalEnvironment_set_ElementContent(ID int, content *ElementContent) error {

	te.elements[ID] = content
	return nil
}

func (te *TerminalEnvironment) Fini() {
	te.screen.Fini()
}
