package main

import (
	"fmt"
	_ "fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rivo/tview"
)

type TextEditor struct {
	rows       map[int]string
	rowSpacing int
	cursor     position
}

type position struct {
	x int
	y int
}

func (txt *TextEditor) getRow(pos position) string {
	return txt.rows[Math.min(pos.y, len(txt.rows))]
}

func (txt *TextEditor) getChar(pos position) string {
	return txt.rows[Math.min(pos.y, len(txt.rows))][Math.max(pos.x-1, 0):Math.min(pos.x, len(txt.rows[Math.min(pos.y, len(txt.rows))]))]
}

func (txt *TextEditor) showCursor(rows map[int]string) string {
	text := ""
	//vFile = tview.Escape(v.text[0:Math.max(v.position-1, 0)]) + color + tview.Escape(v.text[Math.max(v.position-1, 0):v.position]) + "[:-]" + tview.Escape(v.text[v.position:])
	for iter, i := range rows {
		if iter == txt.cursor.y {

			i = i[0:Math.max(txt.cursor.x-1, 0)] + "[:#00aeff]" + i[Math.max(txt.cursor.x-1, 0):Math.max(txt.cursor.x, 0)] + "[:-]" + i[Math.max(txt.cursor.x, 0):]

			//		i[Math.max(txt.cursor.x-1, 0):Math.min(txt.cursor.x, len(txt.rows[Math.min(txt.cursor.y, len(txt.rows))]))]
		}
		text += i + "\n"
	}

	return text
}

func (txt *TextEditor) getFormat() string {
	final := ""

	for LOOP := 1; LOOP <= len(txt.rows); LOOP++ {
		index := LOOP
		element := tview.Escape(txt.rows[LOOP])

		space := ""
		mDigits := iterativeDigitsCount(len(txt.rows))
		sDigits := iterativeDigitsCount(index)
		NumberColor := "[" + pColorTheme.lnColor + ":" + pColorTheme.lnbgColor + ":" + pColorTheme.lnstyleColor + "]"

		for _, et := range pColorTheme.keywords {

			element = strings.ReplaceAll(element, et.name, "["+et.color+"]"+et.name+"[#ffffff]")

			//			return i
			//ping <- true
		}

		//lineColor := "[" + pColorTheme.lnColor + ":" + pColorTheme.lnbgColor + ":" + pColorTheme.lnstyleColor + "]"

		for i := 1; i < (mDigits-sDigits)+1; i++ {
			space += " "
		}

		lineNumbers := space + strconv.Itoa(index) + " [-:-:-]"

		txt.rowSpacing = len(lineNumbers) - (7 + len(strconv.Itoa(index)))

		final += NumberColor + lineNumbers + element + "\n"

	}

	return final
}

func (txt *TextEditor) getLocation() position {
	return position{x: txt.cursor.x + txt.rowSpacing, y: txt.cursor.y}
}

func (txt *TextEditor) finalize() string {
	final := ""

	for LOOP := 1; LOOP <= len(txt.rows); LOOP++ {
		//	fmt.Println(LOOP)
		index := LOOP
		element := txt.rows[LOOP]

		space := ""
		mDigits := iterativeDigitsCount(len(txt.rows))
		sDigits := iterativeDigitsCount(index)
		NumberColor := "[" + pColorTheme.lnColor + ":" + pColorTheme.lnbgColor + ":" + pColorTheme.lnstyleColor + "]"

		if index == txt.cursor.y {

			//app.Stop()
			//fmt.Println(Math.max(txt.cursor.x-1, 0))
			//fmt.Println(Math.min(txt.cursor.x, len(element)))
			element = tview.Escape(element[0:Math.min(Math.max(txt.cursor.x, 0), len(txt.rows[txt.cursor.y]))]) + "[:#00aeff] " + "[:-]" + tview.Escape(element[Math.min(Math.max(txt.cursor.x, 0), len(txt.rows[txt.cursor.y])):])
			//element = element[0:Math.max(txt.cursor.x-1, 0)] + "[:#00aeff]" + element[Math.max(txt.cursor.x-1, 0):Math.min(txt.cursor.x, 0)] + "[:-]" + element[Math.max(txt.cursor.x, 0):]
			//element = tview.Escape(element[0:Math.max(txt.cursor.x-1, 0)]) + "[:#00aeff]" + tview.Escape(element[Math.max(txt.cursor.x-1, 0):Math.min(txt.cursor.x, len(element))]) + "[:-]" + tview.Escape(element[Math.max(txt.cursor.x, 0):])
		}

		for _, et := range pColorTheme.keywords {

			element = strings.ReplaceAll(element, et.name, "["+et.color+"]"+et.name+"[#ffffff]")

			//			return i
			//ping <- true
		}

		//lineColor := "[" + pColorTheme.lnColor + ":" + pColorTheme.lnbgColor + ":" + pColorTheme.lnstyleColor + "]"

		for i := 1; i < (mDigits-sDigits)+1; i++ {
			space += " "
		}

		lineNumbers := space + strconv.Itoa(index) + " [-:-:-]"

		txt.rowSpacing = len(lineNumbers) - (7 + len(strconv.Itoa(index)))

		final += NumberColor + lineNumbers + element + "\n"

	}

	return final
}

func (txt *TextEditor) setLocation(pos position, spacing bool) {
	if spacing {
		txt.cursor = position{x: pos.x - txt.rowSpacing, y: pos.y}
	} else {
		txt.cursor = position{x: pos.x, y: pos.y}
	}
}

func (txt *TextEditor) newLine() {
	//app.Stop()
	//fmt.Println(txt.rows[txt.cursor.y] + "--------")
	txt.rows[txt.cursor.y] = txt.rows[txt.cursor.y][0:Math.max(txt.cursor.x, 0)] + "\n" + txt.rows[txt.cursor.y][Math.max(txt.cursor.x, 0):]
	temp := txt.getText()
	list := strings.Split(temp, "\n")
	txt.rows = make(map[int]string)

	//	fmt.Println(strings.Split(temp, "\n"))

	//fmt.Println(temp + "-------------")

	for index, element := range list {
		if index+1 != len(list) {
			txt.rows[index+1] = element
			//		fmt.Println(index+1, ":" + element)
		}
	}

	txt.moveDown()

}

func (txt *TextEditor) getText() string {
	text := ""

	fmt.Println(len(txt.rows))

	for LOOP := 1; LOOP <= len(txt.rows); LOOP++ {
		i := txt.rows[LOOP]
		text += i + "\n"
	}

	return text
}

func (txt *TextEditor) loadText(text string) {
	yMap := make(map[int]string)
	for _, element := range strings.Split(text, "\n") {
		yMap[len(yMap)+1] = element
	}
	txt.rows = yMap

	txt.cursor = position{x: 0, y: 0}
}

func (txt *TextEditor) initCursor() {
	txt.cursor = position{x: 0, y: 0}
}

func (txt *TextEditor) deleteWord() {
	lastWord := regexp.MustCompile(`\S+\s*$`)
	newText := lastWord.ReplaceAllString(txt.rows[txt.cursor.y][:txt.cursor.x], "") + txt.rows[txt.cursor.y][txt.cursor.x:]
	txt.cursor.x = Math.max(txt.cursor.x-(len(txt.rows[txt.cursor.y])-len(newText)), 0)

	if txt.cursor.x < 0 {
		txt.cursor.x = 0
	}

	txt.rows[txt.cursor.y] = newText
}

func (txt *TextEditor) deleteChar() {
	//txt.rows[Math.min(len(txt.rows), txt.cursor.y)] = txt.rows[Math.min(len(txt.rows), txt.cursor.y)][0:Math.max(txt.cursor.x-1, 0)] + txt.rows[Math.min(len(txt.rows), txt.cursor.y)][Math.max(txt.cursor.x, 0):]

	if len(txt.rows[txt.cursor.y]) <= 0 && txt.cursor.x <= 0 {
		app.Stop()

		fmt.Println(":", len(txt.rows), ":")
		delete(txt.rows, txt.cursor.y)
		fmt.Println(":", len(txt.rows), ":")
		temp := txt.getText()
		fmt.Println(temp)
		list := strings.Split(temp, "\n")
		fmt.Println("len:", len(list), ":")
		txt.rows = make(map[int]string)
		for index, element := range list {
			txt.rows[index+1] = element
			fmt.Println(index + 1)
		}

		fmt.Println(":", len(txt.rows), ":")

		txt.moveUp()
	} else if txt.cursor.x > 0 {
		txt.rows[txt.cursor.y] = txt.rows[txt.cursor.y][0:Math.min(Math.max(txt.cursor.x-1, 0), len(txt.rows[txt.cursor.y]))] + txt.rows[txt.cursor.y][txt.cursor.x:]
		txt.moveLeft()
	} else if txt.cursor.x <= 0 {
		row := txt.rows[txt.cursor.y]
		//pos := txt.cursor.y

		delete(txt.rows, txt.cursor.y)

		txt.rows[txt.cursor.y-1] += row

		temp := txt.getText()
		list := strings.Split(temp, "\n")
		txt.rows = make(map[int]string)

		//	fmt.Println(strings.Split(temp, "\n"))

		//fmt.Println(temp + "-------------")

		for index, element := range list {
			txt.rows[index+1] = element
			//		fmt.Println(index+1, ":" + element)
		}

		txt.moveUp()
	}
}

func (txt *TextEditor) addChar(e string) {
	txt.rows[txt.cursor.y] = txt.rows[txt.cursor.y][0:txt.cursor.x] + e + txt.rows[txt.cursor.y][txt.cursor.x:]
	txt.cursor.x++
}

func (txt *TextEditor) moveUp() {
	txt.cursor.y = Math.max(0, txt.cursor.y-1)
	if txt.cursor.x > len(txt.rows[txt.cursor.y]) {
		txt.cursor.x = len(txt.rows[txt.cursor.y])
	}
	if txt.cursor.x < len(txt.rows[txt.cursor.y]) {
		txt.cursor.x = len(txt.rows[txt.cursor.y])
	}
}

func (txt *TextEditor) moveDown() {
	txt.cursor.y = Math.min(len(txt.rows), txt.cursor.y+1)
	if txt.cursor.x > len(txt.rows[txt.cursor.y]) {
		txt.cursor.x = len(txt.rows[txt.cursor.y])
	}
	if txt.cursor.x < len(txt.rows[txt.cursor.y]) {
		txt.cursor.x = len(txt.rows[txt.cursor.y])
	}
}

func (txt *TextEditor) enter() {
	txt.rows[len(txt.rows)+1] = txt.rows[txt.cursor.y][txt.cursor.x:]
	txt.rows[txt.cursor.y] = txt.rows[txt.cursor.y][0:txt.cursor.x]
	txt.cursor.y++
}

func (txt *TextEditor) moveRight() {
	if txt.cursor.x+1 >= len(txt.rows[txt.cursor.y]) {
		txt.cursor.y = Math.min(txt.cursor.y+1, len(txt.rows))
		txt.cursor.x = 0
	} else {
		txt.cursor.x = Math.min(txt.cursor.x+1, len(txt.rows[txt.cursor.y]))
	}
	//txt.cursor.x = Math.max()
}

func (txt *TextEditor) moveLeft() {
	if txt.cursor.x-1 <= 0 {
		txt.cursor.y = Math.max(txt.cursor.y-1, 1)
		txt.cursor.x = len(txt.rows[txt.cursor.y])
	} else {
		txt.cursor.x = Math.max(txt.cursor.x-1, 0)
	}
}

func (txt *TextEditor) moveWordLeft() {
	txt.cursor.x = len(regexp.MustCompile(`\S+\s*$`).ReplaceAllString(txt.rows[txt.cursor.y][:txt.cursor.x], ""))
}

func (txt *TextEditor) moveWordRight() {
	txt.cursor.x = len(txt.rows[txt.cursor.y]) - len(regexp.MustCompile(`^\s*\S+\s*`).ReplaceAllString(txt.rows[txt.cursor.y][txt.cursor.x:], ""))
}
