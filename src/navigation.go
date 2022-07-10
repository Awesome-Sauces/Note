package main

import (
	"strconv"
	"strings"

	"github.com/rivo/tview"
)

func updatePos(e bool) {
	if e {
		position++
	} else {
		if position > 0 {
			position--
		}
	}
}

func setPos(e int) {
	if e > len(file) {
		return
	}

	position = e

}

func refresh(quickRefresh bool) {

	var vFile string

	var zFile string

	if position > len(file) {
		position = len(file)

	}

	var color string = "[#00aeff::lb]"

	if position >= 0 {
		vFile = TextColorGlobal + tview.Escape(file[0:position]) + color + "|[-:-:-]" + TextColorGlobal + tview.Escape(file[position:])
	} else {
		vFile = TextColorGlobal + tview.Escape(file) + color + "|[-:-:-]" + TextColorGlobal
	}

	for index, element := range strings.Split(vFile, "\n") {

		zFile += "[red]" + strconv.Itoa(index+1) + TextColorGlobal + " " + element + "\n"

	}

	text.Clear()

	text.Write([]byte(zFile))

	//fmt.Fprintf(text, "%s", zFile)
}
